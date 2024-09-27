package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*StaffInfo)(nil)

type StaffInfo struct {
	models.BaseModel
	StatusCode    string `db:"statusCode" json:"statusCode"`
	Role          string `db:"role" json:"role"`
	UserId        string `db:"userId" json:"userId"`
	WorkingUnitId string `db:"workingUnitId" json:"workingUnitId,omitempty"`
}

func (m *StaffInfo) TableName() string {
	return "staff_info"
}

func StaffInfoQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&StaffInfo{})
}
