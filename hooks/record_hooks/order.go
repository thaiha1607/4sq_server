package record_hooks

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/thaiha1607/4sq_server/constants/enum/order_status"
)

func registerOrderHooks(app *pocketbase.PocketBase) {
	forbidInvalidStatusOnCreateOrder(app)
	assignStaffAutomatically(app)
}

func forbidInvalidStatusOnCreateOrder(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("orders").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Pending.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError("invalid_status_code", "When creating an order, the status code must be 'Pending'"),
			})
		}
		return nil
	})
}

func assignStaffAutomatically(app *pocketbase.PocketBase) {}
