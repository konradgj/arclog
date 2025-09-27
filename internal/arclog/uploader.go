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

func (ctx *Context) UploadLog(job UploadJob, anonymous, detailedwvw bool) {
	err := ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
		Status:       string(db.StatusUploading),
		StatusReason: string(db.ReasonUploading),
		ID:           job.Upload.ID,
	})
	if err != nil {
		ctx.Logger.Error("error updating upload", "file", job.Upload.FilePath, "err", err)
		return
	}
	ctx.Logger.Debug("updated upload in db", "upload", job.Upload.FilePath, "status", db.StatusUploading)

	opts := dpsreport.UploadContentOptions{
		UserToken:   ctx.Config.UserToken,
		Anonymous:   anonymous,
		DetailedWvW: detailedwvw,
	}

	ctx.RateLimiter.Wait()
	resp, err := ctx.DpsReportClient.UploadContent(job.Upload.FilePath, opts)
	if err != nil && resp == nil {
		ctx.Logger.Error("could not upload", "err", err)

		err = ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
			Status:       string(db.StatusFailed),
			StatusReason: string(db.ReasonHttp),
			ID:           job.Upload.ID,
		})
		if err != nil {
			ctx.Logger.Error("error updating upload in db", "file", job.Upload.FilePath, "err", err)
			return
		}
	}
	if err != nil && resp != nil {
		ctx.Logger.Error("error decoding response", "err", err)
	}

	err = ctx.St.Queries.UpdateUploadUrl(context.Background(), database.UpdateUploadUrlParams{
		Status:       string(db.StatusUploaded),
		StatusReason: string(db.ReasonSuccess),
		Url:          db.WrapNullStr(resp.Permalink),
		ID:           job.Upload.ID,
	})
	if err != nil {
		ctx.Logger.Error("error updating upload", "file", job.Upload.ID, "err", err)
		return
	}

	ctx.Logger.Info("successfully uploaded to arcdps", "file", job.Upload.FilePath)
}
