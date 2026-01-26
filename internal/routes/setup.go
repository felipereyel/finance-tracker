package routes

import (
	"github.com/pocketbase/pocketbase/core"
)

func SetupRoutes(se *core.ServeEvent) error {
	se.Router.GET("/", homeRedirect)

	se.Router.GET("/summary", accountSummary)
	se.Router.GET("/charts-content", accountChart)

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

// e.Request.BasicAuth()
