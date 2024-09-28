package record_hooks

import (
	"net/http"

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
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
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

func updateOrderWhenShipmentDelivered(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("shipments").Add(func(e *core.RecordUpdateEvent) error {
		allowedStatus := []string{
			shipment_status.Delivered.ID(),
		}
		if !lo.Contains(allowedStatus, e.Record.GetString("statusCodeId")) {
			shipmentItems, err := dbquery.GetShipmentItemsByShipmentId(app.Dao(), e.Record.Id)
			if err != nil {
				return apis.NewNotFoundError("", map[string]validation.Error{
					"shipmentId": validation.NewError("not_found", "Shipment not found"),
				})
			}
			for _, shipmentItem := range shipmentItems {
				orderItem, err := dbquery.GetSingleOrderItem(app.Dao(), shipmentItem.OrderItemId)
				if err != nil {
					return apis.NewNotFoundError("", nil)
				}
				orderItem.ReceivedQty += shipmentItem.Qty
				if orderItem.ReceivedQty > orderItem.OrderedQty {
					return apis.NewBadRequestError("", map[string]validation.Error{
						"qty": validation.NewError("invalid_qty", "Received quantity is greater than ordered quantity"),
					})
				}
				orderItem.MarkAsNotNew()
				if err := app.Dao().Save(orderItem); err != nil {
					return err
				}
			}
			order, err := dbquery.GetSingleOrder(app.Dao(), e.Record.GetString("orderId"))
			if err != nil {
				return apis.NewNotFoundError("", nil)
			}
			statusAllowed := []string{
				order_status.Shipped.ID(),
				order_status.PartiallyShipped.ID(),
				order_status.FailedDeliveryAttempt.ID(),
				order_status.PartiallyDelivered.ID(),
			}
			if !lo.Contains(statusAllowed, order.StatusCodeId) {
				return nil
			}
			// Get all order items
			orderItems, err := dbquery.GetOrderItemsByOrderId(app.Dao(), order.Id)
			if err != nil {
				return apis.NewNotFoundError("", nil)
			}
			allDelivered := lo.EveryBy(orderItems, func(orderItem *custom_models.OrderItem) bool {
				return orderItem.ReceivedQty == orderItem.OrderedQty
			})
			if allDelivered {
				order.StatusCodeId = order_status.Delivered.ID()
			} else {
				if order.StatusCodeId == order_status.PartiallyDelivered.ID() {
					return nil
				}
				order.StatusCodeId = order_status.PartiallyDelivered.ID()
			}
			order.MarkAsNotNew()
			if err := app.Dao().Save(order); err != nil {
				return err
			}
			return nil
		}
		return nil
	})
}
