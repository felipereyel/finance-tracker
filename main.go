package main

import (
	"fintracker/pkgs/handlers"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()
	app.Settings()
	publicDir := os.Getenv("PUBLIC_DIR")

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/healthz", handlers.Healthz)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDir), true))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
