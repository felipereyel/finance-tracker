package database

import "fintracker/internal/models"

type Database interface {
	// RetrieveTaskById(taskId string) (models.Task, error)
	// UpdateTask(task models.Task) error

	ListWallets() ([]models.Wallet, error)

	CreateAsset(asset models.Asset) error
	ListAssetAggregates(wallet string, assetType string) ([]models.AssetAggregate, error)
	GetAssetAggregateById(assetId string) (models.AssetAggregate, error)

	CreatePrice(price models.Price) error

	// InsertUser(user models.User) error
	// RetrieveUserById(id string) (models.User, error)
	// RetrieveUserByName(username string) (models.User, error)
}
