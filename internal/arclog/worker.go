package arclog

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/logger"
)

type UploadJob struct {
	Upload database.Upload
}

func (ctx *Context) StartWorkerPool(numWorkers int) chan<- UploadJob {
	jobs := make(chan UploadJob, 100)

	for i := range numWorkers {
		go func(id int) {
			for job := range jobs {
				logger.Debug("Worker started", "worker", id, "file", job.Upload.FilePath)

				err := ctx.St.Queries.UpdateUploadStatus(context.Background(), database.UpdateUploadStatusParams{
					Status:       string(db.StatusUploading),
					StatusReason: string(db.ReasonUploading),
					ID:           job.Upload.ID,
				})
				if err != nil {
					logger.Error("error updating upload", "worker", id, "err", err)
					return
				}
				logger.Debug("updated upload in db", "upload", job.Upload.FilePath, "status", db.StatusUploading)

				ctx.SimulateUpload(job)
			}
		}(i)
	}

	return jobs
}

func (ctx *Context) EnqueuePending(jobs chan<- UploadJob) {
	uploads, err := ctx.St.Queries.ListUploadsByStatus(context.Background(), string(db.StatusPending))
	if err != nil {
		logger.Error("could not list pending uploads", "err", err)
		return
	}

	for _, u := range uploads {
		jobs <- UploadJob{Upload: u}
	}
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
