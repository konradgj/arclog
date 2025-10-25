package arclog

import (
	"context"

	"github.com/konradgj/arclog/internal/fsutil"
)

func (ctx *Context) RunRmCmd(filename string, delete bool) {
	deletedLog, err := ctx.St.Queries.DeleteCbtlogByFilename(context.Background(), filename)
	if err != nil {
		ctx.Logger.Errorw("could not get file from db", "filename", filename, "err", err)
		return
	}
	ctx.Logger.Infow("removed log from db", "filename", deletedLog.Filename)

	if delete {
		path := ctx.Config.GetLogFilePath(deletedLog)
		err := fsutil.RmFile(path)
		if err != nil {
			ctx.Logger.Errorw("could not delete file", "file", filename, "err", err)
		}
		ctx.Logger.Infow("deleted file from os", "path", path)
	}
}
