package dbquery

import (
	"example.com/4sq_server/custom_models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

func GetWarehouseAssignmentsByInternalOrderId(dao *daos.Dao, internalOrderId string) ([]*custom_models.WarehouseAssignment, error) {
	var assignments []*custom_models.WarehouseAssignment
	err := custom_models.WarehouseAssignmentQuery(dao).
		Where(dbx.HashExp{"internalOrderId": internalOrderId}).
		All(&assignments)
	if err != nil {
		return nil, err
	}
	return assignments, nil
}
