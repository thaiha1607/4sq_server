package cronjobs

import (
	"log/slog"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/thaiha1607/4sq_server/custom_models"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
)

func modifyPendingInternalOrdersOlderThan48Hours(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		before24Hours := time.Now().AddDate(0, 0, -2).Format(types.DefaultDateLayout)

		jobId := "modifyPendingInternalOrdersOlderThan48Hours"

		scheduler.MustAdd(jobId, "0 12 * * *", func() {
			app.Logger().Info("Running job: " + jobId)
			internalOrders := []*custom_models.InternalOrder{}
			err := custom_models.InternalOrderQuery(app.Dao()).
				Where(dbx.NewExp(
					"statusCodeId = {:status} AND created < {:date}",
					dbx.Params{
						"status": order_status.Pending.ID(),
						"date":   before24Hours,
					},
				)).
				All(&internalOrders)
			if err != nil {
				return
			}
			count := 0
			for _, internalOrder := range internalOrders {
				internalOrder.ShipmentId = ""
				if err := app.Dao().Save(internalOrder); err != nil {
					return
				}
				count++
			}
			app.Logger().Info("Modified pending internal orders older than 48 hours", slog.Int("count", count))
		})

		scheduler.Start()
		return nil
	})
}
