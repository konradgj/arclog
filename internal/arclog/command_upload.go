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

	ctx.EnqueuePending(jobs)
	close(jobs)

	wg.Wait()
	ctx.Logger.Debug("All workers shut down")
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
		ctx.Logger.Error("could not start watcher: %w", err)
		os.Exit(1)
	}
	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	<-cancelCtx.Done()
	fmt.Println("shutting down...")
	if ctx.Watcher != nil {
		ctx.Watcher.Close()
	}
	close(jobs)
	wg.Wait()
	ctx.Logger.Debug("all workers shut down")
	fmt.Println("All workers shut down")
}

func (ctx *Context) UploadLog(cbtlog database.Cbtlog, anonymous, detailedwvw bool) {
	err := ctx.St.Queries.UpdateCtblogUploadStatus(context.Background(), database.UpdateCtblogUploadStatusParams{
		UploadStatus:       string(db.StatusUploading),
		UploadStatusReason: string(db.ReasonUploading),
		ID:                 cbtlog.ID,
	})
	if err != nil {
		ctx.Logger.Error("error updating upload", "file", cbtlog.Filename, "err", err)
		return
	}
	ctx.Logger.Debug("updated upload in db", "upload", cbtlog.Filename, "status", db.StatusUploading)

	opts := dpsreport.UploadContentOptions{
		UserToken:   ctx.Config.UserToken,
		Anonymous:   anonymous,
		DetailedWvW: detailedwvw,
	}

	ctx.RateLimiter.Wait()
	resp, err := ctx.DpsReportClient.UploadContent(ctx.Config.GetLogFilePath(cbtlog), opts)
	if err != nil && resp == nil {
		ctx.Logger.Error("could not upload", "err", err)

		err = ctx.St.Queries.UpdateCtblogUploadStatus(context.Background(), database.UpdateCtblogUploadStatusParams{
			UploadStatus:       string(db.StatusFailed),
			UploadStatusReason: string(db.ReasonHttp),
			ID:                 cbtlog.ID,
		})
		if err != nil {
			ctx.Logger.Error("error updating upload in db", "file", cbtlog.Filename, "err", err)
			return
		}
	}
	if err != nil && resp != nil {
		ctx.Logger.Error("error decoding response", "err", err)
	}

	err = ctx.St.Queries.UpdateCbtlogUrl(context.Background(), database.UpdateCbtlogUrlParams{
		UploadStatus:       string(db.StatusUploaded),
		UploadStatusReason: string(db.ReasonSuccess),
		Url:                db.WrapNullStr(resp.Permalink),
		ID:                 cbtlog.ID,
	})
	if err != nil {
		ctx.Logger.Error("error updating upload", "file", cbtlog.ID, "err", err)
		return
	}

	ctx.Logger.Info("successfully uploaded to arcdps", "file", cbtlog.Filename)
}
