package arclog

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/dpsreport"
	"github.com/konradgj/arclog/internal/logger"
)

func (ctx *Context) RunPendingUploads(anonymous, detailedwvw bool, cancelCtx context.Context) {
	jobs, wg := ctx.StartWorkerPool(4, anonymous, detailedwvw)

	ctx.EnqueuePending(jobs)
	close(jobs)

	wg.Wait()
	logger.Info("All workers shut down")
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
		logger.Error("could not start watcher: %w", err)
		os.Exit(1)
	}
	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	<-cancelCtx.Done()
	logger.Info("shutting down...")
	if ctx.Watcher != nil {
		ctx.Watcher.Close()
	}
	close(jobs)
	wg.Wait()
	logger.Info("All workers shut down")
}

func (ctx *Context) UploadLog(job UploadJob, anonymous, detailedwvw bool) {
	err := ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
		Status:       string(db.StatusUploading),
		StatusReason: string(db.ReasonUploading),
		ID:           job.Upload.ID,
	})
	if err != nil {
		logger.Error("error updating upload", "file", job.Upload.ID, "err", err)
		return
	}
	logger.Debug("updated upload in db", "upload", job.Upload.FilePath, "status", db.StatusUploading)

	opts := dpsreport.UploadContentOptions{
		UserToken:   ctx.Config.UserToken,
		Anonymous:   anonymous,
		DetailedWvW: detailedwvw,
	}

	ctx.RateLimiter.Wait()
	resp, err := ctx.DpsReportClient.UploadContent(job.Upload.FilePath, opts)
	if err != nil && resp == nil {
		logger.Error("could not upload", "err", err)

		err = ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
			Status:       string(db.StatusFailed),
			StatusReason: string(db.ReasonQueueFull),
			ID:           job.Upload.ID,
		})
		if err != nil {
			logger.Error("error updating upload", "file", job.Upload.ID, "err", err)
			return
		}
	}
	if err != nil && resp != nil {
		logger.Error("error decoding", "err", err)
	}

	err = ctx.St.Queries.UpdateUploadUrl(context.Background(), database.UpdateUploadUrlParams{
		Status:       string(db.StatusUploaded),
		StatusReason: string(db.ReasonSuccess),
		Url:          db.WrapNullStr(resp.Permalink),
		ID:           job.Upload.ID,
	})
	if err != nil {
		logger.Error("error updating upload", "file", job.Upload.ID, "err", err)
		return
	}

	logger.Info("successfully uploaded", "file", job.Upload.FilePath)
}

func (ctx *Context) SimulateUpload(job UploadJob) {
	// Simulate variable upload time
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	success := rand.Intn(100) < 80 // 80% success rate

	if success {
		url := "https://testurl.com/" + strconv.Itoa(int(job.Upload.ID))
		ctx.St.Queries.UpdateUploadUrl(context.Background(), database.UpdateUploadUrlParams{
			Status:       string(db.StatusUploaded),
			StatusReason: string(db.ReasonSuccess),
			Url:          db.WrapNullStr(url),
			ID:           job.Upload.ID,
		})
		logger.Debug("Simulated upload success",
			"file", job.Upload.FilePath,
			"delay", delay,
		)
	} else {
		ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
			Status:       string(db.StatusFailed),
			StatusReason: string(db.ReasonQueueFull),
			ID:           job.Upload.ID,
		})
		logger.Warn("Simulated upload failed",
			"file", job.Upload.FilePath,
			"delay", delay,
		)
	}
}
