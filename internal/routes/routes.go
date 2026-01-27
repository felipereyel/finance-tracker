package routes

import (
	"fintracker/internal/controllers"
	"fintracker/internal/urls"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func SetupRoutes(c controllers.Controllers) func(*core.ServeEvent) error {
	return func(se *core.ServeEvent) error {
		se.Router.GET(urls.StaticsURL("{path...}"), staticsHandler)
		se.Router.GET(urls.DiscardURL, discardHandler)
		se.Router.GET(urls.HealthzURL, healthzHandler)
		se.Router.GET(urls.Root, homeRedirect)

		authenticatedGroup := se.Router.Group(urls.AuthenticatedUrl)
		setupAuthenticatedRoutes(authenticatedGroup, c)

		return se.Next()
	}
}

func setupAuthenticatedRoutes(group *router.RouterGroup[*core.RequestEvent], c controllers.Controllers) {
	group.BindFunc(withControllerClousure(c, basicAuthMiddleware))

	group.GET(urls.AssetsRedirectPath, assetRedirect)

	group.GET(urls.SummaryPath, withControllerClousure(c, accountSummary))
	group.GET(urls.SummaryChartPath, withControllerClousure(c, accountChart))

	group.GET(urls.AssetsPath, withControllerClousure(c, assetList))
	group.POST(urls.AssetsPath, withControllerClousure(c, assetCreate))
	group.GET(urls.AssetsPopupPath, withControllerClousure(c, assetCreatePopup))

	scopedAssetsGroup := group.Group(urls.AssetIdGroupPath)
	setupScopedAssetsRoutes(scopedAssetsGroup, c)

	scopedPricesGroup := group.Group(urls.PriceIdGroupPath)
	setupScopedPricesRoutes(scopedPricesGroup, c)
}

func withControllerClousure(c controllers.Controllers, handler func(controllers.Controllers, *core.RequestEvent) error) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		return handler(c, e)
	}
}

func homeRedirect(e *core.RequestEvent) error {
	return e.Redirect(302, urls.AssetsURL)
}

func discardHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func healthzHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}
