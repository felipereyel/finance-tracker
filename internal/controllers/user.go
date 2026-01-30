package controllers

import (
	"fintracker/internal/repositories/database"
)

type userController struct {
	db database.Database
}

func (controller userController) GetUserIdFromCredentials(email string, password string) (string, error) {
	return controller.db.GetUserIdFromCredentials(email, password)
}

func (controller userController) ChechUserOwnsWallet(userId string, walletId string) error {
	return controller.db.ChechUserOwnsWallet(userId, walletId)
}

func (controller userController) ChechUserOwnsAsset(userId string, assetId string) error {
	return controller.db.ChechUserOwnsAsset(userId, assetId)
}

func (controller userController) ChechUserOwnsPrice(userId string, priceId string) error {
	return controller.db.ChechUserOwnsPrice(userId, priceId)
}
