package database

import "fintracker/internal/models"

type Database interface {
	// CreateTask(task models.Task) error
	// RetrieveTaskById(taskId string) (models.Task, error)
	// UpdateTask(task models.Task) error
	// DeleteTask(taskId string) error
	ListAssets(wallet string) ([]models.Asset, error)

	// InsertUser(user models.User) error
	// RetrieveUserById(id string) (models.User, error)
	// RetrieveUserByName(username string) (models.User, error)
}
