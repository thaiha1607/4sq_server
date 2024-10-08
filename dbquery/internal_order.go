package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
)

func GetInternalOrdersByOrderId(dao *daos.Dao, orderId string) ([]*custom_models.InternalOrder, error) {
	var internalOrder []*custom_models.InternalOrder
	err := custom_models.
		InternalOrderQuery(dao).
		Where(dbx.HashExp{"rootOrderId": orderId}).
		All(&internalOrder)
	if err != nil {
		return nil, err
	}
	return internalOrder, nil
}

func GetInternalOrdersByShipmentId(dao *daos.Dao, shipmentId string) ([]*custom_models.InternalOrder, error) {
	var internalOrders []*custom_models.InternalOrder
	err := custom_models.
		InternalOrderQuery(dao).
		Where(dbx.HashExp{"shipmentId": shipmentId}).
		OrderBy("created DESC").
		All(&internalOrders)
	if err != nil {
		return nil, err
	}
	return internalOrders, nil
}
