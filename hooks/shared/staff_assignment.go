package shared

import (
	"log/slog"

	"example.com/4sq_server/custom_models"
	"example.com/4sq_server/dbquery"
	"example.com/4sq_server/utils"
	"example.com/4sq_server/utils/enum/assignment_status"
	"example.com/4sq_server/utils/enum/invoice_type"
	"example.com/4sq_server/utils/enum/order_status"
	"example.com/4sq_server/utils/enum/order_type"
	"example.com/4sq_server/utils/enum/shipment_status"
	"example.com/4sq_server/utils/enum/shipment_type"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/samber/lo"
)

func AssignWarehouseStaff(dao *daos.Dao, logger *slog.Logger, orderRecord *models.Record) error {
	logger.Info("Assigning warehouse staffs", "order_id", orderRecord.Id)
	// Get order items
	orderItems, err := dbquery.GetOrderItemsByOrderId(dao, orderRecord.Id)
	if err != nil || len(orderItems) == 0 {
		logger.Error("Failed to get order items", "error", err)
		return nil
	}
	type internalOrderItem struct {
		Qty         int64
		OrderItemId string
	}
	// Group internal order items by working unit ID
	groupedInternalOrderItems := make(map[string][]internalOrderItem)
	productQuantitySummaries, err := dbquery.GetBatchProductQuantitySummaries(
		dao,
		lo.Map(
			orderItems,
			func(item *custom_models.OrderItem, _ int) string {
				return item.ProductCategoryId
			},
		),
	)
	if err != nil {
		logger.Error("Failed to get product quantity summaries", "error", err)
		return nil
	}
	for _, orderItem := range orderItems {
		// Skip the order item if assigned quantity is equal to ordered quantity
		if orderItem.AssignedQty == orderItem.OrderedQty {
			continue
		}
		productQuantitySummary, ok := lo.Find(
			productQuantitySummaries,
			func(summary *custom_models.ProductQuantitySummary) bool {
				return summary.Id == orderItem.ProductCategoryId
			},
		)
		// We skip the order item if the product quantity summary is not found or the total quantity is 0
		// With the product quantity summary is less than the ordered quantity, we will get all available quantity
		if !ok || productQuantitySummary.TotalQty == 0 {
			continue
		}
		// Get product quantities by category ID. It has been already sorted by priority DESC, qty DESC
		productQuantities, err := dbquery.GetProductQuantitiesByCategoryID(
			dao,
			orderItem.ProductCategoryId,
		)
		if err != nil {
			logger.Error("Failed to get product quantities", "error", err)
			return nil
		}
		orderItemQty := orderItem.OrderedQty - orderItem.AssignedQty
		for _, productQuantity := range productQuantities {
			// If the order item quantity is 0, we break the loop
			if orderItemQty == 0 {
				break
			}
			// Because we sort by priority DESC, qty DESC, we need to check if the quantity is 0
			if productQuantity.Qty == 0 {
				continue
			}
			assignedQty := min(orderItemQty, productQuantity.Qty)
			groupedInternalOrderItems[productQuantity.WorkingUnitID] = append(
				groupedInternalOrderItems[productQuantity.WorkingUnitID],
				internalOrderItem{
					Qty:         assignedQty,
					OrderItemId: orderItem.Id,
				},
			)
			orderItemQty -= assignedQty
		}
		if orderItem.AssignedQty == orderItem.OrderedQty-orderItemQty {
			continue
		}
		orderItem.AssignedQty = orderItem.OrderedQty - orderItemQty
		if err := dao.Save(orderItem); err != nil {
			logger.Error("Failed to save order item", "error", err)
			return nil
		}
	}
	// Check if grouped internal order items are empty
	if len(groupedInternalOrderItems) == 0 {
		logger.Error(
			"Cannot create internal orders because of not enough product quantities",
			"order_id",
			orderRecord.Id,
		)
		return nil
	}

	// To create a shipment, we need to get the first final invoice
	invoices, err := dbquery.GetInvoicesByOrderId(dao, orderRecord.Id)
	if err != nil || len(invoices) == 0 {
		logger.Error("Failed to get invoices", "error", err)
		return nil
	}
	// Only choose first invoice that have type FINAL
	finalInvoices := lo.Filter(invoices,
		func(invoice *custom_models.Invoice, _ int) bool {
			return invoice.Type == string(invoice_type.Final)
		},
	)
	if len(finalInvoices) == 0 {
		logger.Error("No final invoice found", "order_id", orderRecord.Id)
		return nil
	}
	invoice := finalInvoices[0]
	shipentCreation := &custom_models.Shipment{
		Type: string(shipment_type.Outbound),
		// Set shipment date and delivery date to UNIX time 0
		// This is because we don't have the exact date
		// Let it be empty is a big mistake because it will be set to 0001-01-01 00:00:00 +0000 UTC
		// And it will cause a lot of problems
		ShipmentDate: utils.ZeroUnixTime,
		DeliveryDate: utils.ZeroUnixTime,
		Note:         orderRecord.GetString("note"),
		OrderId:      orderRecord.Id,
		InvoiceId:    invoice.Id,
		StatusCodeId: shipment_status.Pending.ID(),
	}
	if err := dao.Save(shipentCreation); err != nil {
		logger.Error("Failed to save shipment", "error", err)
		return nil
	}

	err = dao.RunInTransaction(func(txDao *daos.Dao) error {
		for workingUnitID, internalOrderItems := range groupedInternalOrderItems {
			internalOrderCreation := &custom_models.InternalOrder{
				Type:             string(order_type.Transfer),
				Note:             orderRecord.GetString("note"),
				StatusCodeId:     order_status.Pending.ID(),
				RootOrderId:      orderRecord.Id,
				ShipmentId:       shipentCreation.Id,
				SrcWorkingUnitId: workingUnitID,
				DstWorkingUnitId: utils.DeliveryOfficeID,
			}
			if err := txDao.Save(internalOrderCreation); err != nil {
				logger.Error("Failed to save internal order", "error", err)
				return nil
			}
			internalOrderItemCreations := lo.Map(
				internalOrderItems,
				func(item internalOrderItem, _ int) *custom_models.InternalOrderItem {
					return &custom_models.InternalOrderItem{
						Qty:             item.Qty,
						RollQty:         0,
						InternalOrderId: internalOrderCreation.Id,
						OrderItemId:     item.OrderItemId,
					}
				},
			)
			for _, item := range internalOrderItemCreations {
				if err := txDao.Save(item); err != nil {
					logger.Error("Failed to save internal order item", "error", err)
					return err
				}
			}
			warehouseStaffs, err := dbquery.GetStaffsByWorkingUnitId(txDao, workingUnitID)
			if err != nil || len(warehouseStaffs) == 0 {
				logger.Error("Failed to get warehouse staffs", "error", err)
				return nil
			}
			if len(warehouseStaffs) > 1 {
				warehouseStaffs = lo.Shuffle(warehouseStaffs)
			}
			// Get only one staff
			selectedStaff := warehouseStaffs[0]
			warehouseAssignment := &custom_models.WarehouseAssignment{
				Status:          string(assignment_status.Assigned),
				StaffId:         selectedStaff.Id,
				InternalOrderId: internalOrderCreation.Id,
			}
			if err := txDao.Save(warehouseAssignment); err != nil {
				logger.Error("Failed to save warehouse assignment", "error", err)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	logger.Info("Warehouse staffs assigned", "order_id", orderRecord.Id)
	return nil
}
