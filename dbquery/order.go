package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/pocketbase/dbx"
)

func GetSingleOrder(dao *daos.Dao, id string) (*custom_models.Order, error) {
	var order *custom_models.Order

	err := custom_models.OrderQuery(dao).
		Where(dbx.HashExp{"id": id}).
		One(&order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
