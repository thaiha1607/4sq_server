package custom_models

import (
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*ProductQuantitySummary)(nil)

type ProductQuantitySummary struct {
	models.BaseModel
	CategoryId string `db:"categoryId" json:"categoryId"`
	TotalQty   string `db:"totalQty" json:"totalQty"`
}

func (m *ProductQuantitySummary) TableName() string {
	return "product_quantity_summary"
}
