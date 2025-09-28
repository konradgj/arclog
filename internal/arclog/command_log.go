package arclog

import (
	"context"
	"fmt"

	"github.com/konradgj/arclog/internal/database"
)

func (ctx *Context) AddLogsToDb(path string) {
	logPaths, err := GetAllFilePaths(path)
	if err != nil {
		ctx.Logger.Errorw("could not get logs", "err", err)
		return
	}
	if len(logPaths) == 0 {
		fmt.Println("Found 0 logs in given path")
		return
	}

	for _, logPath := range logPaths {
		name, relPath, err := ctx.Config.GetLogNameAndRelativePath(logPath)
		if err != nil {
			ctx.Logger.Errorw("error getting filename", "err", err)
		}

		_, err = ctx.St.Queries.GetCbtlogByFileName(context.Background(), name)
		if err == nil {
			ctx.Logger.Infow("File already exists in db", "name", name)
			continue
		}

		_, err = ctx.St.Queries.CreateCbtlog(context.Background(), database.CreateCbtlogParams{
			Filename:     name,
			RelativePath: relPath,
		})
		if err != nil {
			ctx.Logger.Errorw("could not add file to db", "file", name, "err", err)
		}
	}
}
