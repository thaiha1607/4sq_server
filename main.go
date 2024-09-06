package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/thaiha1607/4sq_server/hooks"
)

func main() {
	app := pocketbase.New()

	hooks.RegisterHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
