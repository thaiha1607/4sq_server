package shipment_type

// ShipmentType represents the type of shipment.
type ShipmentType string

const (
	// Outbound represents an outbound shipment.
	Outbound ShipmentType = "outbound"
	// Inbound represents an inbound shipment.
	Inbound ShipmentType = "inbound"
	// Transfer represents a transfer shipment.
	Transfer ShipmentType = "transfer"
	// Return represents a return shipment.
	Return ShipmentType = "return"
	// Exchange represents an exchange shipment.
	Exchange ShipmentType = "exchange"
	// Other represents any other type of shipment.
	Other ShipmentType = "other"
)
