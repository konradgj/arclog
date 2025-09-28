package arclog

import (
	"sync"

	"github.com/konradgj/arclog/internal/database"
)

type UploadJob struct {
	Cbtlog database.Cbtlog
}

func (ctx *Context) StartWorkerPool(numWorkers int, anonymous, detailedwvw bool) (chan<- UploadJob, *sync.WaitGroup) {
	jobs := make(chan UploadJob, 100)
	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				ctx.Logger.Debugw("worker started", "worker", id, "file", job.Cbtlog.Filename)
				ctx.UploadLog(job.Cbtlog, anonymous, detailedwvw)
			}
		}(i)
	}

	return jobs, &wg
}

func (ctx *Context) EnqueueCbtlogs(cbtlogs []database.Cbtlog, jobs chan<- UploadJob) {
	for _, l := range cbtlogs {
		jobs <- UploadJob{Cbtlog: l}
	}
}
