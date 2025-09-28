package arclog

import (
	"context"
	"fmt"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
)

func (ctx *Context) addCbtlogToDb(logPath string) (*database.Cbtlog, error) {
	var cbtlog database.Cbtlog

	name, relPath, err := ctx.Config.GetLogNameAndRelativePath(logPath)
	if err != nil {
		return nil, fmt.Errorf("error getting filename: %w", err)
	}

	cbtlog, err = ctx.St.Queries.GetCbtlogByFileName(context.Background(), name)
	if err == nil {
		ctx.Logger.Infow("File already exists in db", "name", name)
		return &cbtlog, nil
	}

	cbtlog, err = ctx.St.Queries.CreateCbtlog(context.Background(), database.CreateCbtlogParams{
		Filename:     name,
		RelativePath: relPath,
	})
	if err != nil {
		return nil, fmt.Errorf("could not add file %s to db: %w", name, err)
	}
	return &cbtlog, nil
}

func (ctx *Context) updateLogStatus(id int64, status, reason string, url ...string) {
	if len(url) > 0 {
		err := ctx.St.Queries.UpdateCbtlogUrl(context.Background(), database.UpdateCbtlogUrlParams{
			UploadStatus:       status,
			UploadStatusReason: reason,
			Url:                db.WrapNullStr(url[0]),
			ID:                 id,
		})
		if err != nil {
			ctx.Logger.Errorw("could not update log with URL", "logID", id, "err", err)
		}
		return
	}

	err := ctx.St.Queries.UpdateCtblogUploadStatus(context.Background(), database.UpdateCtblogUploadStatusParams{
		UploadStatus:       status,
		UploadStatusReason: reason,
		ID:                 id,
	})
	if err != nil {
		ctx.Logger.Errorw("could not update log status", "logID", id, "err", err)
	}
}

func (ctx *Context) getPendingCbtlogs() ([]database.Cbtlog, error) {
	cbtlogs, err := ctx.St.Queries.ListCbtlogsByUploadStatus(context.Background(), string(db.StatusPending))
	if err != nil {
		return nil, fmt.Errorf("could not list pending uploads: %w", err)
	}

	return cbtlogs, nil
}
