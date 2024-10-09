package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*ProductQuantity)(nil)

type ProductQuantity struct {
	models.BaseModel
	Priority      int64  `db:"priority" json:"priority"`
	Qty           int64  `db:"qty" json:"qty,omitempty"`
	CategoryID    string `db:"categoryId" json:"categoryId"`
	WorkingUnitID string `db:"workingUnitId" json:"workingUnitId"`
}

func (m *ProductQuantity) TableName() string {
	return "product_quantities"
}

func ProductQuantityQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ProductQuantity{})
}
