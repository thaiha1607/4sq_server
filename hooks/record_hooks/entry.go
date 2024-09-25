package record_hooks

import "github.com/pocketbase/pocketbase"

func RegisterRecordHooks(app *pocketbase.PocketBase) {
	forbidInvalidOrderStatus(app)
	forbidInvalidInternalOrderStatus(app)
	forbidInvalidInvoiceStatus(app)
	forbidInvalidShipmentStatus(app)
	assignWarehouseStaff(app)
	changeOrderStatusToProcessing(app)
	assignDeliveryStaff(app)
	changeOrderStatusToShipped(app)
	changeOrderStatusToDelivered(app)
	addTransactionHistoryHooks(app)
}
