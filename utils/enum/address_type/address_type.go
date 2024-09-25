package address_type

// AddressType represents the type of an address.
type AddressType string

const (
	// Home is the type of a home address.
	Home AddressType = "home"
	// Work is the type of a work address.
	Work AddressType = "work"
	// Billing is the type of a billing address.
	Billing AddressType = "billing"
	// Shipping is the type of a shipping address.
	Shipping AddressType = "shipping"
	// Other is the type of an other address.
	Other AddressType = "other"
)
