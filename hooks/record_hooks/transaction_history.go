package record_hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/thaiha1607/4sq_server/dbquery"
	"github.com/thaiha1607/4sq_server/utils"
)

func addTransactionHistoryHooks(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest(utils.AllowedTransactionHistoryEntities...).Add(func(e *core.RecordCreateEvent) error {
		return dbquery.CreateNewTransactionHistory(app.Dao(), e.Collection.Name, e.Record)
	})
	app.OnRecordBeforeUpdateRequest(utils.AllowedTransactionHistoryEntities...).Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById(e.Collection.Name, e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Record from "+e.Collection.Name+" not found", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			return dbquery.CreateNewTransactionHistory(app.Dao(), e.Collection.Name, e.Record)
		}
		return nil
	})
}
