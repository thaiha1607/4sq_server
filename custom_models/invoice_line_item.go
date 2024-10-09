package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*InvoiceLineItem)(nil)

type InvoiceLineItem struct {
	models.BaseModel
	InvoiceId   string  `db:"invoiceId" json:"invoiceId"`
	OrderItemId string  `db:"orderItemId" json:"orderItemId"`
	UnitPrice   float64 `db:"unitPrice" json:"unitPrice"`
	Note        string  `db:"note" json:"note,omitempty"`
}

func (m *InvoiceLineItem) TableName() string {
	return "invoice_line_items"
}

func InvoiceLineItemQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&InvoiceLineItem{})
}
