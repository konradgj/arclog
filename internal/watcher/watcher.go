package watcher

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/app"
	"github.com/konradgj/arclog/internal/database"
)

func Watch(ctx *app.Context) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) && strings.Contains(event.Name, ".zevtc") {
					ctx.Log.Info("new event", "event", event.Name)
					upload, err := ctx.St.Queries.CreateUpload(context.Background(), database.CreateUploadParams{
						FilePath: event.Name,
					})
					if err != nil {
						ctx.Log.Error("error adding upload to db", "err", err)
					}
					ctx.Log.Info("added upload to db", "file_path", upload.FilePath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				ctx.Log.Error("Error", "err", err)
			}
		}
	}()

	if err != nil {
		ctx.Log.Error("Could not umarshal config", "err", err)
	}

	// Add a path.
	err = watcher.Add(ctx.Config.LogPath)
	if err != nil {
		ctx.Log.Error("Could not add path to watcher", "err", err)
		os.Exit(1)
	}
	// Add subdirs recursivly
	err = filepath.WalkDir(ctx.Config.LogPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		ctx.Log.Error("Could not add path to watcher", "err", err)
		os.Exit(1)
	}

	fmt.Printf("Started watching dir: %s\n", ctx.Config.LogPath)

	// Block main goroutine forever.
	<-make(chan struct{})
}
