package arclog

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/dpsreport"
)

func (ctx *Context) RunPendingUploads(anonymous, detailedwvw bool, cancelCtx context.Context) {
	jobs, wg := ctx.StartWorkerPool(4, anonymous, detailedwvw)
	fmt.Println("Uploading pending logs...")
	ctx.EnqueuePending(jobs)
	close(jobs)

	wg.Wait()
	ctx.Logger.Debugw("All workers shut down...")
}

func (ctx *Context) RunWatchUploads(anonymous, detailedwvw bool, cancelCtx context.Context) {
	jobs, wg := ctx.StartWorkerPool(4, anonymous, detailedwvw)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			ctx.EnqueuePending(jobs)
		}
	}()

	err := ctx.NewWatcher(jobs, cancelCtx)
	if err != nil {
		ctx.Logger.Errorw("could not start watcher", "err", err)
		os.Exit(1)
	}
	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	<-cancelCtx.Done()
	fmt.Println("Shutting down...")
	if ctx.Watcher != nil {
		ctx.Watcher.Close()
	}
	close(jobs)
	wg.Wait()
	ctx.Logger.Debugw("all workers shut down")
	fmt.Println("All workers shut down")
}

func (ctx *Context) UploadLog(cbtlog database.Cbtlog, anonymous, detailedwvw bool) {
	filePath := ctx.Config.GetLogFilePath(cbtlog)

	exists, err := FileExists(filePath)
	if err != nil {
		ctx.Logger.Errorw("could not stat file", "file", filePath, "err", err)
		return
	}
	if !exists {
		ctx.Logger.Warnw("file missing, skipping upload", "file", filePath)

		ctx.updateLogStatus(cbtlog.ID, string(db.StatusFailed), string(db.ReasonFileMissing))
		return
	}

	ctx.updateLogStatus(cbtlog.ID, string(db.StatusUploading), string(db.ReasonUploading))
	ctx.Logger.Debugw("updated upload in db", "upload", cbtlog.Filename, "status", db.StatusUploading)

	ctx.RateLimiter.Wait()

	resp, err := ctx.DpsReportClient.UploadContent(filePath, dpsreport.UploadContentOptions{
		UserToken:   ctx.Config.UserToken,
		Anonymous:   anonymous,
		DetailedWvW: detailedwvw,
	})
	if err != nil {
		reason := db.ErrMapToReason(err)

		ctx.Logger.Errorw("upload failed", "file", filePath, "err", err, "reason", reason)

		ctx.updateLogStatus(cbtlog.ID, string(db.StatusFailed), reason)
		return
	}

	ctx.updateLogStatus(cbtlog.ID, string(db.StatusUploaded), string(db.ReasonSuccess), resp.Permalink)
	ctx.Logger.Infow("successfully uploaded to arcdps", "file", cbtlog.Filename)
}
