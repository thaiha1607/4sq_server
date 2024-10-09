package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*InternalOrder)(nil)

type InternalOrder struct {
	models.BaseModel
	Type             string `db:"type" json:"type"`
	Note             string `db:"note" json:"note,omitempty"`
	StatusCodeId     string `db:"statusCodeId" json:"statusCodeId"`
	RootOrderId      string `db:"rootOrderId" json:"rootOrderId"`
	ShipmentId       string `db:"shipmentId" json:"shipmentId,omitempty"`
	SrcWorkingUnitId string `db:"srcWorkingUnitId" json:"srcWorkingUnitId"`
	DstWorkingUnitId string `db:"dstWorkingUnitId" json:"dstWorkingUnitId,omitempty"`
}

func (m *InternalOrder) TableName() string {
	return "internal_orders"
}

func InternalOrderQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&InternalOrder{})
}
