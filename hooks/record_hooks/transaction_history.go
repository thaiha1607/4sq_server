package record_hooks

import (
	"slices"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/thaiha1607/4sq_server/custom_models"
)

var allowedTransactionHistoryEntities = []string{
	"internal_orders",
	"invoices",
	"orders",
	"shipments",
}

func createNewTransactionHistory(app *pocketbase.PocketBase, entityType string, r *models.Record) error {
	if !slices.Contains(allowedTransactionHistoryEntities, entityType) {
		return nil
	}
	model := &custom_models.TransactionHistory{
		EntityType:   entityType,
		EntityId:     r.Id,
		StatusCodeId: r.GetString("statusCodeId"),
		Note:         "",
	}
	if err := app.Dao().Save(model); err != nil {
		return err
	}
	return nil
}

func registerTransactionHistory(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest(allowedTransactionHistoryEntities...).Add(func(e *core.RecordCreateEvent) error {
		return createNewTransactionHistory(app, e.Collection.Name, e.Record)
	})
	app.OnRecordAfterUpdateRequest(allowedTransactionHistoryEntities...).Add(func(e *core.RecordUpdateEvent) error {
		return createNewTransactionHistory(app, e.Collection.Name, e.Record)
	})
}
