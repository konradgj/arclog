package arclog

import (
	"fmt"
)

func (ctx *Context) AddLogsToDb(paths []string) {
	var logPaths []string
	for _, path := range paths {
		ps, err := GetAllFilePaths(path)
		if err != nil {
			ctx.Logger.Errorw("could not get logs", "err", err)
			return
		}
		if len(ps) == 0 {
			fmt.Printf("Found 0 logs in given path: %s\n", path)
			continue
		}

		logPaths = append(logPaths, ps...)
	}

	for _, logPath := range logPaths {
		_, err := ctx.addCbtlogToDb(logPath)
		if err != nil {
			ctx.Logger.Errorw("could not add log to db", "err", err)
		}
	}
}
