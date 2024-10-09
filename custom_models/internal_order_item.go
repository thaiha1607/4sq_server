package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*InternalOrderItem)(nil)

type InternalOrderItem struct {
	models.BaseModel
	Qty             int64  `db:"qty" json:"qty,omitempty"`
	RollQty         int64  `db:"rollQty" json:"rollQty,omitempty"`
	Note            string `db:"note" json:"note,omitempty"`
	InternalOrderId string `db:"internalOrderId" json:"internalOrderId"`
	OrderItemId     string `db:"orderItemId" json:"orderItemId"`
}

func (m *InternalOrderItem) TableName() string {
	return "internal_order_items"
}

func InternalOrderItemQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&InternalOrderItem{})
}
