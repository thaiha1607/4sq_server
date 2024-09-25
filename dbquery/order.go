package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetSingleOrder(dao *daos.Dao, id string) (*custom_models.Order, error) {
	var order *custom_models.Order

	err := custom_models.OrderItemQuery(dao).
		Where(dbx.HashExp{"id": id}).
		One(&order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
