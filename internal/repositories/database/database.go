package database

import (
	"fintracker/internal/models"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type database struct {
	app core.App
}

func NewDatabaseRepo(app core.App) database {
	return database{app}
}

func (db *database) CreateAsset(asset models.Asset) error {
	result, err := db.app.DB().Insert("assets", dbx.Params{
		"id":            asset.Id,
		"name":          asset.Name,
		"type":          asset.Type,
		"wallet":        asset.Wallet,
		"initial_price": asset.InitialPrice,
		"buy_date":      asset.BuyDate,
		"comment":       asset.Comment,
	}).Execute()

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when inserting asset")
	}

	return nil
}

var GetCurrentSummaryQuery = `
	SELECT 
		a.id,
		a.name,
		a.type,
		a.wallet,
		w.name as wallet_name,
		a.initial_price,
		a.buy_date,
		ap.value as last_price,
		ap.logged_at as last_date
	FROM 
		assets a,
		asset_prices ap,
		wallets w
	WHERE 
		a.sell_date = '' AND
		a.id = ap.asset_id AND 
		a.wallet = w.id AND
		ap.logged_at = (SELECT MAX(logged_at) FROM asset_prices WHERE asset_id = a.id)
	ORDER BY
		w.name,
		ap.logged_at DESC
`

func (db *database) GetCurrentSummary(wallet string, asset_type string) ([]models.AssetAggregate, error) {
	var assets []models.AssetAggregate
	if err := db.app.DB().NewQuery(GetCurrentSummaryQuery).All(&assets); err != nil {
		return nil, err
	}

	var filtered []models.AssetAggregate
	for _, asset := range assets {
		if wallet != "" && asset.Wallet != wallet {
			continue
		}

		if asset_type != "" && asset.Type != asset_type {
			continue
		}

		filtered = append(filtered, asset)
	}

	return filtered, nil
}

func (db *database) ListWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	if err := db.app.DB().Select("id", "name").From("wallets").All(&wallets); err != nil {
		return nil, err
	}

	return wallets, nil
}

func (db *database) CreatePrice(price models.Price) error {
	result, err := db.app.DB().Insert("asset_prices", dbx.Params{
		"id":        price.Id,
		"asset_id":  price.AssetId,
		"value":     price.Value,
		"logged_at": price.Logged,
		"gain":      price.Gain,
		"comment":   price.Comment,
	}).Execute()

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when inserting price")
	}

	return nil
}
