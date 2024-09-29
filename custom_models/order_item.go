package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*OrderItem)(nil)

type OrderItem struct {
	models.BaseModel
	OrderedQty        int64   `db:"orderedQty" json:"orderedQty"`
	ReceivedQty       int64   `db:"receivedQty" json:"receivedQty,omitempty"`
	ShippedQty        int64   `db:"shippedQty" json:"shippedQty,omitempty"`
	AssignedQty       int64   `db:"assignedQty" json:"assignedQty,omitempty"`
	UnitPrice         float64 `db:"unitPrice" json:"unitPrice"`
	Note              string  `db:"note" json:"note,omitempty"`
	OrderId           string  `db:"orderId" json:"orderId"`
	ProductCategoryId string  `db:"productCategoryId" json:"productCategoryId"`
}

func (m *OrderItem) TableName() string {
	return "order_items"
}

func OrderItemQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&OrderItem{})
}
