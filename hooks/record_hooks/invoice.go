package record_hooks

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/samber/lo"
	"github.com/thaiha1607/4sq_server/utils"
)

func forbidInvalidInvoiceStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("invoices").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("invoices", e.Record.Id)
		if err != nil {
			return apis.NewApiError(500, "Something happened on our end", nil)
		}
		if old.GetString("statusCodeId") != e.Record.GetString("statusCodeId") {
			value, ok := utils.InvoiceStatusCodeTransitions[old.GetString("statusCodeId")]
			if !ok {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"statusCodeId": validation.NewError("invalid_status_code", "Invalid status code"),
				})
			}
			if !lo.Contains(value, e.Record.GetString("statusCodeId")) {
				return apis.NewBadRequestError("", map[string]validation.Error{
					"statusCodeId": validation.NewError("invalid_status_code", "Invalid status code transition"),
				})
			}
		}
		return nil
	})
}
