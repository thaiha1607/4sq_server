package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*ShipmentAssignment)(nil)

type ShipmentAssignment struct {
	models.BaseModel
	Status     string `db:"status" json:"status"`
	Note       string `db:"note" json:"note,omitempty"`
	OtherInfo  string `db:"otherinfo" json:"otherInfo,omitempty"`
	ShipmentId string `db:"shipmentid" json:"shipmentId"`
	StaffId    string `db:"staffid" json:"staffId,omitempty"`
}

func (m *ShipmentAssignment) TableName() string {
	return "shipment_assignments"
}

func ShipmentAssignmentQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ShipmentAssignment{})
}
