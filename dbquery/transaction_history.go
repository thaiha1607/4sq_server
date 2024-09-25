package dbquery

import (
	"slices"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/thaiha1607/4sq_server/custom_models"
	"github.com/thaiha1607/4sq_server/utils"
)

func CreateNewTransactionHistory(dao *daos.Dao, entityType string, r *models.Record) error {
	if !slices.Contains(utils.AllowedTransactionHistoryEntities, entityType) {
		return nil
	}
	model := &custom_models.TransactionHistory{
		EntityType:   entityType,
		EntityId:     r.Id,
		StatusCodeId: r.GetString("statusCodeId"),
		Note:         "",
	}
	if err := dao.Save(model); err != nil {
		return err
	}
	return nil
}
