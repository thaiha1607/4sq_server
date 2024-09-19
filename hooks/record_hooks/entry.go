package record_hooks

import "github.com/pocketbase/pocketbase"

func RegisterRecordHooks(app *pocketbase.PocketBase) {
	registerOrderHooks(app)
	registerTransactionHistory(app)
}
