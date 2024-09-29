package record_hooks

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/samber/lo"
	"github.com/thaiha1607/4sq_server/dbquery"
	"github.com/thaiha1607/4sq_server/hooks/shared"
	"github.com/thaiha1607/4sq_server/utils"
	"github.com/thaiha1607/4sq_server/utils/enum/order_status"
)

func forbidInvalidOrderStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("orders").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Pending.ID() {
			return apis.NewBadRequestError("", map[string]validation.Error{
				"statusCodeId": validation.NewError(
					"invalid_status_code",
					"When creating an order, the status code must be 'Pending'",
				),
			})
		}
		return nil
	})
	app.OnRecordBeforeUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := dbquery.GetSingleOrder(app.Dao(), e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		if old.StatusCodeId != e.Record.GetString("statusCodeId") {
			value, ok := utils.OrderStatusCodeTransitions[old.StatusCodeId]
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

func assignWarehouseStaff(app *pocketbase.PocketBase) {
	app.OnRecordBeforeUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := dbquery.GetSingleOrder(app.Dao(), e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		// Only assign warehouse staffs when the order is changed from Pending to Confirmed
		if old.StatusCodeId != order_status.Pending.ID() || e.Record.GetString("statusCodeId") != order_status.Confirmed.ID() {
			return nil
		}
		return shared.AssignWarehouseStaff(app.Dao(), app.Logger(), e.Record)
	})
}
