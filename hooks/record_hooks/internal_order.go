package record_hooks

import (
	"math/rand/v2"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/samber/lo"
	"github.com/thaiha1607/4sq_server/custom_models"
	"github.com/thaiha1607/4sq_server/dbquery"
	"github.com/thaiha1607/4sq_server/utils"
	"github.com/thaiha1607/4sq_server/utils/enum/assignment_status"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
	"github.com/thaiha1607/4sq_server/utils/enum/shipment_status"
	"github.com/thaiha1607/4sq_server/utils/enum/staff_role"
)

func forbidInvalidInternalOrderStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("internal_orders").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Pending.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError("invalid_status_code", "When creating an internal order, the status code must be 'Pending'"),
			})
		}
		return nil
	})
	app.OnRecordBeforeUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("internal_orders", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Internal order not found", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			value, ok := utils.InternalOrderStatusCodeTransitions[old.GetString("statusCodeId")]
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

func assignDeliveryStaff(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("internal_orders", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Internal order not found", nil)
		}
		shipmentIdGone := old.GetString("shipmentId") != e.Record.GetString("shipmentId")
		isProcessingToShipped := old.GetString("statusCodeId") == order_status.Processing.ID() &&
			e.Record.GetString("statusCodeId") == order_status.Shipped.ID()
		if !shipmentIdGone && !isProcessingToShipped {
			return nil
		}
		shipmentId := old.GetString("shipmentId")
		internalOrders, err := dbquery.GetInternalOrdersByShipmentId(app.Dao(), shipmentId)
		if err != nil || len(internalOrders) == 0 {
			app.Logger().Error("Failed to get internal orders", "error", err)
			return nil
		}
		internalOrders = lo.DropWhile(internalOrders, func(internalOrder *custom_models.InternalOrder) bool {
			return internalOrder.Id == e.Record.Id
		})
		allShippedInternalOrder := lo.Filter(internalOrders, func(internalOrder *custom_models.InternalOrder, _ int) bool {
			return internalOrder.StatusCodeId == order_status.Shipped.ID()
		})
		if isProcessingToShipped {
			newModel := &custom_models.InternalOrder{
				Type:         e.Record.GetString("type"),
				Note:         e.Record.GetString("note"),
				StatusCodeId: order_status.Shipped.ID(),
				RootOrderId:  e.Record.GetString("rootOrderId"),
				ShipmentId:   shipmentId,
			}
			newModel.SetId(e.Record.Id)
			allShippedInternalOrder = append(allShippedInternalOrder, newModel)
		}
		readyToDeliver := len(lo.Filter(internalOrders, func(internalOrder *custom_models.InternalOrder, _ int) bool {
			return internalOrder.StatusCodeId == order_status.Pending.ID() ||
				internalOrder.StatusCodeId == order_status.Processing.ID() ||
				internalOrder.StatusCodeId == order_status.WaitingForAction.ID()
		})) == 0

		if readyToDeliver {
			app.Logger().Info("Assigning delivery staff", "shipmentId", shipmentId)
			// Update shipment status to Processed
			shipment, err := dbquery.GetSingleShipment(app.Dao(), shipmentId)
			if err != nil {
				app.Logger().Error("Failed to get shipment", "error", err)
				return nil
			}
			shipment.StatusCodeId = shipment_status.Processed.ID()
			shipment.MarkAsNotNew()
			if err := app.Dao().Save(shipment); err != nil {
				app.Logger().Error("Failed to save shipment", "error", err)
				return nil
			}
			// Create shipment items
			shipmentItems := make(map[string]*custom_models.ShipmentItem)
			for _, internalOrder := range allShippedInternalOrder {
				internalOrderItems, err := dbquery.GetInternalOrderItemsByInternalOrderId(app.Dao(), internalOrder.Id)
				if err != nil {
					app.Logger().Error("Failed to get internal order items", "error", err)
					return nil
				}
				for _, internalOrderItem := range internalOrderItems {
					key := shipmentId + internalOrderItem.OrderItemId
					if _, ok := shipmentItems[key]; !ok {
						shipmentItems[key] = &custom_models.ShipmentItem{
							ShipmentId:  shipmentId,
							OrderItemId: internalOrderItem.OrderItemId,
							Qty:         internalOrderItem.Qty,
							RollQty:     internalOrderItem.RollQty,
						}
					} else {
						shipmentItems[key].Qty += internalOrderItem.Qty
						shipmentItems[key].RollQty += internalOrderItem.RollQty
					}
				}
			}
			// Save shipment items
			for _, shipmentItem := range shipmentItems {
				if err := app.Dao().Save(shipmentItem); err != nil {
					app.Logger().Error("Failed to save shipment item", "error", err)
					return nil
				}
			}
			// Update shipment status to Shipped
			fiveDaysLater, _ := types.ParseDateTime(time.Now().AddDate(0, 0, 5))
			shipment.ShipmentDate = types.NowDateTime()
			shipment.DeliveryDate = fiveDaysLater
			shipment.StatusCodeId = shipment_status.Shipped.ID()
			shipment.MarkAsNotNew()
			if err := app.Dao().Save(shipment); err != nil {
				app.Logger().Error("Failed to save shipment", "error", err)
				return nil
			}
			// Get delivery staff
			deliveryStaffs, err := dbquery.GetStaffsByRole(app.Dao(), string(staff_role.Delivery))
			if err != nil {
				app.Logger().Error("Failed to get delivery staffs", "error", err)
				return nil
			}
			staffIdx := rand.IntN(len(deliveryStaffs))
			staffId := deliveryStaffs[staffIdx].Id
			shipmentAssignment := &custom_models.ShipmentAssignment{
				Status:     string(assignment_status.Assigned),
				ShipmentId: shipmentId,
				StaffId:    staffId,
			}
			if err := app.Dao().Save(shipmentAssignment); err != nil {
				app.Logger().Error("Failed to save shipment assignment", "error", err)
				return nil
			}
			app.Logger().Info("Assigned delivery staff", "shipmentId", shipmentId)
		}
		return nil
	})
}
