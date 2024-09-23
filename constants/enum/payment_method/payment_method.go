package payment_method

// PaymentMethod represents the method of payment for a transaction.
type PaymentMethod string

const (
	// Cash represents a payment method using physical currency.
	Cash PaymentMethod = "cash"
	// EFT represents a payment method using electronic funds transfer.
	EFT PaymentMethod = "eft"
	// GiftCard represents a payment method using a gift card.
	GiftCard PaymentMethod = "gift_card"
	// CreditCard represents a payment method using a credit card.
	CreditCard PaymentMethod = "credit_card"
	// DebitCard represents a payment method using a debit card.
	DebitCard PaymentMethod = "debit_card"
	// PrepaidCard represents a payment method using a prepaid card.
	PrepaidCard PaymentMethod = "prepaid_card"
	// Check represents a payment method using a check.
	Check PaymentMethod = "check"
	// Other represents a payment method using another form of payment.
	Other PaymentMethod = "other"
)
