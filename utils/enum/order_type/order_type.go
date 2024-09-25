package order_type

// OrderType represents the type of an order.
type OrderType string

const (
	// Sale represents a sale order.
	Sale OrderType = "sale"

	// Return represents a return order.
	Return OrderType = "return"

	// Exchange represents an exchange order.
	Exchange OrderType = "exchange"

	// Transfer represents a transfer order.
	Transfer OrderType = "transfer"

	// Other represents an order type that does not fit into any of the predefined categories.
	Other OrderType = "other"
)
