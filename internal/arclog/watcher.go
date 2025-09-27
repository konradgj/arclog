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
)

func (ctx *Context) RunWatch(cancelCtx context.Context) {
	err := ctx.NewWatcher(nil, cancelCtx)
	if err != nil {
		ctx.Logger.Error("could not start watcher", "err", err)
		os.Exit(1)
	}
	defer ctx.Watcher.Close()
	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	<-cancelCtx.Done()
	fmt.Println("\nshutting down...")
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
					ctx.Logger.Error("error in filepath", "err", err)
					continue
				}

				if st.IsDir() && !slices.Contains(watcher.WatchList(), event.Name) {
					watcher.Add(event.Name)
				}

				if !strings.HasSuffix(event.Name, ".zevtc") {
					continue
				}

				ctx.Logger.Debug("new event", event.Name)
				upload, err := ctx.St.Queries.CreateUpload(context.Background(), database.CreateUploadParams{
					FilePath: event.Name,
				})
				if err != nil {
					ctx.Logger.Error("error adding upload to db", "err", err)
					continue
				}

				ctx.Logger.Info("added upload to db", "file_path", upload.FilePath)

				if jobs != nil {
					jobs <- UploadJob{Upload: upload}
					ctx.Logger.Debug("enqueued upload job", "file", upload.FilePath)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					ctx.Logger.Warn("watcher errors channel closed")
					return
				}
				if err != nil {
					ctx.Logger.Error("watcher error", "err", err)
					return
				}
			}
		}
	}()

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
