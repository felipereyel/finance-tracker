package controllers

import (
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"
)

type priceController struct {
	db database.Database
}

func (controller priceController) GetPrice(priceId string) (models.Price, error) {
	price, err := controller.db.GetPriceById(priceId)
	if err != nil {
		return models.EmptyPrice, err
	}

	return price, nil
}

func (controller priceController) UpdatePrice(priceId string, dto models.PriceUpdateDTO) error {
	price, err := controller.db.GetPriceById(priceId)
	if err != nil {
		return err
	}

	price.Comment = dto.Comment
	price.Updated = models.GenerateTimestamp()

	return controller.db.UpdatePrice(price)
}

func (controller priceController) ListPrices(assetId string) ([]models.Price, error) {
	return controller.db.ListPricesByAssetId(assetId)
}

func (controller priceController) CreatePrice(assetId string, dto models.PriceCreateDTO) error {
	newPrice := models.CreateNewPrice(assetId, dto)
	return controller.db.CreatePrice(newPrice)
}

func (controller priceController) ListPricesEnrich(assetId string) (models.AssetAggregate, []models.Price, error) {
	assetAgg, err := controller.db.GetAssetAggregateById(assetId)
	if err != nil {
		return models.EmptyAssetAggregate, nil, err
	}

	prices, err := controller.db.ListPricesByAssetId(assetId)
	if err != nil {
		return models.EmptyAssetAggregate, nil, err
	}

	return assetAgg, prices, nil
}
