package cronjobs

import (
	"errors"
	"log/slog"
	"time"

	"example.com/4sq_server/hooks/shared"
	"example.com/4sq_server/utils/enum/order_status"
	pocketbase "github.com/AlperRehaYAZGAN/postgresbase"
	"github.com/AlperRehaYAZGAN/postgresbase/core"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/tools/cron"
	"github.com/AlperRehaYAZGAN/postgresbase/tools/types"
	"github.com/pocketbase/dbx"
)

func assignStaffToCompleteUnfinishedOrdersUpdatedMoreThan3DaysAgo(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		lastThreeDays := time.Now().AddDate(0, 0, -3).Format(types.DefaultDateLayout)

		jobId := "assignStaffToCompleteUnfinishedOrdersUpdatedMoreThan3DaysAgo"

		scheduler.MustAdd(jobId, "0 4 * * 1-5", func() {
			app.Logger().Info("Running job: " + jobId)
			orderRecords, err := app.Dao().FindRecordsByExpr(
				"orders",
				dbx.NewExp(
					"statusCodeId IN ({:partial_delivered}, {:partial_shipped}) AND updated < {:date}",
					dbx.Params{
						"partial_delivered": order_status.PartiallyDelivered.ID(),
						"partial_shipped":   order_status.PartiallyShipped.ID(),
						"date":              lastThreeDays,
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
					"Error assigning staff to complete unfinished orders updated more than 3 days ago",
					"err",
					err,
				)
				return
			}

			app.Logger().Info("Modified pending internal orders older than 3 days", slog.Int("count", count))
		})

		scheduler.Start()
		return nil
	})
}

func assignStaffToCompleteConfirmedAndProcessingOrdersThatHaveAllInternalOrdersCancelledUpdatedMoreThan3DaysAgo(
	app *pocketbase.PocketBase,
) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		threeDaysAgo := time.Now().AddDate(0, 0, -3).Format(types.DefaultDateLayout)

		jobId := "assignStaffToCompleteConfirmedAndProcessingOrdersThatHaveAllInternalOrdersCancelledUpdatedMoreThan3DaysAgo"

		scheduler.MustAdd(jobId, "0 5 * * 1-5", func() {
			app.Logger().Info("Running job: " + jobId)
			orderRecords, err := app.Dao().FindRecordsByExpr(
				"orders",
				dbx.NewExp(
					"statusCodeId IN ({:confirmed}, {:processing}) AND updated < {:date} AND id NOT IN (SELECT rootOrderId FROM internal_orders WHERE statusCodeId != {:cancelled})",
					dbx.Params{
						"confirmed":  order_status.Confirmed.ID(),
						"processing": order_status.Processing.ID(),
						"date":       threeDaysAgo,
						"cancelled":  order_status.Cancelled.ID(),
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
					"Error assigning staff to complete confirmed and processing orders that have all internal orders cancelled",
					"err",
					err,
				)
				return
			}

			app.Logger().Info("Modified confirmed and processing orders with all internal orders cancelled", slog.Int("count", count))
		})

		scheduler.Start()
		return nil
	})
}
