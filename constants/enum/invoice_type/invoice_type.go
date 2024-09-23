package invoice_type

// InvoiceType represents the type of an invoice.
type InvoiceType string

const (
	// ProForma represents a preliminary invoice sent to a buyer in advance of a shipment or delivery of goods.
	// It typically contains a description of the items, the total amount due, and other important details,
	// but it is not a demand for payment.
	ProForma InvoiceType = "pro_forma"

	// Regular represents a standard invoice issued for goods or services rendered.
	Regular InvoiceType = "regular"

	// PastDue represents an invoice that has not been paid by the due date.
	PastDue InvoiceType = "past_due"

	// Retainer represents an invoice issued for a retainer fee, which is an advance payment for services.
	Retainer InvoiceType = "retainer"

	// Interim represents an invoice issued for partial payment before the final invoice.
	Interim InvoiceType = "interim"

	// Timesheet represents an invoice based on hours worked, typically used for freelance or contract work.
	Timesheet InvoiceType = "timesheet"

	// Final represents the last invoice issued for a project, indicating that no further invoices will be issued.
	Final InvoiceType = "final"

	// Credit represents an invoice issued to correct or reduce the amount of a previous invoice.
	Credit InvoiceType = "credit"

	// Debit represents an invoice issued to increase the amount of a previous invoice.
	Debit InvoiceType = "debit"

	// Mixed represents an invoice that includes both credit and debit items.
	Mixed InvoiceType = "mixed"

	// Commercial represents an invoice issued for commercial transactions, typically between businesses.
	Commercial InvoiceType = "commercial"

	// Recurring represents an invoice issued on a regular basis, such as monthly or annually.
	Recurring InvoiceType = "recurring"

	// Other represents an invoice type that does not fit into any of the predefined categories.
	Other InvoiceType = "other"
)
