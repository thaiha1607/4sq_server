package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*WorkingUnit)(nil)

type WorkingUnit struct {
	models.BaseModel
	Name      string `db:"name" json:"name"`
	Type      string `db:"type" json:"type"`
	ImageUrl  string `db:"imageUrl" json:"imageUrl,omitempty"`
	AddressId string `db:"addressId" json:"addressId,omitempty"`
}

func (m *WorkingUnit) TableName() string {
	return "working_units"
}

func WorkingUnitQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&WorkingUnit{})
}
