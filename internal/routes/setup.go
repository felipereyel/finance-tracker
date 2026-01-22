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
	se.Router.POST("/assets/{asset_id}", assetUpdate)

	se.Router.GET("/assets/{asset_id}/chart", assetChart)
	se.Router.GET("/assets/{asset_id}/prices", assetPriceTable)

	se.Router.GET("/assets/{asset_id}/price-popup", priceCreatePopup)
	se.Router.POST("/assets/{asset_id}/prices", priceCreate)

	se.Router.GET("/prices/{price_id}", priceDetails)
	se.Router.POST("/prices/{price_id}", priceUpdate)

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

// authRecord := e.Auth

// isGuest := e.Auth == nil

// // the same as "e.Auth != nil && e.Auth.IsSuperuser()"
// isSuperuser := e.HasSuperuserAuth()
