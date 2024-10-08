package record_hooks

import (
	"errors"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
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
		if e.Record.GetString("statusCodeId") == order_status.Pending.ID() {
			return nil
		}
		return apis.NewBadRequestError("", map[string]validation.Error{
			"statusCodeId": validation.NewError("invalid_status_code", "When creating an internal order, the status code must be 'Pending'"),
		})
	})
	app.OnRecordBeforeUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("internal_orders", e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		if old.GetString("statusCodeId") == e.Record.GetString("statusCodeId") {
			return nil
		}
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
		return nil
	})
}

func updateOrderWhenInternalOrderProcessing(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Processing.ID() {
			return nil
		}
		err := app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			order, err := dbquery.GetSingleOrder(txDao, e.Record.GetString("rootOrderId"))
			if err != nil {
				return apis.NewNotFoundError("", map[string]validation.Error{
					"rootOrderId": validation.NewError("not_found", "Order not found"),
				})
			}
			if order.StatusCodeId != order_status.Confirmed.ID() {
				return nil
			}
			order.StatusCodeId = order_status.Processing.ID()
			order.MarkAsNotNew()
			if err := txDao.Save(order); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
}

func updateOrderItemWhenInternalOrderCancelled(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Cancelled.ID() {
			return nil
		}
		app.Logger().Info("Updating order item when internal order cancelled", "internalOrderId", e.Record.Id)
		err := app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			internalOrderItems, err := dbquery.GetInternalOrderItemsByInternalOrderId(txDao, e.Record.Id)
			if err != nil {
				return apis.NewNotFoundError("", map[string]validation.Error{
					"internalOrderId": validation.NewError("not_found", "Internal order items not found"),
				})
			}
			var transactionErr error = nil
			for _, internalOrderItem := range internalOrderItems {
				orderItem, err := dbquery.GetSingleOrderItem(txDao, internalOrderItem.OrderItemId)
				if err != nil {
					transactionErr = errors.Join(transactionErr, err)
					continue
				}
				orderItem.AssignedQty -= internalOrderItem.Qty
				if orderItem.AssignedQty < 0 {
					orderItem.AssignedQty = 0
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
			app.Logger().Error("Failed to update order item when internal order cancelled", "error", err)
			return nil
		}
		app.Logger().Info("Updated order item when internal order cancelled", "internalOrderId", e.Record.Id)
		return nil
	})
}

func updateOrderItemWhenInternalOrderItemQtyChanged(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("internal_order_items").Add(func(e *core.RecordUpdateEvent) error {
		old, err := dbquery.GetSingleInternalOrderItem(app.Dao(), e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		oldQty := old.Qty
		newQty := int64(e.Record.GetInt("qty"))
		if oldQty == newQty {
			return nil
		}

		orderItem, err := dbquery.GetSingleOrderItem(app.Dao(), e.Record.GetString("orderItemId"))
		if err != nil {
			return apis.NewNotFoundError("", nil)
		}
		amountOfChange := newQty - oldQty
		if newQty < oldQty {
			orderItem.AssignedQty += amountOfChange
			if orderItem.AssignedQty < 0 {
				orderItem.AssignedQty = 0
			}
		} else {
			orderItem.AssignedQty += amountOfChange
			if orderItem.AssignedQty > orderItem.OrderedQty {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"qty": validation.NewError("invalid_qty", "Assigned quantity exceeds ordered quantity"),
				})
			}
		}

		orderItem.MarkAsNotNew()
		if err := app.Dao().Save(orderItem); err != nil {
			return err
		}

		return nil
	})
}

