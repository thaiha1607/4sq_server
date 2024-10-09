package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*Invoice)(nil)

type Invoice struct {
	models.BaseModel
	TotalAmount   float64 `db:"totalAmount" json:"totalAmount"`
	PaidAmount    float64 `db:"paidAmount" json:"paidAmount,omitempty"`
	Type          string  `db:"type" json:"type"`
	PaymentMethod string  `db:"paymentMethod" json:"paymentMethod"`
	Note          string  `db:"note" json:"note,omitempty"`
	OrderId       string  `db:"orderId" json:"orderId"`
	StatusCodeId  string  `db:"statusCodeId" json:"statusCodeId"`
	RootInvoiceId string  `db:"rootInvoiceId" json:"rootInvoiceId,omitempty"`
	OtherInfo     string  `db:"otherInfo" json:"otherInfo,omitempty"`
}

func (m *Invoice) TableName() string {
	return "invoices"
}

func InvoiceQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Invoice{})
}
