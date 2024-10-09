package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*StaffInfo)(nil)

type StaffInfo struct {
	models.BaseModel
	StatusCode    string `db:"statusCode" json:"statusCode"`
	Role          string `db:"role" json:"role"`
	UserId        string `db:"userId" json:"userId"`
	WorkingUnitId string `db:"workingUnitId" json:"workingUnitId"`
}

func (m *StaffInfo) TableName() string {
	return "staff_info"
}

func StaffInfoQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&StaffInfo{})
}
