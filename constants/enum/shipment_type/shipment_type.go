package shipment_type

// ShipmentType represents the type of shipment.
type ShipmentType string

const (
	// Outbound represents an outbound shipment.
	Outbound ShipmentType = "Outbound"
	// Inbound represents an inbound shipment.
	Inbound ShipmentType = "Inbound"
	// Transfer represents a transfer shipment.
	Transfer ShipmentType = "Transfer"
	// Return represents a return shipment.
	Return ShipmentType = "Return"
	// Exchange represents an exchange shipment.
	Exchange ShipmentType = "Exchange"
	// Other represents any other type of shipment.
	Other ShipmentType = "Other"
)
