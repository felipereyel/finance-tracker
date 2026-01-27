package cmd

import (
	"fintracker/internal/config"
	"fintracker/internal/controllers"
	"fintracker/internal/repositories/database"
	"fintracker/internal/routes"

	"fintracker/internal/migrations"
	"log"

	"github.com/pocketbase/pocketbase"
)

func Root() {
	app := pocketbase.New()
	db := database.NewDatabaseRepo(app)
	c := controllers.NewControllers(db)
	app.OnServe().BindFunc(routes.SetupRoutes(c))

	cfg := config.GetServerConfigs()
	if cfg.AutoMigrate {
		migrations.Register()
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
