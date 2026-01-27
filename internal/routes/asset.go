package routes

import (
	"fintracker/internal/components"
	"fintracker/internal/controllers"
	"fintracker/internal/models"
	"fintracker/internal/urls"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func setupScopedAssetsRoutes(group *router.RouterGroup[*core.RequestEvent], c controllers.Controllers) {
	group.BindFunc(withControllerClousure(c, assetScopeCheckMiddleware))

	group.GET(urls.Root, withControllerClousure(c, assetDetails))
	group.POST(urls.Root, withControllerClousure(c, assetUpdate))

	group.GET(urls.AssetIdPricesPath, withControllerClousure(c, assetPriceTable))
	group.POST(urls.AssetIdPricesPath, withControllerClousure(c, assetPriceCreate))
	group.GET(urls.AssetIdPricesChartPath, withControllerClousure(c, assetPriceChart))
	group.GET(urls.AssetIdPricesPopupPath, withControllerClousure(c, assetPricePopup))
}

func assetRedirect(e *core.RequestEvent) error {
	// TODO: add query to urls file
	wallet := e.Request.URL.Query().Get("wallet")
	asset_type := e.Request.URL.Query().Get("type")
	e.Response.Header().Set("HX-Redirect", urls.AssetsURLWithQuey(wallet, asset_type))
	return e.JSON(200, map[string]any{"success": true})
}

func assetList(c controllers.Controllers, e *core.RequestEvent) error {
	walletFilter := e.Request.URL.Query().Get("wallet")
	typeFilter := e.Request.URL.Query().Get("type")

	summary, err := c.Asset.SummarizeAssets(walletFilter, typeFilter)
	if err != nil {
		return err
	}

	return sendPage(e, components.AssetSummaryPage(summary))
}

func accountChart(c controllers.Controllers, e *core.RequestEvent) error {
	// TODO: add filters to chart

	summary, err := c.Asset.SummarizeAssets("", "")
	if err != nil {
		return err
	}

	return components.SummaryChart(summary, e.Response)
}

func accountSummary(c controllers.Controllers, e *core.RequestEvent) error {
	// TODO: add filters to chart

	summary, err := c.Asset.SummarizeAssets("", "")
	if err != nil {
		return err
	}

	return sendPage(e, components.SummaryPage(summary))
}

func assetCreatePopup(c controllers.Controllers, e *core.RequestEvent) error {
	// TODO: init popup with selected fiels based in query
	// wallet := e.Request.URL.Query().Get("wallet")
	// asset_type := e.Request.URL.Query().Get("type")

	options, err := c.Asset.GetAssetOptions()
	if err != nil {
		return err
	}

	return sendPage(e, components.NewAsset(options))
}

func assetCreate(c controllers.Controllers, e *core.RequestEvent) error {
	assetDTO := models.AssetCreateDTO{}
	if err := e.BindBody(&assetDTO); err != nil {
		return err
	}

	assetId, err := c.Asset.CreateAsset(assetDTO)
	if err != nil {
		return err
	}

	e.Response.Header().Set("HX-Redirect", urls.AssetIdGroupURL(assetId))
	return e.JSON(200, map[string]any{"success": true})
}

func assetDetails(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	asset, err := c.Asset.GetAssetAggregate(assetId)
	if err != nil {
		fmt.Println("Error retrieving asset:", err)
		return err
	}

	return sendPage(e, components.AssetDetailsPage(asset))
}

func assetPriceTable(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	assetAgg, prices, err := c.Price.ListPricesEnrich(assetId)
	if err != nil {
		fmt.Println("Error retrieving info:", err)
		return err
	}

	return sendPage(e, components.AssetPricesPage(assetAgg, prices))
}

func assetUpdate(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	assetDTO := models.AssetUpdateDTO{}
	if err := e.BindBody(&assetDTO); err != nil {
		return err
	}

	if err := c.Asset.UpdateAsset(assetId, assetDTO); err != nil {
		return err
	}

	return e.JSON(200, map[string]any{"success": true})
}

func assetPriceChart(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	prices, err := c.Price.ListPrices(assetId)
	if err != nil {
		return err
	}

	return components.PriceChart(prices, e.Response)
}

func assetPricePopup(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	asset, err := c.Asset.GetAssetAggregate(assetId)
	if err != nil {
		fmt.Println("Error retrieving asset:", err)
		return err
	}

	return sendPage(e, components.NewPrice(asset))
}

func assetPriceCreate(c controllers.Controllers, e *core.RequestEvent) error {
	assetId := e.Request.PathValue(urls.AssetIdPathParam)

	priceDTO := models.PriceCreateDTO{AssetId: assetId}
	if err := e.BindBody(&priceDTO); err != nil {
		return err
	}

	if err := c.Price.CreatePrice(assetId, priceDTO); err != nil {
		return err
	}

	e.Response.Header().Set("HX-Redirect", urls.AssetIdGroupURL(priceDTO.AssetId))
	return e.JSON(200, map[string]any{"success": true})
}
