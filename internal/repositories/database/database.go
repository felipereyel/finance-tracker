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

// User methods

func (db database) GetUserIdFromCredentials(email string, password string) (string, error) {
	user, err := db.app.FindAuthRecordByEmail("users", email)
	if err != nil {
		fmt.Println("auth record not found", err)
		return "", err
	}

	if !user.ValidatePassword(password) {
		fmt.Println("pwd not matched")
		return "", fmt.Errorf("invalid password")
	}

	return user.Id, nil
}

func (db database) ChechUserOwnsAsset(userId string, assetId string) error {
	query := db.app.DB().Select("wallets.user_id as user", "assets.id as asset").
		From("wallets").
		InnerJoin("assets", dbx.NewExp("wallets.id = assets.wallet")).
		Where(dbx.And(
			dbx.NewExp("wallets.user_id = {:userId}", dbx.Params{"userId": userId}),
			dbx.NewExp("assets.id = {:assetId}", dbx.Params{"assetId": assetId}),
		)).
		Build()

	data := struct {
		User  string
		Asset string
	}{}

	if err := query.One(&data); err != nil {
		return err
	}

	if data.Asset != assetId || data.User != userId {
		return fmt.Errorf("no rows affected when checking records")
	}

	return nil
}

func (db database) ChechUserOwnsPrice(userId string, priceId string) error {
	query := db.app.DB().Select("wallets.user_id as user", "asset_prices.id as price").
		From("wallets").
		InnerJoin("assets", dbx.NewExp("wallets.id = assets.wallet")).
		InnerJoin("asset_prices", dbx.NewExp("assets.id = asset_prices.asset_id")).
		Where(dbx.And(
			dbx.NewExp("wallets.user_id = {:userId}", dbx.Params{"userId": userId}),
			dbx.NewExp("asset_prices.id = {:priceId}", dbx.Params{"priceId": priceId}),
		))

	data := struct {
		User  string
		Price string
	}{}

	if err := query.Build().One(&data); err != nil {
		return err
	}

	if data.Price != priceId || data.User != userId {
		return fmt.Errorf("no rows affected when checking records")
	}

	return nil
}

// Wallet methods

func (db database) ListWallets(userId string) ([]models.Wallet, error) {
	query := db.app.DB().
		Select("id", "name").
		From("wallets").
		Where(dbx.NewExp("wallets.user_id = {:userId}", dbx.Params{"userId": userId}))

	var wallets []models.Wallet
	if err := query.All(&wallets); err != nil {
		return nil, err
	}

	return wallets, nil
}

// Asset methods

func (db database) CreateAsset(asset models.Asset) error {
	result, err := db.app.DB().Insert("assets", dbx.Params{
		"id":            asset.Id,
		"name":          asset.Name,
		"type":          asset.Type,
		"wallet":        asset.Wallet,
		"initial_price": asset.InitialPrice,
		"buy_date":      asset.BuyDate,
		"comment":       asset.Comment,
		"created":       asset.Created,
		"updated":       asset.Updated,
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

func (db database) UpdateAsset(asset models.Asset) error {
	result, err := db.app.DB().Update("assets", dbx.Params{
		"comment":   asset.Comment,
		"sell_date": asset.SellDate,
		"updated":   asset.Updated,
	}, dbx.HashExp{"id": asset.Id}).Execute()

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when updating asset")
	}

	return nil
}

func (db database) GetAssetById(assetId string) (models.Asset, error) {
	var asset models.Asset
	if err := db.app.DB().Select("id", "name", "type", "wallet", "initial_price", "buy_date", "sell_date", "comment", "created", "updated").From("assets").Where(dbx.HashExp{"id": assetId}).One(&asset); err != nil {
		return models.Asset{}, err
	}

	return asset, nil
}

// AssetAggregate methods

var AssetAggregatesSelectFragment = `
	SELECT
		a.id,
		a.name,
		a.type,
		a.wallet,
		w.name as wallet_name,
		a.initial_price,
		a.buy_date,
		ap.value as last_price,
		ap.logged_at as last_date,
		a.sell_date as sell_date,
		a.comment as comment
	FROM
		assets a,
		asset_prices ap,
		wallets w
	WHERE
		a.id = ap.asset_id AND
		a.wallet = w.id AND
		ap.logged_at = (SELECT MAX(logged_at) FROM asset_prices WHERE asset_id = a.id)
`

func (db database) ListAssetAggregates(userId string) ([]models.AssetAggregate, error) {
	var ListAssetAggregatesQuery = AssetAggregatesSelectFragment + ` AND w.user_id = {:id} AND a.sell_date = '' ORDER BY w.name, ap.logged_at DESC`

	var assets []models.AssetAggregate
	if err := db.app.DB().NewQuery(ListAssetAggregatesQuery).Bind(dbx.Params{"id": userId}).All(&assets); err != nil {
		return nil, err
	}

	return assets, nil
}

func (db database) GetAssetAggregateById(assetId string) (models.AssetAggregate, error) {
	var GetAssetAggregateByIdQuery = AssetAggregatesSelectFragment + ` AND a.id = {:id} LIMIT 1`
	var asset models.AssetAggregate
	if err := db.app.DB().NewQuery(GetAssetAggregateByIdQuery).Bind(dbx.Params{"id": assetId}).One(&asset); err != nil {
		return models.AssetAggregate{}, err
	}

	return asset, nil
}

// Price methods

func (db database) CreatePrice(price models.Price) error {
	result, err := db.app.DB().Insert("asset_prices", dbx.Params{
		"id":        price.Id,
		"asset_id":  price.AssetId,
		"value":     price.Value,
		"logged_at": price.LoggedAt,
		"gain":      price.Gain,
		"comment":   price.Comment,
		"created":   price.Created,
		"updated":   price.Updated,
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

func (db database) UpdatePrice(price models.Price) error {
	result, err := db.app.DB().Update("asset_prices", dbx.Params{
		"comment": price.Comment,
		"updated": price.Updated,
	}, dbx.HashExp{"id": price.Id}).Execute()
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected when updating price")
	}

	return nil
}

func (db database) GetPriceById(priceId string) (models.Price, error) {
	var price models.Price
	if err := db.app.DB().Select("id", "asset_id", "value", "logged_at", "gain", "comment", "created", "updated").From("asset_prices").Where(dbx.HashExp{"id": priceId}).One(&price); err != nil {
		return models.Price{}, err
	}

	return price, nil
}

func (db database) ListPricesByAssetId(assetId string) ([]models.Price, error) {
	var prices []models.Price
	if err := db.app.DB().Select("id", "asset_id", "value", "logged_at", "gain", "comment", "created", "updated").From("asset_prices").Where(dbx.HashExp{"asset_id": assetId}).OrderBy("logged_at DESC").All(&prices); err != nil {
		return nil, err
	}

	return prices, nil
}
