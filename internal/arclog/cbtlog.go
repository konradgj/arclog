package arclog

import (
	"context"

	"github.com/konradgj/arclog/internal/database"
	"github.com/konradgj/arclog/internal/db"
)

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
