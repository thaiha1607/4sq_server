package record_hooks

import "github.com/pocketbase/pocketbase"

func RegisterRecordHooks(app *pocketbase.PocketBase) {
	forbidInvalidOrderStatus(app)
	forbidInvalidInternalOrderStatus(app)
	forbidInvalidInvoiceStatus(app)
	forbidInvalidShipmentStatus(app)
	updateOrderWhenShipmentDelivered(app)
	updateDailyIncomeWhenInvoiceIsPaid(app)
	updateOrderItemWhenInternalOrderCancelled(app)
	assignWarehouseStaff(app)
	assignDeliveryStaff(app)
	addTransactionHistoryHooks(app)
}
