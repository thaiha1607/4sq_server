package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/pocketbase/dbx"
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

func GetSingleOrderItem(dao *daos.Dao, id string) (*custom_models.OrderItem, error) {
	var orderItem *custom_models.OrderItem

	err := custom_models.OrderItemQuery(dao).
		Where(dbx.HashExp{"id": id}).
		One(&orderItem)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}
