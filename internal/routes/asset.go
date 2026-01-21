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
		AssetId: newAsset.Id,
		Value:   newAsset.InitialPrice,
		Logged:  newAsset.BuyDate,
		Comment: "Initial price",
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

	return sendPage(e, components.AssetDetailsPage(asset))
}

// func taskNew(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	task, err := tc.CreateTask(user.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return c.Redirect("/edit/" + task.Id)
// }

// func taskEdit(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	taskId := c.Params("id")
// 	task, err := tc.RetrieveTask(user.ID, taskId)

// 	if err != nil {
// 		return err
// 	}

// 	return sendPage(c, components.TaskEditPage(task))
// }

// func taskSave(tc *controllers.PriceController, c *fiber.Ctx, user models.User) error {
// 	var taskId = c.Params("id")
// 	var taskChange controllers.TaskChange
// 	err := c.BodyParser(&taskChange)
// 	if err != nil {
// 		return err
// 	}

// 	if err := tc.UpdateTask(user.ID, taskId, taskChange); err != nil {
// 		return err
// 	}

// 	return c.SendStatus(fiber.StatusOK)
// }
