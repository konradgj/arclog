package watcher

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/konradgj/arclog/internal/appconfig"
	"github.com/konradgj/arclog/internal/logger"
)

func Watch() {
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
				logger.Info("New event", "event", event)
				if event.Has(fsnotify.Write) {
					logger.Debug("New event", "modified", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error("Error", "err", err)
			}
		}
	}()

	cfg, err := appconfig.Unmarshal()
	if err != nil {
		logger.Error("Could not umarshal config", "err", err)
	}

	// Add a path.
	err = watcher.Add(cfg.LogPath)
	if err != nil {
		logger.Error("Could not add path to watcher", "err", err)
		os.Exit(1)
	}
	// Add subdirs recursivly
	err = filepath.WalkDir(cfg.LogPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		logger.Error("Could not add path to watcher", "err", err)
		os.Exit(1)
	}

	fmt.Printf("Started watching dir: %s\n", cfg.LogPath)

	// Block main goroutine forever.
	<-make(chan struct{})
}
