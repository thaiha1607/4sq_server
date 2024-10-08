package record_hooks

import "github.com/pocketbase/pocketbase"

func RegisterRecordHooks(app *pocketbase.PocketBase) {
	forbidInvalidOrderStatus(app)
	forbidInvalidInternalOrderStatus(app)
	forbidInvalidInvoiceStatus(app)
	forbidInvalidShipmentStatus(app)
	addTransactionHistoryHooks(app)
	assignWarehouseStaff(app)
	updateOrderWhenInternalOrderProcessing(app)
	updateOrderItemWhenInternalOrderItemQtyChanged(app)
	updateOrderItemWhenInternalOrderCancelled(app)
	assignDeliveryStaff(app)
	updateOrderWhenShipmentDelivered(app)
	updateDailyIncomeWhenInvoiceIsPaid(app)
}
