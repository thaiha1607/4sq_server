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
	"github.com/thaiha1607/4sq_server/custom_models"
	"github.com/thaiha1607/4sq_server/dbquery"
	"github.com/thaiha1607/4sq_server/utils/enum/assignment_status"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
)

// Change pending internal orders older than 5 days to CANCELLED
func modifyPendingInternalOrdersOlderThan5Days(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		lastFiveDays := time.Now().AddDate(0, 0, -5).Format(types.DefaultDateLayout)

		jobId := "modifyPendingInternalOrdersOlderThan5Days"

		scheduler.MustAdd(jobId, "0 4 * * 1-5", func() {
			app.Logger().Info("Running job: " + jobId)
			internalOrders := []*custom_models.InternalOrder{}
			err := custom_models.InternalOrderQuery(app.Dao()).
				Where(dbx.NewExp(
					"statusCodeId = {:status} AND updated < {:date}",
					dbx.Params{
						"status": order_status.Pending.ID(),
						"date":   lastFiveDays,
					},
				)).
				All(&internalOrders)
			if err != nil {
				return
			}
			count := 0
			err = app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
				var transactionErr error = nil
				for _, internalOrder := range internalOrders {
					internalOrder.StatusCodeId = order_status.Cancelled.ID()
					internalOrder.Note = "Hủy đơn hàng tự động do quá thời gian chờ xác nhận (5 ngày)"
					internalOrder.ShipmentId = ""
					if err := txDao.Save(internalOrder); err != nil {
						transactionErr = errors.Join(transactionErr, err)
						continue
					}
					warehouseAssignments, err := dbquery.GetWarehouseAssignmentsByInternalOrderId(txDao, internalOrder.Id)
					if err != nil {
						transactionErr = errors.Join(transactionErr, err)
						continue
					}
					for _, warehouseAssignment := range warehouseAssignments {
						warehouseAssignment.Status = string(assignment_status.Cancelled)
						warehouseAssignment.Note = "Hủy đơn hàng tự động do quá thời gian chờ xác nhận (5 ngày)"
						if err := txDao.Save(warehouseAssignment); err != nil {
							transactionErr = errors.Join(transactionErr, err)
							continue
						}
					}
					count++
				}
				return nil
			})
			if err != nil {
				app.Logger().Error("Failed to modify pending internal orders older than 5 days", "error", err)
				return
			}
			app.Logger().Info("Modified pending internal orders older than 5 days", slog.Int("count", count))
		})
		scheduler.Start()
		return nil
	})
}
