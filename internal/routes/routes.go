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
	rc := NewRouteController(group, c)
	rc.BindFunc(basicAuthMiddleware)

	rc.GET(urls.AssetsRedirectPath, assetRedirect)

	rc.GET(urls.SummaryPath, accountSummary)
	rc.GET(urls.SummaryChartPath, accountChart)

	rc.GET(urls.AssetsPath, assetList)
	rc.POST(urls.AssetsPath, assetCreate)
	rc.GET(urls.AssetsPopupPath, assetCreatePopup)

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

type RouteController struct {
	router *router.RouterGroup[*core.RequestEvent]
	c      controllers.Controllers
}

func NewRouteController(router *router.RouterGroup[*core.RequestEvent], c controllers.Controllers) *RouteController {
	return &RouteController{router: router, c: c}
}

func (rc *RouteController) GET(path string, handler func(controllers.Controllers, *core.RequestEvent) error) {
	rc.router.GET(path, withControllerClousure(rc.c, handler))
}

func (rc *RouteController) POST(path string, handler func(controllers.Controllers, *core.RequestEvent) error) {
	rc.router.POST(path, withControllerClousure(rc.c, handler))
}

func (rc *RouteController) BindFunc(handler func(controllers.Controllers, *core.RequestEvent) error) {
	rc.router.BindFunc(withControllerClousure(rc.c, handler))
}

func discardHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}

func healthzHandler(e *core.RequestEvent) error {
	return e.JSON(200, map[string]any{"success": true})
}
