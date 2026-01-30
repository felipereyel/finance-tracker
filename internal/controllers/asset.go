package controllers

import (
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"
)

type assetController struct {
	db database.Database
}

func (controller assetController) GetAssetAggregate(assetId string) (models.AssetAggregate, error) {
	return controller.db.GetAssetAggregateById(assetId)
}

func (controller assetController) UpdateAsset(assetId string, dto models.AssetUpdateDTO) error {
	asset, err := controller.db.GetAssetById(assetId)
	if err != nil {
		return err
	}

	asset.SellDate = dto.SellDate
	asset.Comment = dto.Comment
	asset.Updated = models.GenerateTimestamp()

	return controller.db.UpdateAsset(asset)
}

func (controller assetController) CreateAsset(dto models.AssetCreateDTO) (string, error) {
	newAsset := models.CreateNewAsset(dto)
	if err := controller.db.CreateAsset(newAsset); err != nil {
		return "", err
	}

	priceDTO := models.PriceCreateDTO{
		Value:    newAsset.InitialPrice,
		LoggedAt: newAsset.BuyDate,
		Comment:  "Initial price",
	}

	newPrice := models.CreateNewPrice(newAsset.Id, priceDTO)
	if err := controller.db.CreatePrice(newPrice); err != nil {
		return "", err
	}

	return newAsset.Id, nil
}

func (controller assetController) SummarizeAssets(userId string, walletFilter string, typeFilter string) (models.Summary, error) {
	summary := models.Summary{
		Total:      0,
		Aggregates: make([]models.AssetAggregate, 0),

		AssetTypes:   models.AssetTypes,
		SelectedType: typeFilter,

		SelectedWallet: walletFilter,
		Wallets:        make([][]string, 0),
	}

	aggregated, err := controller.db.ListAssetAggregates(userId)
	if err != nil {
		return models.EmptySummary, err
	}

	for _, asset := range aggregated {
		if walletFilter != "" && asset.Wallet != walletFilter {
			continue
		}

		if typeFilter != "" && asset.Type != typeFilter {
			continue
		}

		summary.Aggregates = append(summary.Aggregates, asset)
		summary.Total += asset.LastPrice
	}

	wallets, err := controller.db.ListWallets(userId)
	if err != nil {
		return models.EmptySummary, err
	}

	for _, wallet := range wallets {
		summary.Wallets = append(summary.Wallets, []string{wallet.Id, wallet.Name})
	}

	return summary, nil
}

func (controller assetController) GetAssetOptions(userId string) (models.NewAssetOptions, error) {
	wallets, err := controller.db.ListWallets(userId)
	if err != nil {
		return models.EmptyNewAssetOptions, err
	}

	options := models.NewAssetOptions{
		AssetTypes: models.AssetTypes,
		Wallets:    make([][]string, 0),
	}

	for _, wallet := range wallets {
		options.Wallets = append(options.Wallets, []string{wallet.Id, wallet.Name})
	}

	return options, nil
}
