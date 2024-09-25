package staff_role

// StaffRole represents the role of a staff member in the organization.
type StaffRole string

const (
	// Salesperson represents a staff member who is responsible for sales.
	Salesperson StaffRole = "salesperson"
	// Warehouse represents a staff member who works in the warehouse.
	Warehouse StaffRole = "warehouse"
	// Delivery represents a staff member who is responsible for deliveries.
	Delivery StaffRole = "delivery"
	// Other represents a staff member with a role that does not fit into the other categories.
	Other StaffRole = "other"
)
