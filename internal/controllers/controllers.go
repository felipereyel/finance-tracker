package controllers

import (
	"fintracker/internal/repositories/database"
)

type Controllers struct {
	Price priceController
	Asset assetController
	User  userController
}

func NewControllers(db database.Database) Controllers {
	return Controllers{
		Price: priceController{db},
		Asset: assetController{db},
		User:  userController{db},
	}
}
