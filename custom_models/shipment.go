package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ models.Model = (*Shipment)(nil)

type Shipment struct {
	models.BaseModel
	Type         string         `db:"type" json:"type"`
	ShipmentDate types.DateTime `db:"shipmentDate" json:"shipmentDate,omitempty"`
	DeliveryDate types.DateTime `db:"deliveryDate" json:"deliveryDate,omitempty"`
	Note         string         `db:"note" json:"note,omitempty"`
	OrderId      string         `db:"orderId" json:"orderId"`
	InvoiceId    string         `db:"invoiceId" json:"invoiceId"`
	StatusCodeId string         `db:"statusCodeId" json:"statusCodeId"`
}

func (m *Shipment) TableName() string {
	return "shipments"
}

func ShipmentQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Shipment{})
}
