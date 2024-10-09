package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

// Get all invoices by order id, sorted descending by created_at
func GetInvoicesByOrderId(dao *daos.Dao, orderId string) ([]*custom_models.Invoice, error) {
	var invoices []*custom_models.Invoice
	err := custom_models.
		InvoiceQuery(dao).
		Where(dbx.HashExp{"orderId": orderId}).
		OrderBy("created DESC").
		All(&invoices)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
