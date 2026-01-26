package routes

import (
	"fintracker/internal/controllers"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func SetupRoutes(c controllers.Controllers) func(*core.ServeEvent) error {
	return func(se *core.ServeEvent) error {
		se.Router.GET("/_statics/{path...}", staticsHandler)
		se.Router.GET("/_discard", discardHandler)
		se.Router.GET("/_healthz", healthzHandler)
		se.Router.GET("/", homeRedirect)

		authenticatedGroup := se.Router.Group("/u")
		setupAuthenticatedRoutes(authenticatedGroup, c)

		return se.Next()
	}
}

func setupAuthenticatedRoutes(group *router.RouterGroup[*core.RequestEvent], c controllers.Controllers) {
	group.BindFunc(basicAuthMiddleware)

	group.GET("/summary", withControllerClousure(c, accountSummary))
	group.GET("/summary-chart", withControllerClousure(c, accountChart))

	group.GET("/assets-redirect", assetRedirect)

	group.GET("/assets", withControllerClousure(c, assetList))
	group.POST("/assets", withControllerClousure(c, assetCreate))
	group.GET("/assets-popup", withControllerClousure(c, assetCreatePopup))

	scopedAssetsGroup := group.Group("/assets/{asset_id}")
	setupScopedAssetsRoutes(scopedAssetsGroup, c)

	scopedPricesGroup := group.Group("/prices/{price_id}")
	setupScopedPricesRoutes(scopedPricesGroup, c)
}

func withControllerClousure(c controllers.Controllers, handler func(controllers.Controllers, *core.RequestEvent) error) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		return handler(c, e)
	}
}

func discardHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func healthzHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func homeRedirect(e *core.RequestEvent) error {
	return e.Redirect(302, "/u/assets")
}
