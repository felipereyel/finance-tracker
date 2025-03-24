package routes

import (
	"github.com/pocketbase/pocketbase/core"
)

func SetupRoutes(se *core.ServeEvent) error {
	// dbRepo := database.NewDatabaseRepo(se.App)

	// / -> redirect to /assets
	se.Router.GET("/", homeRedirect)

	// /assets?wallet=XXX -> asset list (home)
	se.Router.GET("/assets", assetList)

	// get  new asset popup
	// post new asset

	// get asset
	// put asset
	// get new price popup
	// post new price

	// asset price
	// put asset price

	// app.Get("/", uc, tc, taskList)
	// app.Get("/new", uc, tc, taskNew)
	// app.Get("/edit/:id", uc, tc, taskEdit)
	// app.Post("/edit/:id", uc, tc, taskSave)

	se.Router.GET("/statics/{path...}", assetsHandler)
	// app.Use("/healthz", healthzHandler)
	// app.Use(notFoundHandler)

	return se.Next()
}
