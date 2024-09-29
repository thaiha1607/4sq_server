package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*ProductQuantityHistory)(nil)

type ProductQuantityHistory struct {
	models.BaseModel
	CategoryId     string  `db:"categoryId" json:"categoryId"`
	AmountOfChange float64 `db:"amountOfChange" json:"amountOfChange"`
}

func (m *ProductQuantityHistory) TableName() string {
	return "product_quantity_history"
}

func ProductQuantityHistoryQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ProductQuantityHistory{})
}
