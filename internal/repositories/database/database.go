package database

import (
	"fintracker/internal/models"

	"github.com/pocketbase/pocketbase/core"
)

type database struct {
	app core.App
}

func NewDatabaseRepo(app core.App) database {
	return database{app}
}

// func (db *database) CreateTask(task models.Task) error {
// 	query := `INSERT INTO tasks (id, title, description, owner_id) VALUES (?, ?, ?, ?)`
// 	_, err := db.conn.Exec(query, task.Id, task.Title, task.Description, task.OwnerId)
// 	return err
// }

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

var ListWalletsQuery = `
	SELECT id, name FROM wallets
`

func (db *database) ListWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	if err := db.app.DB().NewQuery(ListWalletsQuery).All(&wallets); err != nil {
		return nil, err
	}

	return wallets, nil
}

// func (db *database) GetCurrentSummary(wallet string) ([]models.Asset, error) {
// 	baseQuery := db.app.DB().Select(models.AssetFields...).From("assets").AndWhere(dbx.HashExp{"sell_date": ""})
// 	if wallet != "" {
// 		baseQuery = baseQuery.AndWhere(dbx.HashExp{"wallet": wallet})
// 	}

// 	var assets []models.Asset
// 	err := baseQuery.All(&assets)
// 	return assets, err
// }

// func (db *database) RetrieveTaskById(taskId string) (models.Task, error) {
// 	query := `SELECT id, title, owner_id, description FROM tasks WHERE id = ?`
// 	row := db.conn.QueryRow(query, taskId)

// 	var task models.Task
// 	err := row.Scan(&task.Id, &task.Title, &task.OwnerId, &task.Description)
// 	if err != nil {
// 		return models.EmptyTask, err
// 	}

// 	return task, nil
// }

// func (db *database) DeleteTask(taskId string) error {
// 	query := `DELETE FROM tasks WHERE id = ?`
// 	_, err := db.conn.Exec(query, taskId)
// 	return err
// }

// func (db *database) UpdateTask(task models.Task) error {
// 	query := `UPDATE tasks SET title = ?, description = ? WHERE id = ?`
// 	_, err := db.conn.Exec(query, task.Title, task.Description, task.Id)
// 	return err
// }

// func (db *database) InsertUser(user models.User) error {
// 	query := `INSERT INTO users (id, username, pswd_hash) VALUES (?, ?, ?)`
// 	_, err := db.conn.Exec(query, user.ID, user.Username, user.PswdHash)
// 	return err
// }

// func (db *database) RetrieveUserByName(username string) (models.User, error) {
// 	query := `SELECT id, username, pswd_hash FROM users WHERE username = ?`
// 	row := db.conn.QueryRow(query, username)

// 	var user models.User
// 	err := row.Scan(&user.ID, &user.Username, &user.PswdHash)
// 	if err != nil {
// 		return models.EmptyUser, err
// 	}

// 	return user, nil
// }

// func (db *database) RetrieveUserById(id string) (models.User, error) {
// 	query := `SELECT id, username, pswd_hash FROM users WHERE id = ?`
// 	row := db.conn.QueryRow(query, id)

// 	var user models.User
// 	err := row.Scan(&user.ID, &user.Username, &user.PswdHash)
// 	if err != nil {
// 		return models.EmptyUser, err
// 	}

// 	return user, nil
// }
