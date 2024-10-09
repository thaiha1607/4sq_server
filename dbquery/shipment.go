package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/pocketbase/dbx"
)

func GetShipmentsByOrderId(dao *daos.Dao, orderId string) ([]*custom_models.Shipment, error) {
	var shipments []*custom_models.Shipment
	err := custom_models.
		ShipmentQuery(dao).
		Where(dbx.HashExp{"orderId": orderId}).
		All(&shipments)
	if err != nil {
		return nil, err
	}
	return shipments, nil
}

func GetSingleShipment(dao *daos.Dao, id string) (*custom_models.Shipment, error) {
	var shipment *custom_models.Shipment
	err := custom_models.
		ShipmentQuery(dao).
		Where(dbx.HashExp{"id": id}).
		One(&shipment)
	if err != nil {
		return nil, err
	}
	return shipment, nil
}
