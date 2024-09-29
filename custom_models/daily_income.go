package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*DailyIncome)(nil)

type DailyIncome struct {
	models.BaseModel
	AmountOfChange float64 `db:"amountOfChange" json:"amountOfChange"`
}

func (m *DailyIncome) TableName() string {
	return "daily_income"
}

func DailyIncomeQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&DailyIncome{})
}
