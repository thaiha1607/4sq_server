package dbquery

import (
	"slices"

	"example.com/4sq_server/custom_models"
	"example.com/4sq_server/utils"
	"github.com/AlperRehaYAZGAN/postgresbase/daos"
	"github.com/AlperRehaYAZGAN/postgresbase/models"
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
	_ = dao.Save(model)
	return nil
}
