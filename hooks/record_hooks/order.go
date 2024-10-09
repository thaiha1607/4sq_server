package record_hooks

import (
	"net/http"

	"example.com/4sq_server/dbquery"
	"example.com/4sq_server/hooks/shared"
	"example.com/4sq_server/utils"
	"example.com/4sq_server/utils/enum/order_status"
	pocketbase "github.com/AlperRehaYAZGAN/postgresbase"
	"github.com/AlperRehaYAZGAN/postgresbase/apis"
	"github.com/AlperRehaYAZGAN/postgresbase/core"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/samber/lo"
)

func forbidInvalidOrderStatus(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("orders").Add(func(e *core.RecordCreateEvent) error {
		if e.Record.GetString("statusCodeId") == order_status.Pending.ID() {
			return nil
		}
		return apis.NewBadRequestError("", map[string]validation.Error{
			"statusCodeId": validation.NewError(
				"invalid_status_code",
				"When creating an order, the status code must be 'Pending'",
			),
		})
	})
	app.OnRecordBeforeUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		old, err := dbquery.GetSingleOrder(app.Dao(), e.Record.Id)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Something happened on our end", nil)
		}
		if old.StatusCodeId == e.Record.GetString("statusCodeId") {
			return nil
		}
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
		return nil
	})
}

func assignWarehouseStaff(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("orders").Add(func(e *core.RecordUpdateEvent) error {
		if e.Record.GetString("statusCodeId") != order_status.Confirmed.ID() {
			return nil
		}
		return shared.AssignWarehouseStaff(app.Dao(), app.Logger(), e.Record)
	})
}
