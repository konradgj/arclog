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
		ctx.Logger.Infow("File already exists in db", "name", cbtlog.Filename, "path", cbtlog.RelativePath.String)
		return &cbtlog, nil
	}

	cbtlog, err = ctx.St.Queries.CreateCbtlog(context.Background(), database.CreateCbtlogParams{
		Filename:     name,
		RelativePath: relPath,
	})
	if err != nil {
		return nil, fmt.Errorf("could not add file %s to db: %w", name, err)
	}

	ctx.Logger.Infow("Added file to db", "name", cbtlog.Filename, "path", cbtlog.RelativePath.String)
	return &cbtlog, nil
}

func (ctx *Context) updateLogStatus(id int64, status, reason string) {
	err := ctx.St.Queries.UpdateCtblogUploadStatus(context.Background(), database.UpdateCtblogUploadStatusParams{
		UploadStatus:       status,
		UploadStatusReason: reason,
		ID:                 id,
	})
	if err != nil {
		ctx.Logger.Errorw("could not update log status", "logID", id, "err", err)
	}
}

func (ctx *Context) getCbtlogsByStatus(status db.UploadStatus) ([]database.Cbtlog, error) {
	cbtlogs, err := ctx.St.Queries.ListCbtlogsByUploadStatus(context.Background(), string(status))
	if err != nil {
		return nil, fmt.Errorf("could not list pending uploads: %w", err)
	}

	return cbtlogs, nil
}

func (ctx *Context) listCbtlogsByFilters(uploadStatus, relativePath, date, fromDate, toDate string, challengeMode, encounterSuccess *bool) ([]database.Cbtlog, error) {
	cbtlogs, err := ctx.St.Queries.ListCbtlogsByFilters(context.Background(), database.ListCbtlogsByFiltersParams{
		UploadStatus:     db.WrapNullStr(uploadStatus),
		RelativePath:     db.WrapNullStr(relativePath),
		Date:             db.WrapNullStr(date),
		FromDate:         db.WrapNullStr(fromDate),
		ToDate:           db.WrapNullStr(toDate),
		ChallengeMode:    db.WrapNullBool(challengeMode),
		EncounterSuccess: db.WrapNullBool(encounterSuccess),
	})
	if err != nil {
		return nil, fmt.Errorf("could not list logs: %w", err)
	}

	return cbtlogs, nil
}
