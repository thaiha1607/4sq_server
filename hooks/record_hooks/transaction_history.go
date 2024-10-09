package record_hooks

import (
	"net/http"

	"example.com/4sq_server/dbquery"
	"example.com/4sq_server/utils"
	pocketbase "github.com/AlperRehaYAZGAN/postgresbase"
	"github.com/AlperRehaYAZGAN/postgresbase/apis"
	"github.com/AlperRehaYAZGAN/postgresbase/core"
)

func addTransactionHistoryHooks(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest(utils.AllowedTransactionHistoryEntities...).Add(func(e *core.RecordCreateEvent) error {
		return dbquery.CreateNewTransactionHistory(app.Dao(), e.Collection.Name, e.Record)
	})
	app.OnRecordBeforeUpdateRequest(utils.AllowedTransactionHistoryEntities...).Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById(e.Collection.Name, e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			return dbquery.CreateNewTransactionHistory(app.Dao(), e.Collection.Name, e.Record)
		}
		return nil
	})
}
