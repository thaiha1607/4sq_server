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
	"github.com/thaiha1607/4sq_server/utils/enum/assignment_status"
	"github.com/thaiha1607/4sq_server/utils/enum/invoice_type"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
	"github.com/thaiha1607/4sq_server/utils/enum/order_type"
	"github.com/thaiha1607/4sq_server/utils/enum/shipment_status"
	"github.com/thaiha1607/4sq_server/utils/enum/shipment_type"
	"github.com/thaiha1607/4sq_server/utils/enum/staff_role"
)

func forbidInvalidOrderStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("orders").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Pending.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError("invalid_status_code", "When creating an order, the status code must be 'Pending'"),
			})
		}
		return nil
	})
	app.OnRecordBeforeUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("orders", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Order not found", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			value, ok := utils.OrderStatusCodeTransitions[old.GetString("statusCodeId")]
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

func assignWarehouseStaff(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("orders", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Order not found", nil)
		}
		// Only assign warehouse staffs when the order is changed from Pending to Confirmed
		if old.GetString("statusCodeId") == order_status.Pending.ID() &&
			e.Record.GetString("statusCodeId") == order_status.Confirmed.ID() {
			app.Logger().Info("Assigning warehouse staffs", "order_id", e.Record.Id)
			// Get order items
			orderItems, err := dbquery.GetOrderItemsByOrderId(app.Dao(), e.Record.Id)
			if err != nil {
				app.Logger().Error("Failed to get order items", "error", err)
				return nil
			}
			// Get warehouse staffs
			warehouseStaffs, err := dbquery.GetStaffsByRole(app.Dao(), string(staff_role.Warehouse))
			if err != nil {
				app.Logger().Error("Failed to get warehouse staffs", "error", err)
				return nil
			}
			warehouseStaffs = lo.Shuffle(warehouseStaffs)
			requiredStaffs := min(len(orderItems)/5+1, len(warehouseStaffs))
			selectedStaff := warehouseStaffs[:requiredStaffs]
			itemLenPerStaff := len(orderItems) / requiredStaffs
			// Split order items into numStaffs parts. Leftover items will be appended to the last staff
			groupedOrderItems := lo.Chunk(orderItems, itemLenPerStaff)
			// Append leftover items to the last staff if needed
			if len(groupedOrderItems) > requiredStaffs {
				groupedOrderItems[requiredStaffs-1] = append(groupedOrderItems[requiredStaffs-1], groupedOrderItems[requiredStaffs]...)
				groupedOrderItems = groupedOrderItems[:requiredStaffs]
			}
			// Get pro forma invoice
			invoices, err := dbquery.GetInvoicesByOrderId(app.Dao(), e.Record.Id)
			if err != nil {
				app.Logger().Error("Failed to get invoices", "error", err)
				return nil
			}
			invoice, ok := lo.First(
				lo.Filter(invoices, func(item *custom_models.Invoice, _ int) bool {
					return item.Type == string(invoice_type.ProForma)
				}),
			)
			if !ok {
				app.Logger().Error("Pro forma invoice not found", "order_id", e.Record.Id)
				return nil
			}
			shipment := &custom_models.Shipment{
				Type:         string(shipment_type.Outbound),
				OrderId:      e.Record.Id,
				InvoiceId:    invoice.Id,
				StatusCodeId: shipment_status.Pending.ID(),
			}
			if err := app.Dao().Save(shipment); err != nil {
				app.Logger().Error("Failed to save shipment", "error", err)
				return nil
			}
			staffIdx := 0
			for range requiredStaffs {
				// Create internal order
				internal_order := &custom_models.InternalOrder{
					Type:         string(order_type.Transfer),
					StatusCodeId: order_status.Pending.ID(),
					RootOrderId:  e.Record.Id,
					ShipmentId:   shipment.Id,
				}
				if err := app.Dao().Save(internal_order); err != nil {
					app.Logger().Error("Failed to save internal order", "error", err)
					return nil
				}
				// Create warehouse staff assignment
				assignment := &custom_models.WarehouseAssignment{
					Status:           string(assignment_status.Assigned),
					StaffId:          selectedStaff[staffIdx].Id,
					SrcWorkingUnitId: selectedStaff[staffIdx].WorkingUnitId,
					DstWorkingUnitId: utils.DeliveryOfficeID,
					InternalOrderId:  internal_order.Id,
				}
				if err := app.Dao().Save(assignment); err != nil {
					app.Logger().Error("Failed to save warehouse assignment", "error", err)
					return nil
				}
				// Create internal order items
				for _, item := range groupedOrderItems[staffIdx] {
					internal_order_item := &custom_models.InternalOrderItem{
						Qty:             item.OrderedQty,
						InternalOrderId: internal_order.Id,
						OrderItemId:     item.Id,
					}
					if err := app.Dao().Save(internal_order_item); err != nil {
						app.Logger().Error("Failed to save internal order item", "error", err)
						return nil
					}
				}
				staffIdx++
			}
			app.Logger().Info("Warehouse staffs assigned", "order_id", e.Record.Id)
		}
		return nil
	})
}