func updateOrderWhenInternalOrderShipped(dao *daos.Dao, internalOrderId string, orderId string) error {
	internalOrderItems, err := dbquery.GetInternalOrderItemsByInternalOrderId(dao, internalOrderId)
	if err != nil {
		return apis.NewNotFoundError("", map[string]validation.Error{
			"internalOrderId": validation.NewError("not_found", "Internal order items not found"),
		})
	}
	for _, internalOrderItem := range internalOrderItems {
		orderItem, err := dbquery.GetSingleOrderItem(dao, internalOrderItem.OrderItemId)
		if err != nil {
			return apis.NewNotFoundError("", nil)
		}
		orderItem.ShippedQty += internalOrderItem.Qty
		if orderItem.ShippedQty > orderItem.OrderedQty {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"qty": validation.NewError("invalid_qty", "Shipped quantity exceeds ordered quantity"),
			})
		}
		orderItem.MarkAsNotNew()
		if err := dao.Save(orderItem); err != nil {
			return err
		}
	}
	order, err := dbquery.GetSingleOrder(dao, orderId)
	if err != nil {
		return apis.NewNotFoundError("", nil)
	}
	statusAllowed := []string{
		order_status.Processing.ID(),
		order_status.WaitingForAction.ID(),
		order_status.PartiallyShipped.ID(),
	}
	if !lo.Contains(statusAllowed, order.StatusCodeId) {
		return nil
	}
	// Get all order items
	orderItems, err := dbquery.GetOrderItemsByOrderId(dao, orderId)
	if err != nil {
		return apis.NewNotFoundError("", nil)
	}
	allShipped := lo.EveryBy(orderItems, func(orderItem *custom_models.OrderItem) bool {
		return orderItem.ShippedQty == orderItem.OrderedQty
	})
	if allShipped {
		order.StatusCodeId = order_status.Shipped.ID()
	} else {
		if order.StatusCodeId == order_status.PartiallyShipped.ID() {
			return nil
		}
		order.StatusCodeId = order_status.PartiallyShipped.ID()
	}
	order.MarkAsNotNew()
	if err := dao.Save(order); err != nil {
		return err
	}
	return nil
}

func updateQuantityInWarehouse(
	dao *daos.Dao,
	logger *slog.Logger,
	internalOrderId string,
	workingUnitId string,
	delta int64,
) error {
	internalOrderItems, err := dbquery.GetInternalOrderItemsByInternalOrderId(dao, internalOrderId)
	if err != nil {
		return apis.NewNotFoundError("", map[string]validation.Error{
			"internalOrderId": validation.NewError("not_found", "Internal order items not found"),
		})
	}
	for _, internalOrderItem := range internalOrderItems {
		orderItem, err := dbquery.GetSingleOrderItem(dao, internalOrderItem.OrderItemId)
		if err != nil {
			return apis.NewNotFoundError("", nil)
		}
		productQuantity, err := dbquery.GetSingleProductQuantityByCategoryIDAndWorkingUnitID(
			dao,
			orderItem.ProductCategoryId,
			workingUnitId,
		)
		if err != nil {
			if internalOrderItem.Qty == 0 {
				logger.Warn("Failed to get product quantity", "data", map[string]string{
					"categoryId":    orderItem.ProductCategoryId,
					"workingUnitId": workingUnitId,
				})
				continue
			} else {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"qty": validation.NewError("invalid_qty", "Not enough quantity in warehouse"),
				})
			}
		}
		productQuantity.Qty += (delta * internalOrderItem.Qty)
		if productQuantity.Qty < 0 {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"qty": validation.NewError("invalid_qty", "Not enough quantity in warehouse"),
			})
		}
		productQuantity.MarkAsNotNew()
		if err := dao.Save(productQuantity); err != nil {
			return err
		}
		// Update product quantity history
		productQuantityHistory := &custom_models.ProductQuantityHistory{
			CategoryId:     orderItem.ProductCategoryId,
			AmountOfChange: float64(delta * internalOrderItem.Qty),
		}
		if err := dao.Save(productQuantityHistory); err != nil {
			return err
		}
	}
	return nil
}

