package record_hooks

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/samber/lo"
	"github.com/thaiha1607/4sq_server/custom_models"
	"github.com/thaiha1607/4sq_server/dbquery"
	"github.com/thaiha1607/4sq_server/utils"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
	"github.com/thaiha1607/4sq_server/utils/enum/shipment_status"
)

func forbidInvalidShipmentStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("shipments").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != shipment_status.Pending.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError("invalid_status_code", "When creating an shipment, the status code must be 'Pending'"),
			})
		}
		return nil
	})
	app.OnRecordBeforeUpdateRequest("shipments").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("shipments", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Shipment not found", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			value, ok := utils.ShipmentStatusCodeTransitions[old.GetString("statusCodeId")]
			if !ok {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"statusCodeId": validation.NewError("invalid_status_code", "Invalid status code"),
				})
			}
			if !lo.Contains(value, e.Record.GetString("statusCodeId")) {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"statusCodeId": validation.NewError("invalid_status_code", "Invalid status code transition"),
				})
			}
		}
		return nil
	})
}

func changeOrderStatusToDelivered(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("shipments").Add(func(e *core.RecordUpdateEvent) error {
		shipments, err := dbquery.GetShipmentsByOrderId(app.Dao(), e.Record.GetString("orderId"))
		if err != nil {
			app.Logger().Error("Failed to get shipments", "error", err)
			return nil
		}
		order, err := dbquery.GetSingleOrder(app.Dao(), e.Record.GetString("orderId"))
		if err != nil {
			app.Logger().Error("Failed to get order", "error", err)
			return nil
		}
		allDelivered := len(lo.Filter(shipments, func(shipment *custom_models.Shipment, _ int) bool {
			return shipment.StatusCodeId == shipment_status.Delivered.ID()
		})) == len(shipments)
		isThisShipmentDelivered := e.Record.GetString("statusCodeId") == order_status.Delivered.ID()
		if (order.StatusCodeId == order_status.Shipped.ID() ||
			order.StatusCodeId == order_status.PartiallyShipped.ID() ||
			order.StatusCodeId == order_status.FailedDeliveryAttempt.ID() ||
			order.StatusCodeId == order_status.PartiallyDelivered.ID()) &&
			(allDelivered || isThisShipmentDelivered) {
			var newStatusCodeId string
			if allDelivered {
				newStatusCodeId = order_status.Delivered.ID()
			} else {
				newStatusCodeId = order_status.PartiallyDelivered.ID()
			}
			order.StatusCodeId = newStatusCodeId
			order.MarkAsNotNew()
			if err := app.Dao().Save(order); err != nil {
				app.Logger().Error("Failed to save order", "error", err)
				return nil
			}
			app.Logger().Info("Order status changed to Shipped", "order_id", order.Id)
		}

		return nil
	})
}
