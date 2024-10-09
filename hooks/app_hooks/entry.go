package app_hooks

import (
	"os"

	pocketbase "github.com/AlperRehaYAZGAN/postgresbase"
	"github.com/AlperRehaYAZGAN/postgresbase/apis"
	"github.com/AlperRehaYAZGAN/postgresbase/core"
)

func RegisterAppHooks(app *pocketbase.PocketBase) {
	// Serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})
}
