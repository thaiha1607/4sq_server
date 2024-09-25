package dbquery

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/thaiha1607/4sq_server/custom_models"
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
