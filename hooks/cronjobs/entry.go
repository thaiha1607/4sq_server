package cronjobs

import "github.com/pocketbase/pocketbase"

func RegisterCronJobs(app *pocketbase.PocketBase) {
	modifyPendingInternalOrdersOlderThan48Hours(app)
}
