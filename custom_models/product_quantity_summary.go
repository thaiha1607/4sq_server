package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*ProductQuantitySummary)(nil)

type ProductQuantitySummary struct {
	models.BaseModel
	TotalQty int64 `db:"totalQty" json:"totalQty"`
}

func (m *ProductQuantitySummary) TableName() string {
	return "product_quantity_summary"
}

func ProductQuantitySummaryQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ProductQuantitySummary{})
}
