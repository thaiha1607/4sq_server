package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Order)(nil)

type Order struct {
	models.BaseModel
	Type         string `db:"type" json:"type"`
	Priority     int64  `db:"priority" json:"priority,omitempty"`
	Rating       int64  `db:"rating" json:"rating,omitempty"`
	Note         string `db:"note" json:"note,omitempty"`
	OtherInfo    string `db:"otherInfo" json:"otherInfo,omitempty"`
	RootOrderId  string `db:"rootOrderId" json:"rootOrderId,omitempty"`
	CreatorId    string `db:"creatorId" json:"creatorId"`
	GuestId      string `db:"guestId" json:"guestId,omitempty"`
	StatusCodeId string `db:"statusCodeId" json:"statusCodeId"`
	AddressId    string `db:"addressId" json:"addressId"`
}

func (m *Order) TableName() string {
	return "orders"
}

func OrderQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Order{})
}
