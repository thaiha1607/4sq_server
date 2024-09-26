package record_hooks

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/samber/lo"
	"github.com/thaiha1607/4sq_server/utils"
	"github.com/thaiha1607/4sq_server/utils/enum/invoice_status"
)

func forbidInvalidInvoiceStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("invoices").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != invoice_status.Draft.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError("invalid_status_code", "When creating an invoice, the status code must be 'Draft'"),
			})
		}
		return nil
	})
	app.OnRecordBeforeUpdateRequest("invoices").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("invoices", e.Record.Id)
		if err != nil {
			return apis.NewNotFoundError("Invoice not found", nil)
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