package database

import "fintracker/internal/models"

type Database interface {
	// User methods
	GetUserIdFromCredentials(email string, password string) (string, error)
	ChechUserOwnsAsset(userId string, assetId string) error
	ChechUserOwnsPrice(userId string, assetId string) error

	// Wallet methods
	ListWallets(userId string) ([]models.Wallet, error)

	// Asset methods
	CreateAsset(asset models.Asset) error
	UpdateAsset(asset models.Asset) error
	GetAssetById(assetId string) (models.Asset, error)

	// AssetAggregate methods
	ListAssetAggregates(userId string) ([]models.AssetAggregate, error)
	GetAssetAggregateById(assetId string) (models.AssetAggregate, error)

	// Price methods
	CreatePrice(price models.Price) error
	UpdatePrice(price models.Price) error
	GetPriceById(priceId string) (models.Price, error)
	ListPricesByAssetId(assetId string) ([]models.Price, error)
}
