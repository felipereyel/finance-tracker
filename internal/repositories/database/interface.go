package database

import "fintracker/internal/models"

type Database interface {
	// RetrieveTaskById(taskId string) (models.Task, error)
	// UpdateTask(task models.Task) error

	CreateAsset(asset models.Asset) error
	ListAssets(wallet string) ([]models.Asset, error)

	CreatePrice(price models.Price) error

	// InsertUser(user models.User) error
	// RetrieveUserById(id string) (models.User, error)
	// RetrieveUserByName(username string) (models.User, error)
}
