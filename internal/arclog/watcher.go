package arclog

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/logger"
)

func (ctx *Context) RunWatch(cancelCtx context.Context) {
	err := ctx.NewWatcher(nil, cancelCtx)
	if err != nil {
		logger.Error("Could not start watcher", "err", err)
	}
	defer ctx.Watcher.Close()
	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	<-cancelCtx.Done()
	logger.Info("shutting down...")
	if ctx.Watcher != nil {
		ctx.Watcher.Close()
	}
}

func (ctx *Context) NewWatcher(jobs chan<- UploadJob, cancelCtx context.Context) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("could not create new watcher: %w", err)
	}
	ctx.Watcher = watcher

	go func() {
		for {
			select {
			case <-cancelCtx.Done():
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if !event.Has(fsnotify.Create) {
					continue
				}

				st, err := os.Stat(event.Name)
				if err != nil {
					logger.Error("error in filepath", "err", err)
					continue
				}

				if st.IsDir() && !slices.Contains(watcher.WatchList(), event.Name) {
					watcher.Add(event.Name)
				}

				if !strings.HasSuffix(event.Name, ".zevtc") {
					continue
				}

				logger.Debug("new event", "event", event.Name)
				upload, err := ctx.St.Queries.CreateUpload(context.Background(), database.CreateUploadParams{
					FilePath: event.Name,
				})
				if err != nil {
					logger.Error("error adding upload to db", "err", err)
					continue
				}

				logger.Info("added upload to db", "file_path", upload.FilePath)

				if jobs != nil {
					jobs <- UploadJob{Upload: upload}
					logger.Debug("enqueued upload job", "file", upload.FilePath)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					logger.Warn("watcher errors channel closed")
					return
				}
				if err != nil {
					logger.Error("watcher error", "err", err)
					return
				}
			}
		}
	}()

	err = ctx.Config.Unmarshal()
	if err != nil {
		return fmt.Errorf("could not umarshal config: %w", err)
	}

	err = watcher.Add(ctx.Config.LogPath)
	if err != nil {
		return fmt.Errorf("could not add path to watcher: %w", err)
	}
	// Add subdirs recursivly
	err = filepath.WalkDir(ctx.Config.LogPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not add path to watcher: %w", err)
	}

	return nil
}
