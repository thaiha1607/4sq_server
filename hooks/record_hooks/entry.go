package record_hooks

import "github.com/pocketbase/pocketbase"

func RegisterRecordHooks(app *pocketbase.PocketBase) {
	forbidInvalidOrderStatus(app)
	forbidInvalidInternalOrderStatus(app)
	forbidInvalidInvoiceStatus(app)
	forbidInvalidShipmentStatus(app)
	updateOrderItemWhenInternalOrderItemQtyChanged(app)
	updateOrderItemWhenInternalOrderCancelled(app)
	updateOrderWhenShipmentDelivered(app)
	updateDailyIncomeWhenInvoiceIsPaid(app)
	assignWarehouseStaff(app)
	assignDeliveryStaff(app)
	addTransactionHistoryHooks(app)
}
