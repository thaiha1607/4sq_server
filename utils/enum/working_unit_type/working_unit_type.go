package working_unit_type

// WorkingUnitType represents the type of working unit.
type WorkingUnitType string

const (
	// Warehouse represents a warehouse working unit.
	Warehouse WorkingUnitType = "warehouse"

	// Office represents an office working unit.
	Office WorkingUnitType = "office"

	// Delivery represents a delivery working unit.
	Delivery WorkingUnitType = "delivery"

	// Other represents any other type of working unit.
	Other WorkingUnitType = "other"
)
