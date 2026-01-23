package routes

import (
	"fintracker/internal/components"
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func priceCreatePopup(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	assetId := e.Request.PathValue("asset_id")

	asset, err := db.GetAssetAggregateById(assetId)
	if err != nil {
		fmt.Println("Error retrieving asset:", err)
		return err
	}

	return sendPage(e, components.NewPrice(asset))
}

func priceCreate(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)

	assetId := e.Request.PathValue("asset_id")
	priceDTO := models.PriceCreateDTO{AssetId: assetId}
	if err := e.BindBody(&priceDTO); err != nil {
		return err
	}

	newPrice := models.CreateNewPrice(priceDTO)
	if err := db.CreatePrice(newPrice); err != nil {
		return err
	}

	e.Response.Header().Set("HX-Redirect", "/assets/"+priceDTO.AssetId)
	return e.JSON(200, map[string]any{"success": true})
}

func priceDetails(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	priceId := e.Request.PathValue("price_id")

	price, err := db.GetPriceById(priceId)
	if err != nil {
		fmt.Println("Error retrieving price:", err)
		return err
	}

	return sendPage(e, components.PricePage(price))
}

func priceUpdate(e *core.RequestEvent) error {
	db := database.NewDatabaseRepo(e.App)
	priceId := e.Request.PathValue("price_id")

	priceDTO := models.PriceUpdateDTO{}
	if err := e.BindBody(&priceDTO); err != nil {
		return err
	}

	price, err := db.GetPriceById(priceId)
	if err != nil {
		return err
	}

	price.Comment = priceDTO.Comment
	price.Updated = models.GenerateTimestamp()

	if err := db.UpdatePrice(price); err != nil {
		return err
	}

	return e.JSON(200, map[string]any{"success": true})
}
