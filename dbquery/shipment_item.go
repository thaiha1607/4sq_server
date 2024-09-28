package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
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
