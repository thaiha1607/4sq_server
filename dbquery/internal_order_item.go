package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func GetInternalOrderItemsByInternalOrderId(dao *daos.Dao, internalOrderId string) ([]*custom_models.InternalOrderItem, error) {
	var internalOrderItems []*custom_models.InternalOrderItem

	err := custom_models.
		InternalOrderItemQuery(dao).
		Where(dbx.HashExp{
			"internalOrderId": internalOrderId,
		},
		).
		All(&internalOrderItems)
	if err != nil {
		return nil, err
	}

	return internalOrderItems, nil
}

func GetSingleInternalOrderItem(dao *daos.Dao, id string) (*custom_models.InternalOrderItem, error) {
	var internalOrderItem *custom_models.InternalOrderItem

	err := custom_models.
		InternalOrderItemQuery(dao).
		Where(dbx.HashExp{
			"id": id,
		},
		).
		One(&internalOrderItem)
	if err != nil {
		return nil, err
	}

	return internalOrderItem, nil
}
