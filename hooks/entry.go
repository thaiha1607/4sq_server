package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/thaiha1607/4sq_server/hooks/app_hooks"
	"github.com/thaiha1607/4sq_server/hooks/record_hooks"
)

// Register all hooks for the app
func RegisterHooks(app *pocketbase.PocketBase) {
	app_hooks.RegisterAppHooks(app)
	record_hooks.RegisterRecordHooks(app)
}
