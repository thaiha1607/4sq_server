package order_status

import "strconv"

type OrderStatus int

const (
	// Pending is the status of an order that has been created but not yet confirmed.
	Pending OrderStatus = iota + 1
	// Confirmed is the status of an order that has been confirmed by the customer.
	Confirmed
	// Processing is the status of an order that is being processed by the seller.
	Processing
	// WaitingForAction is the status of an order that is waiting for an action to be taken.
	WaitingForAction
	// Shipped is the status of an order that has been shipped by the seller.
	Shipped
	// Delivered is the status of an order that has been delivered to the customer.
	Delivered
	// Cancelled is the status of an order that has been cancelled by the customer.
	Cancelled
	// Returned is the status of an order that has been returned by the customer.
	Returned
	// OnHold is the status of an order that has been placed on hold by the seller.
	OnHold
	// FailedDeliveryAttempt is the status of an order that has failed delivery attempts.
	FailedDeliveryAttempt
	// Refunded is the status of an order that has been refunded by the seller.
	Refunded
	// PartiallyShipped is the status of an order that has been partially shipped by the seller.
	PartiallyShipped
	// PartiallyDelivered is the status of an order that has been partially delivered to the customer.
	PartiallyDelivered
	// AwaitingPayment is the status of an order that is awaiting payment from the customer.
	AwaitingPayment
)

// ID returns the ID of the order status.
func (s OrderStatus) ID() string {
	r := strconv.Itoa(int(s))
	for len(r) < 15 {
		r = "0" + r
	}
	return r
}
