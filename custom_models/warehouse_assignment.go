package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*WarehouseAssignment)(nil)

type WarehouseAssignment struct {
	models.BaseModel
	Status           string `db:"status" json:"status"`
	Note             string `db:"note" json:"note,omitempty"`
	OtherInfo        string `db:"otherInfo" json:"otherInfo,omitempty"`
	StaffId          string `db:"staffId" json:"staffId,omitempty"`
	SrcWorkingUnitId string `db:"srcWorkingUnitId" json:"srcWorkingUnitId"`
	DstWorkingUnitId string `db:"dstWorkingUnitId" json:"dstWorkingUnitId,omitempty"`
	InternalOrderId  string `db:"internalOrderId" json:"internalOrderId"`
}

func (m *WarehouseAssignment) TableName() string {
	return "warehouse_assignments"
}

func WarehouseAssignmentQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&WarehouseAssignment{})
}
