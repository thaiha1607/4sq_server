package main

import (
	"log"
	"os"
	"strings"

	"example.com/4sq_server/hooks"
	_ "example.com/4sq_server/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	hooks.RegisterHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