func assignDeliveryStaff(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("internal_orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("internal_orders", e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		shipmentIdGone := old.GetString("shipmentId") != e.Record.GetString("shipmentId")
		isProcessingToShipped := old.GetString("statusCodeId") == order_status.Processing.ID() &&
			e.Record.GetString("statusCodeId") == order_status.Shipped.ID()
		if !shipmentIdGone && !isProcessingToShipped {
			return nil
		}
		if isProcessingToShipped {
			if err := app.Dao().RunInTransaction(
				func(txDao *daos.Dao) error {
					return updateQuantityInWarehouse(
						txDao,
						app.Logger(),
						e.Record.Id,
						e.Record.GetString("srcWorkingUnitId"),
						-1,
					)
				},
			); err != nil {
				return err
			}
			if err := app.Dao().RunInTransaction(
				func(txDao *daos.Dao) error {
					return updateOrderWhenInternalOrderShipped(
						txDao,
						e.Record.Id,
						e.Record.GetString("rootOrderId"),
					)
				},
			); err != nil {
				return err
			}
			shipment, err := dbquery.GetSingleShipment(app.Dao(), old.GetString("shipmentId"))
			if err != nil {
				app.Logger().Error("Failed to get shipment", "error", err)
				return nil
			}
			if shipment.StatusCodeId == shipment_status.Pending.ID() {
				shipment.StatusCodeId = shipment_status.Processed.ID()
				shipment.MarkAsNotNew()
				if err := app.Dao().Save(shipment); err != nil {
					app.Logger().Error("Failed to save shipment", "error", err)
					return nil
				}
			}

		}
		shipmentId := old.GetString("shipmentId")
		internalOrderRecords, err := dbquery.GetInternalOrdersByShipmentId(app.Dao(), shipmentId)
		if err != nil || len(internalOrderRecords) == 0 {
			app.Logger().Error("Failed to get internal orders", "error", err)
			return nil
		}
		internalOrders := lo.Reject(internalOrderRecords, func(internalOrder *custom_models.InternalOrder, _ int) bool {
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
		readyToDeliver := lo.EveryBy(internalOrders, func(internalOrder *custom_models.InternalOrder) bool {
			notReadyStatus := []string{
				order_status.Pending.ID(),
				order_status.Processing.ID(),
				order_status.WaitingForAction.ID(),
			}
			return !lo.Contains(notReadyStatus, internalOrder.StatusCodeId)
		})
		allCancelled := lo.EveryBy(internalOrders, func(internalOrder *custom_models.InternalOrder) bool {
			return internalOrder.StatusCodeId == order_status.Cancelled.ID()
		})
		if !readyToDeliver {
			return nil
		}
		if allCancelled {
			app.Logger().Info("All internal orders are cancelled", "shipmentId", shipmentId)
			shipment, err := dbquery.GetSingleShipment(app.Dao(), shipmentId)
			if err != nil {
				app.Logger().Error("Failed to get shipment", "error", err)
				return nil
			}
			shipment.StatusCodeId = shipment_status.Cancelled.ID()
			shipment.MarkAsNotNew()
			if err := app.Dao().Save(shipment); err != nil {
				app.Logger().Error("Failed to save shipment", "error", err)
				return nil
			}
			return nil
		}

		app.Logger().Info("Assigning delivery staff", "shipmentId", shipmentId)
		err = app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			// Create shipment items
			shipmentItems := make(map[string]*custom_models.ShipmentItem)
			for _, internalOrder := range allShippedInternalOrder {
				internalOrderItems, err := dbquery.GetInternalOrderItemsByInternalOrderId(txDao, internalOrder.Id)
				if err != nil {
					app.Logger().Error("Failed to get internal order items", "error", err)
					return nil
				}
				for _, internalOrderItem := range internalOrderItems {
					if internalOrderItem.Qty == 0 {
						continue
					}
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
				if err := txDao.Save(shipmentItem); err != nil {
					app.Logger().Error("Failed to save shipment item", "error", err)
					return nil
				}
			}
			// Update shipment status to Shipped
			fiveDaysLater, _ := types.ParseDateTime(time.Now().AddDate(0, 0, 5))
			shipment, err := dbquery.GetSingleShipment(txDao, shipmentId)
			if err != nil {
				app.Logger().Error("Failed to get shipment", "error", err)
				return nil
			}
			shipment.ShipmentDate = types.NowDateTime()
			shipment.DeliveryDate = fiveDaysLater
			shipment.StatusCodeId = shipment_status.Shipped.ID()
			shipment.MarkAsNotNew()
			if err := txDao.Save(shipment); err != nil {
				app.Logger().Error("Failed to save shipment", "error", err)
				return nil
			}
			// Get delivery staff
			deliveryStaffs, err := dbquery.GetStaffsByRole(txDao, string(staff_role.Delivery))
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
			if err := txDao.Save(shipmentAssignment); err != nil {
				app.Logger().Error("Failed to save shipment assignment", "error", err)
				return nil
			}
			return nil
		})
		if err != nil {
			return err
		}
		app.Logger().Info("Assigned delivery staff", "shipmentId", shipmentId)
		return nil
	})
}
