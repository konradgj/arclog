package arclog

import (
	"context"
	"sync"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
	"github.com/konradgj/arclog/internal/logger"
)

type UploadJob struct {
	Upload database.Upload
}

func (ctx *Context) StartWorkerPool(numWorkers int, anonymous, detailedwvw bool) (chan<- UploadJob, *sync.WaitGroup) {
	jobs := make(chan UploadJob, 100)
	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				logger.Debug("worker started", "worker", id, "file", job.Upload.FilePath)
				ctx.UploadLog(job, anonymous, detailedwvw)
			}
		}(i)
	}

	return jobs, &wg
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
