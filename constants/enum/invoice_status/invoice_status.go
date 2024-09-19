package invoice_status

import "strconv"

// InvoiceStatus represents the status of an invoice.
type InvoiceStatus int

const (
	// Draft is the status of an invoice that has been created but not yet sent to the customer.
	Draft InvoiceStatus = iota + 1
	// Active is the status of an invoice that has been sent to the customer and is awaiting payment.
	Active
	// Sent is the status of an invoice that has been sent to the customer.
	Sent
	// Disputed is the status of an invoice that has been disputed by the customer.
	Disputed
	// Overdue is the status of an invoice that has not been paid by the due date.
	Overdue
	// Partial is the status of an invoice that has been partially paid by the customer.
	Partial
	// Paid is the status of an invoice that has been paid by the customer.
	Paid
	// Void is the status of an invoice that has been voided by the user.
	Void
	// BadDebt is the status of an invoice that has been deemed uncollectible by the user.
	Debt
	// Reserved is the status of an invoice that has been reserved for future use.
	Reserved
)

// ID returns the ID of the invoice status.
func (s InvoiceStatus) ID() string {
	r := strconv.Itoa(int(s))
	for len(r) < 15 {
		r = "0" + r
	}
	return r
}
