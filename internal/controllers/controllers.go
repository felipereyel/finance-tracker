package controllers

import (
	"fintracker/internal/repositories/database"
)

type Controllers struct {
	Price priceController
	Asset assetController
}

func NewControllers(db database.Database) Controllers {
	return Controllers{
		Price: priceController{db},
		Asset: assetController{db},
	}
}
