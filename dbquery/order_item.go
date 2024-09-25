package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetOrderItemsByOrderId(dao *daos.Dao, orderId string) ([]*custom_models.OrderItem, error) {
	var orderItems []*custom_models.OrderItem

	err := custom_models.OrderItemQuery(dao).
		Where(dbx.HashExp{"orderId": orderId}).
		All(&orderItems)
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
