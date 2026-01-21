package routes

import (
	"github.com/pocketbase/pocketbase/core"
)

func SetupRoutes(se *core.ServeEvent) error {
	// / -> redirect to /assets
	se.Router.GET("/", homeRedirect)

	// /assets?wallet=XXX -> asset list (home)
	se.Router.GET("/assets", assetList)
	se.Router.GET("/assets-redirect", assetRedirect)

	se.Router.GET("/assets-popup", assetCreatePopup)
	se.Router.POST("/assets", assetCreate)

	se.Router.GET("/assets/{asset_id}", assetDetails)
	se.Router.GET("/assets/{asset_id}/", assetDetails)

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
	se.Router.GET("/discard", discardHandler)
	se.Router.GET("/healthz", healthzHandler)

	// app.Use(notFoundHandler)

	return se.Next()
}

func discardHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func healthzHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func homeRedirect(e *core.RequestEvent) error {
	return e.Redirect(302, "/assets")
}
