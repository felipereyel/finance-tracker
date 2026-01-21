package database

import (
	"fintracker/internal/models"

	_ "modernc.org/sqlite"
)

type fakeDatabase struct {
	wallets map[string]models.Wallet
	assets  map[string]models.Asset
	price   map[string]models.Price
}

func NewFakeDatabaseRepo() fakeDatabase {
	return fakeDatabase{
		wallets: make(map[string]models.Wallet),
		assets:  make(map[string]models.Asset),
		price:   make(map[string]models.Price),
	}
}
