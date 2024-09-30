package cronjobs

import "github.com/pocketbase/pocketbase"

func RegisterCronJobs(app *pocketbase.PocketBase) {
	modifyPendingInternalOrdersOlderThan5Days(app)
	assignStaffToCompleteUnfinishedOrdersUpdatedMoreThanThreeDaysAgo(app)
}
