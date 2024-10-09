package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*ShipmentItem)(nil)

type ShipmentItem struct {
	models.BaseModel
	Qty         int64  `db:"qty" json:"qty"`
	RollQty     int64  `db:"rollQty" json:"rollQty,omitempty"`
	ShipmentId  string `db:"shipmentId" json:"shipmentId"`
	OrderItemId string `db:"orderItemId" json:"orderItemId"`
}

func (m *ShipmentItem) TableName() string {
	return "shipment_items"
}

func ShipmentItemQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ShipmentItem{})
}
