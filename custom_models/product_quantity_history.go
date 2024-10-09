package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
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
