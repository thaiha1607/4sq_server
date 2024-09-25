package staff_status

// StaffStatus represents the status of a staff member.
type StaffStatus string

const (
	// Active indicates that the staff member is currently active.
	Active StaffStatus = "active"

	// Inactive indicates that the staff member is currently inactive.
	Inactive StaffStatus = "inactive"

	// Suspended indicates that the staff member is currently suspended.
	Suspended StaffStatus = "suspended"

	// Terminated indicates that the staff member has been terminated.
	Terminated StaffStatus = "terminated"

	// Other indicates that the staff member has a status that does not fit into the other categories.
	Other StaffStatus = "other"
)
