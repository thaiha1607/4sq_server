package utils

import (
	"github.com/thaiha1607/4sq_server/utils/enum/invoice_status"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
	"github.com/thaiha1607/4sq_server/utils/enum/shipment_status"
)

const DeliveryOfficeID = "i2lkp732qf26rvg"

var AllowedTransactionHistoryEntities = []string{
	"internal_orders",
	"invoices",
	"orders",
	"shipments",
}

var OrderStatusCodeTransitions = map[string][]string{
	order_status.Pending.ID(): {
		order_status.Confirmed.ID(),
		order_status.Processing.ID(),
		order_status.Cancelled.ID(),
	},
	order_status.Confirmed.ID(): {
		order_status.Processing.ID(),
		order_status.Cancelled.ID(),
		order_status.OnHold.ID(),
		order_status.AwaitingPayment.ID(),
	},
	order_status.Processing.ID(): {
		order_status.WaitingForAction.ID(),
		order_status.Shipped.ID(),
		order_status.Cancelled.ID(),
		order_status.OnHold.ID(),
		order_status.PartiallyShipped.ID(),
	},
	order_status.WaitingForAction.ID(): {
		order_status.Processing.ID(),
		order_status.Shipped.ID(),
		order_status.Cancelled.ID(),
		order_status.OnHold.ID(),
		order_status.PartiallyShipped.ID(),
	},
	order_status.Shipped.ID(): {
		order_status.Delivered.ID(),
		order_status.Cancelled.ID(),
		order_status.FailedDeliveryAttempt.ID(),
		order_status.PartiallyDelivered.ID(),
	},
	order_status.Delivered.ID(): {},
	order_status.OnHold.ID(): {
		order_status.Processing.ID(),
	},
	order_status.FailedDeliveryAttempt.ID(): {
		order_status.Delivered.ID(),
		order_status.Cancelled.ID(),
	},
	order_status.PartiallyShipped.ID(): {
		order_status.Shipped.ID(),
		order_status.Cancelled.ID(),
		order_status.PartiallyDelivered.ID(),
	},
	order_status.PartiallyDelivered.ID(): {
		order_status.Delivered.ID(),
	},
}

var InternalOrderStatusCodeTransitions = map[string][]string{
	order_status.Pending.ID(): {
		order_status.Processing.ID(),
		order_status.Cancelled.ID(),
	},
	order_status.Processing.ID(): {
		order_status.WaitingForAction.ID(),
		order_status.Shipped.ID(),
		order_status.Cancelled.ID(),
		order_status.OnHold.ID(),
	},
	order_status.WaitingForAction.ID(): {
		order_status.Processing.ID(),
		order_status.Shipped.ID(),
		order_status.Cancelled.ID(),
		order_status.OnHold.ID(),
	},
	order_status.Shipped.ID():   {},
	order_status.Cancelled.ID(): {},
	order_status.OnHold.ID(): {
		order_status.Processing.ID(),
	},
}

var InvoiceStatusCodeTransitions = map[string][]string{
	invoice_status.Draft.ID(): {
		invoice_status.Active.ID(),
		invoice_status.Void.ID(),
		invoice_status.Reserved.ID(),
	},
	invoice_status.Active.ID(): {
		invoice_status.Sent.ID(),
		invoice_status.Void.ID(),
		invoice_status.Reserved.ID(),
	},
	invoice_status.Sent.ID(): {
		invoice_status.Disputed.ID(),
		invoice_status.Overdue.ID(),
		invoice_status.Partial.ID(),
		invoice_status.Paid.ID(),
		invoice_status.Void.ID(),
		invoice_status.Debt.ID(),
		invoice_status.Reserved.ID(),
	},
	invoice_status.Partial.ID(): {
		invoice_status.Paid.ID(),
		invoice_status.Void.ID(),
		invoice_status.Reserved.ID(),
	},
	invoice_status.Paid.ID(): {
		invoice_status.Void.ID(),
		invoice_status.Reserved.ID(),
	},
}

var ShipmentStatusCodeTransitions = map[string][]string{
	shipment_status.Pending.ID(): {
		shipment_status.Processed.ID(),
		shipment_status.Cancelled.ID(),
		shipment_status.OnHold.ID(),
	},
	shipment_status.Processed.ID(): {
		shipment_status.Shipped.ID(),
		shipment_status.Cancelled.ID(),
		shipment_status.OnHold.ID(),
	},
	shipment_status.Shipped.ID(): {
		shipment_status.InTransit.ID(),
		shipment_status.OutForDelivery.ID(),
	},
	shipment_status.InTransit.ID(): {
		shipment_status.OutForDelivery.ID(),
		shipment_status.Delivered.ID(),
		shipment_status.FailedDeliveryAttempt.ID(),
		shipment_status.Returned.ID(),
	},
	shipment_status.OutForDelivery.ID(): {
		shipment_status.Delivered.ID(),
		shipment_status.FailedDeliveryAttempt.ID(),
		shipment_status.Returned.ID(),
	},
	shipment_status.Delivered.ID(): {},
	shipment_status.FailedDeliveryAttempt.ID(): {
		shipment_status.Returned.ID(),
		shipment_status.Cancelled.ID(),
	},
	shipment_status.Returned.ID(): {
		shipment_status.Cancelled.ID(),
	},
	shipment_status.Cancelled.ID(): {},
	shipment_status.OnHold.ID(): {
		shipment_status.Processed.ID(),
	},
}
