package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*InternalOrderItem)(nil)

type InternalOrderItem struct {
	models.BaseModel
	Qty             int64  `db:"qty" json:"qty,omitempty"`
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
