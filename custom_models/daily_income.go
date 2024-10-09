package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
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
