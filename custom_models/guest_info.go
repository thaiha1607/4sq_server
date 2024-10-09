package custom_models

import (
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
	"github.com/pocketbase/dbx"
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
