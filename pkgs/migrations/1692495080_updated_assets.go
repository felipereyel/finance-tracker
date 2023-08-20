package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1rm3ii9vtitvzs4")
		if err != nil {
			return err
		}

		// update
		edit_wallet := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eertwpzm",
			"name": "wallet",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "23wj5do6cr2xwfa",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), edit_wallet)
		collection.Schema.AddField(edit_wallet)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1rm3ii9vtitvzs4")
		if err != nil {
			return err
		}

		// update
		edit_wallet := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eertwpzm",
			"name": "wallet",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "23wj5do6cr2xwfa",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), edit_wallet)
		collection.Schema.AddField(edit_wallet)

		return dao.SaveCollection(collection)
	})
}
