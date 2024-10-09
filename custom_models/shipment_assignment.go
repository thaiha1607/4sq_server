package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
)

var _ models.Model = (*ShipmentAssignment)(nil)

type ShipmentAssignment struct {
	models.BaseModel
	Status     string `db:"status" json:"status"`
	Note       string `db:"note" json:"note,omitempty"`
	ShipmentId string `db:"shipmentid" json:"shipmentId"`
	StaffId    string `db:"staffid" json:"staffId,omitempty"`
}

func (m *ShipmentAssignment) TableName() string {
	return "shipment_assignments"
}

func ShipmentAssignmentQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&ShipmentAssignment{})
}
