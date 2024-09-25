package assignment_status

// AssignmentStatus represents the status of an assignment.
type AssignmentStatus string

const (
	// Pending indicates that the assignment is pending.
	Pending AssignmentStatus = "pending"
	// Assigned indicates that the assignment has been assigned.
	Assigned AssignmentStatus = "assigned"
	// InProgress indicates that the assignment is in progress.
	InProgress AssignmentStatus = "in_progress"
	// Completed indicates that the assignment has been completed.
	Completed AssignmentStatus = "completed"
	// Cancelled indicates that the assignment has been cancelled.
	Cancelled AssignmentStatus = "cancelled"
	// Failed indicates that the assignment has failed.
	Failed AssignmentStatus = "failed"
	// Other indicates that the assignment has a different status.
	Other AssignmentStatus = "other"
)
