package custom_models

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*GuestInfo)(nil)

type GuestInfo struct {
	models.BaseModel
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email,omitempty"`
	Phone     string `db:"phone" json:"phone"`
	AddressId string `db:"addressId" json:"addressId"`
}

func (m *GuestInfo) TableName() string {
	return "guest_info"
}

func GuestInfoQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&GuestInfo{})
}
