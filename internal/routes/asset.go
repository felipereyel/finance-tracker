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
	wallet := e.Request.URL.Query().Get(urls.WalletQueryParam)
	tag := e.Request.URL.Query().Get(urls.TagQueryParam)
	e.Response.Header().Set("HX-Redirect", urls.AssetsURLWithQuery(wallet, tag))
	return e.JSON(200, map[string]any{"success": true})
}

func assetList(c controllers.Controllers, e *core.RequestEvent) error {
	walletFilter := e.Request.URL.Query().Get(urls.WalletQueryParam)
	tagFilter := e.Request.URL.Query().Get(urls.TagQueryParam)
	userId := e.Get(userIdStoreKey).(string)

	summary, err := c.Asset.SummarizeAssets(userId, walletFilter, tagFilter)
	if err != nil {
		return err
	}

	return sendPage(e, components.AssetSummaryPage(summary))
}

func accountChart(c controllers.Controllers, e *core.RequestEvent) error {
	userId := e.Get(userIdStoreKey).(string)
	aggregation := e.Request.URL.Query().Get(urls.AggregationQueryParam)

	summary, err := c.Asset.SummarizeAssets(userId, "", "")
	if err != nil {
		return err
	}

	summary.Aggregation = aggregation
	return components.SummaryChart(summary, e.Response)
}

func accountSummary(c controllers.Controllers, e *core.RequestEvent) error {
	aggregation := e.Request.URL.Query().Get(urls.AggregationQueryParam)
	if aggregation == "" {
		aggregation = "total"
	}

	return sendPage(e, components.SummaryPage(aggregation))
}

func assetCreatePopup(c controllers.Controllers, e *core.RequestEvent) error {
	userId := e.Get(userIdStoreKey).(string)

	options, err := c.Asset.GetAssetOptions(userId)
	if err != nil {
		return err
	}

	return sendPage(e, components.NewAsset(options))
}

// missing wallet scope check
func assetCreate(c controllers.Controllers, e *core.RequestEvent) error {
	userId := e.Get(userIdStoreKey).(string)

	assetDTO := models.AssetCreateDTO{}
	if err := e.BindBody(&assetDTO); err != nil {
		return err
	}

	if err := c.User.ChechUserOwnsWallet(userId, assetDTO.Wallet); err != nil {
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

	priceDTO := models.PriceCreateDTO{}
	if err := e.BindBody(&priceDTO); err != nil {
		return err
	}

	if err := c.Price.CreatePrice(assetId, priceDTO); err != nil {
		return err
	}

	e.Response.Header().Set("HX-Redirect", urls.AssetIdGroupURL(assetId))
	return e.JSON(200, map[string]any{"success": true})
}

func historyPage(c controllers.Controllers, e *core.RequestEvent) error {
	aggregation := e.Request.URL.Query().Get(urls.AggregationQueryParam)
	if aggregation == "" {
		aggregation = "total"
	}

	return sendPage(e, components.HistoryPage(aggregation))
}

func historyChart(c controllers.Controllers, e *core.RequestEvent) error {
	userId := e.Get(userIdStoreKey).(string)
	aggregation := e.Request.URL.Query().Get(urls.AggregationQueryParam)
	if aggregation == "" {
		aggregation = "total"
	}

	history, err := c.History.GetHistory(userId, aggregation)
	if err != nil {
		return err
	}

	return components.HistoryChart(history, e.Response)
}
