package hooks

import (
	"example.com/4sq_server/hooks/app_hooks"
	"example.com/4sq_server/hooks/cronjobs"
	"example.com/4sq_server/hooks/record_hooks"
	"github.com/pocketbase/pocketbase"
)

// Register all hooks for the app
func RegisterHooks(app *pocketbase.PocketBase) {
	cronjobs.RegisterCronJobs(app)
	app_hooks.RegisterAppHooks(app)
	record_hooks.RegisterRecordHooks(app)
}
