package arclog

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/konradgj/arclog/internal/db"
)

func (ctx *Context) RunListCbtlogsByFilter(uploadStatus, relativePath, date, fromDate, toDate string, challengeMode, encounterSucces *bool) {
	cbtlogs, err := ctx.listCbtlogsByFilters(uploadStatus, relativePath, date, fromDate, toDate, challengeMode, encounterSucces)
	if err != nil {
		ctx.Logger.Errorw("could not list logs", "err", err)
		return
	}

	if len(cbtlogs) == 0 {
		fmt.Println("No logs found matching current filters...")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "FILENAME\tRELATIVE_PATH\tURL\tCM\tSUCCESS\tUPLOAD_STATUS\tUPLOAD_STATUS_REASON")

	for _, row := range cbtlogs {
		relPath := db.PrintNullStr(row.RelativePath)
		url := db.PrintNullStr(row.Url)
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\t%d\t%d\t%s\t%s\n",
			row.Filename,
			relPath,
			url,
			db.PrintNullBool(row.ChallengeMode),
			db.PrintNullBool(row.EncounterSuccess),
			row.UploadStatus,
			row.UploadStatusReason,
		)
	}

	w.Flush()
}
