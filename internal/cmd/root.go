package cmd

import (
	"fintracker/internal/config"
	"fintracker/internal/routes"

	"fintracker/internal/migrations"
	"log"

	"github.com/pocketbase/pocketbase"
)

func Root() {
	app := pocketbase.New()
	cfg := config.GetServerConfigs()
	app.OnServe().BindFunc(routes.SetupRoutes)

	if cfg.AutoMigrate {
		migrations.Register()
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
