package record_hooks

import (
	"errors"
	"net/http"

	"example.com/4sq_server/custom_models"
	"example.com/4sq_server/dbquery"
	"example.com/4sq_server/utils"
	"example.com/4sq_server/utils/enum/order_status"
	"example.com/4sq_server/utils/enum/shipment_status"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/samber/lo"
)

func forbidInvalidShipmentStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("shipments").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") == shipment_status.Pending.ID() {
			return nil
		}
		return apis.NewBadRequestError("", map[string]validation.Error{
			"statusCodeId": validation.NewError("invalid_status_code", "When creating an shipment, the status code must be 'Pending'"),
		})
	})
	app.OnRecordBeforeUpdateRequest("shipments").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("shipments", e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		if old.GetString("statusCodeId") == e.Record.GetString("statusCodeId") {
			return nil
		}
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
		return nil
	})
}

func updateOrderWhenShipmentDelivered(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("shipments").Add(func(e *core.RecordUpdateEvent) error {
		allowedStatus := []string{
			shipment_status.Delivered.ID(),
		}
		if !lo.Contains(allowedStatus, e.Record.GetString("statusCodeId")) {
			return nil
		}
		shipmentItems, err := dbquery.GetShipmentItemsByShipmentId(app.Dao(), e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("", map[string]validation.Error{
				"shipmentId": validation.NewError("not_found", "Shipment not found"),
			})
		}
		err = app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			var transactionErr error = nil
			for _, shipmentItem := range shipmentItems {
				orderItem, err := dbquery.GetSingleOrderItem(txDao, shipmentItem.OrderItemId)
				if err != nil {
					transactionErr = errors.Join(transactionErr, err)
					continue
				}
				orderItem.ReceivedQty += shipmentItem.Qty
				if orderItem.ReceivedQty > orderItem.OrderedQty {
					return apis.NewBadRequestError("", map[string]validation.Error{
						"qty": validation.NewError("invalid_qty", "Received quantity is greater than ordered quantity"),
					})
				}
				orderItem.MarkAsNotNew()
				if err := txDao.Save(orderItem); err != nil {
					transactionErr = errors.Join(transactionErr, err)
					continue
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
		order, err := dbquery.GetSingleOrder(app.Dao(), e.Record.GetString("orderId"))
		if err != nil {
			return apis.NewNotFoundError("", nil)
		}
		// Check if the order is allowed to be updated to 'Delivered' or 'Partially Delivered'
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
	})
}
