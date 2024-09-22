package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*InternalOrder)(nil)

type InternalOrder struct {
	models.BaseModel
	Type         string `db:"type" json:"type"`
	Note         string `db:"note" json:"note,omitempty"`
	StatusCodeId string `db:"statusCodeId" json:"statusCodeId"`
	RootOrderId  string `db:"rootOrderId" json:"rootOrderId"`
}

func (m *InternalOrder) TableName() string {
	return "internal_orders"
}

func InternalOrderQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&InternalOrder{})
}
