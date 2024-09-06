package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/thaiha1607/4sq_server/hooks/apphooks"
)

// Register all hooks for the app
func RegisterHooks(app *pocketbase.PocketBase) {
	apphooks.RegisterAppHooks(app)
}
