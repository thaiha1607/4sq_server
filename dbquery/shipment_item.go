package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/pocketbase/dbx"
)

func GetShipmentItemsByShipmentId(dao *daos.Dao, shipmentId string) ([]*custom_models.ShipmentItem, error) {
	var shipmentItems []*custom_models.ShipmentItem

	err := custom_models.ShipmentItemQuery(dao).
		Where(dbx.HashExp{"shipmentId": shipmentId}).
		All(&shipmentItems)
	if err != nil {
		return nil, err
	}

	return shipmentItems, nil
}
