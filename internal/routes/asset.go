package routes

import (
	"fintracker/internal/components"
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func assetList(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	wallet := e.Request.URL.Query().Get("wallet")
	asset_type := e.Request.URL.Query().Get("type")

	aggregated, err := db.ListAssetAggregates(wallet, asset_type)
	if err != nil {
		return err
	}

	wallets, err := db.ListWallets()
	if err != nil {
		return err
	}

	summary := models.Summary{
		Total:      0,
		Aggregates: make([]models.AssetAggregate, 0),

		AssetTypes:   models.AssetTypes,
		SelectedType: asset_type,

		SelectedWallet: wallet,
		Wallets:        make([][]string, 0),
	}

	for _, asset := range aggregated {
		summary.Total += asset.LastPrice
		summary.Aggregates = append(summary.Aggregates, asset)
	}

	for _, wallet := range wallets {
		summary.Wallets = append(summary.Wallets, []string{wallet.Id, wallet.Name})
	}

	return sendPage(e, components.AssetSummaryPage(summary))
}

func accountChart(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	aggregated, err := db.ListAssetAggregates("", "")
	if err != nil {
		return err
	}

	wallets, err := db.ListWallets()
	if err != nil {
		return err
	}

	summary := models.Summary{
		Total:      0,
		Aggregates: make([]models.AssetAggregate, 0),

		AssetTypes:   models.AssetTypes,
		SelectedType: "",

		SelectedWallet: "",
		Wallets:        make([][]string, 0),
	}

	for _, asset := range aggregated {
		summary.Total += asset.LastPrice
		summary.Aggregates = append(summary.Aggregates, asset)
	}

	for _, wallet := range wallets {
		summary.Wallets = append(summary.Wallets, []string{wallet.Id, wallet.Name})
	}

	return components.SummaryChart(summary, e.Response)
}

func accountSummary(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	aggregated, err := db.ListAssetAggregates("", "")
	if err != nil {
		return err
	}

	wallets, err := db.ListWallets()
	if err != nil {
		return err
	}

	summary := models.Summary{
		Total:      0,
		Aggregates: make([]models.AssetAggregate, 0),

		AssetTypes:   models.AssetTypes,
		SelectedType: "",

		SelectedWallet: "",
		Wallets:        make([][]string, 0),
	}

	for _, asset := range aggregated {
		summary.Total += asset.LastPrice
		summary.Aggregates = append(summary.Aggregates, asset)
	}

	for _, wallet := range wallets {
		summary.Wallets = append(summary.Wallets, []string{wallet.Id, wallet.Name})
	}

	return sendPage(e, components.AccountChartsPage(summary))
}

func assetRedirect(e *core.RequestEvent) error {
	wallet := e.Request.URL.Query().Get("wallet")
	asset_type := e.Request.URL.Query().Get("type")
	e.Response.Header().Set("HX-Redirect", "/assets?wallet="+wallet+"&type="+asset_type)
	return e.JSON(200, map[string]any{"success": true})
}

func assetCreatePopup(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	// TODO: init popup with selected fiels based in query
	// wallet := e.Request.URL.Query().Get("wallet")
	// asset_type := e.Request.URL.Query().Get("type")

	wallets, err := db.ListWallets()
	if err != nil {
		return err
	}

	summary := models.NewAssetSummary{
		AssetTypes: models.AssetTypes,
		Wallets:    make([][]string, 0),
	}

	for _, wallet := range wallets {
		summary.Wallets = append(summary.Wallets, []string{wallet.Id, wallet.Name})
	}

	return sendPage(e, components.NewAsset(summary))
}

func assetCreate(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	assetDTO := models.AssetCreateDTO{}
	if err := e.BindBody(&assetDTO); err != nil {
		return err
	}

	newAsset := models.CreateNewAsset(assetDTO)
	if err := db.CreateAsset(newAsset); err != nil {
		return err
	}

	priceDTO := models.PriceCreateDTO{
		AssetId:  newAsset.Id,
		Value:    newAsset.InitialPrice,
		LoggedAt: newAsset.BuyDate,
		Comment:  "Initial price",
	}

	newPrice := models.CreateNewPrice(priceDTO)
	if err := db.CreatePrice(newPrice); err != nil {
		return err
	}

	e.Response.Header().Set("HX-Redirect", "/assets/"+newAsset.Id)
	return e.JSON(200, map[string]any{"success": true})
}

func assetDetails(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	assetId := e.Request.PathValue("asset_id")

	asset, err := db.GetAssetAggregateById(assetId)
	if err != nil {
		fmt.Println("Error retrieving asset:", err)
		return err
	}

	prices, err := db.ListPricesByAssetId(assetId)
	if err != nil {
		fmt.Println("Error retrieving prices:", err)
		return err
	}

	return sendPage(e, components.AssetDetailsPage(asset, prices))
}

func assetPriceTable(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	assetId := e.Request.PathValue("asset_id")

	asset, err := db.GetAssetAggregateById(assetId)
	if err != nil {
		fmt.Println("Error retrieving asset:", err)
		return err
	}

	prices, err := db.ListPricesByAssetId(assetId)
	if err != nil {
		fmt.Println("Error retrieving prices:", err)
		return err
	}

	return sendPage(e, components.AssetPricesPage(asset, prices))
}

func assetUpdate(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	assetId := e.Request.PathValue("asset_id")

	assetDTO := models.AssetUpdateDTO{}
	if err := e.BindBody(&assetDTO); err != nil {
		return err
	}

	asset, err := db.GetAssetById(assetId)
	if err != nil {
		return err
	}

	asset.SellDate = assetDTO.SellDate
	asset.Comment = assetDTO.Comment
	asset.Updated = models.GenerateTimestamp()

	if err := db.UpdateAsset(asset); err != nil {
		return err
	}

	return e.JSON(200, map[string]any{"success": true})
}

func assetChart(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	assetId := e.Request.PathValue("asset_id")

	prices, err := db.ListPricesByAssetId(assetId)
	if err != nil {
		return err
	}

	return components.PriceChart(prices, e.Response)
}
