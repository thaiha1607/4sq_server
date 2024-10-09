package record_hooks

import (
	"example.com/4sq_server/custom_models"
	"example.com/4sq_server/utils"
	"example.com/4sq_server/utils/enum/invoice_status"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/samber/lo"
)

func forbidInvalidInvoiceStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("invoices").Add(func(e *core.RecordUpdateEvent) error {
		old, err := app.Dao().FindRecordById("invoices", e.Record.Id)
		if err != nil {
			return apis.NewApiError(500, "Something happened on our end", nil)
		}
		if old.GetString("statusCodeId") == e.Record.GetString("statusCodeId") {
			return nil
		}
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
		return nil
	})
}

func updateDailyIncomeWhenInvoiceIsPaid(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("invoices").Add(func(e *core.RecordUpdateEvent) error {
		if e.Record.GetString("statusCodeId") == invoice_status.Paid.ID() {
			dailyIncome := &custom_models.DailyIncome{
				AmountOfChange: e.Record.GetFloat("totalAmount"),
			}
			_ = app.Dao().Save(dailyIncome)
		}
		return nil
	})
}
