package shipment_status

import "strconv"

type ShipmentStatus int

const (
	// Pending is the status of a shipment that has been created but not yet processed.
	Pending ShipmentStatus = iota + 1
	// Processed is the status of a shipment that has been processed by the seller.
	Processed
	// Shipped is the status of a shipment that has been shipped by the seller.
	Shipped
	// InTransit is the status of a shipment that is in transit to the customer.
	InTransit
	// OutForDelivery is the status of a shipment that is out for delivery to the customer.
	OutForDelivery
	// Delivered is the status of a shipment that has been delivered to the customer.
	Delivered
	// FailedDeliveryAttempt is the status of a shipment that has failed delivery attempts.
	FailedDeliveryAttempt
	// Returned is the status of a shipment that has been returned by the customer.
	Returned
	// Cancelled is the status of a shipment that has been cancelled by the customer.
	Cancelled
	// OnHold is the status of a shipment that has been placed on hold by the seller.
	OnHold
)

// ID returns the ID of the shipment status.
func (s ShipmentStatus) ID() string {
	r := strconv.Itoa(int(s))
	for len(r) < 15 {
		r = "0" + r
	}
	return r
}
