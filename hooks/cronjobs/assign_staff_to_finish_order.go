package cronjobs

import (
	"errors"
	"log/slog"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/thaiha1607/4sq_server/hooks/shared"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
)

func assignStaffToCompleteUnfinishedOrdersUpdatedMoreThanThreeDaysAgo(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		lastThreeDays := time.Now().AddDate(0, 0, -3).Format(types.DefaultDateLayout)

		jobId := "assignStaffToCompleteUnfinishedOrdersUpdatedMoreThanThreeDaysAgo"

		scheduler.MustAdd(jobId, "0 4 * * 1-5", func() {
			app.Logger().Info("Running job: " + jobId)
			filteredStatus := []string{
				order_status.PartiallyDelivered.ID(),
				order_status.PartiallyShipped.ID(),
			}
			orderRecords, err := app.Dao().FindRecordsByExpr(
				"orders",
				dbx.NewExp(
					"statusCodeId IN {:status} AND updated < {:date}",
					dbx.Params{
						"status": filteredStatus,
						"date":   lastThreeDays,
					},
				),
			)
			if err != nil {
				return
			}
			count := 0
			err = app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
				var transactionErr error = nil
				for _, orderRecord := range orderRecords {
					assignStaffErr := shared.AssignWarehouseStaff(txDao, app.Logger(), orderRecord)
					if assignStaffErr != nil {
						transactionErr = errors.Join(transactionErr, assignStaffErr)
					} else {
						count++
					}
				}
				return transactionErr
			})
			if err != nil {
				app.Logger().Error(
					"Error assigning staff to complete unfinished orders updated more than three days ago",
					"err",
					err,
				)
				return
			}

			app.Logger().Info("Modified pending internal orders older than 48 hours", slog.Int("count", count))
		})

		scheduler.Start()
		return nil
	})
}
