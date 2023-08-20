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

		collection, err := dao.FindCollectionByNameOrId("t6z60bv72qu6riz")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT a.id, a.name, a.type, a.initial_price, a.buy_date, a.comment, ap.value as latest_price, ap.logged_at as latest_date, a.sell_date, a.wallet\nFROM assets a, asset_prices ap\nWHERE a.id = ap.asset_id\nAND ap.logged_at = (SELECT MAX(logged_at) FROM asset_prices WHERE asset_id = a.id)"
		}`), &options)
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("b63svvhq")

		// remove
		collection.Schema.RemoveField("9tu3ywgv")

		// remove
		collection.Schema.RemoveField("wesvd4e9")

		// remove
		collection.Schema.RemoveField("awo9wzew")

		// remove
		collection.Schema.RemoveField("uqeffpqw")

		// remove
		collection.Schema.RemoveField("pjorq9gx")

		// remove
		collection.Schema.RemoveField("v4beuylm")

		// remove
		collection.Schema.RemoveField("4af6dedi")

		// add
		new_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bmafibcb",
			"name": "name",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_name)
		collection.Schema.AddField(new_name)

		// add
		new_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tricsfrz",
			"name": "type",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_type)
		collection.Schema.AddField(new_type)

		// add
		new_initial_price := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "cfox5p1q",
			"name": "initial_price",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_initial_price)
		collection.Schema.AddField(new_initial_price)

		// add
		new_buy_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "vuvec2d3",
			"name": "buy_date",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_buy_date)
		collection.Schema.AddField(new_buy_date)

		// add
		new_comment := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "n1loy8sk",
			"name": "comment",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_comment)
		collection.Schema.AddField(new_comment)

		// add
		new_latest_price := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "4jfbprxq",
			"name": "latest_price",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_latest_price)
		collection.Schema.AddField(new_latest_price)

		// add
		new_latest_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "eg6v257j",
			"name": "latest_date",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_latest_date)
		collection.Schema.AddField(new_latest_date)

		// add
		new_sell_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7zmo9p5x",
			"name": "sell_date",
			"type": "date",
			"required": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_sell_date)
		collection.Schema.AddField(new_sell_date)

		// add
		new_wallet := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "g68jzypr",
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
		}`), new_wallet)
		collection.Schema.AddField(new_wallet)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("t6z60bv72qu6riz")
		if err != nil {
			return err
		}

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT a.id, a.name, a.type, a.initial_price, a.buy_date, a.comment, ap.value as latest_price, ap.logged_at as latest_date, a.sell_date\nFROM assets a, asset_prices ap\nWHERE a.id = ap.asset_id\nAND ap.logged_at = (SELECT MAX(logged_at) FROM asset_prices WHERE asset_id = a.id)"
		}`), &options)
		collection.SetOptions(options)

		// add
		del_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "b63svvhq",
			"name": "name",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_name)
		collection.Schema.AddField(del_name)

		// add
		del_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "9tu3ywgv",
			"name": "type",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_type)
		collection.Schema.AddField(del_type)

		// add
		del_initial_price := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wesvd4e9",
			"name": "initial_price",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), del_initial_price)
		collection.Schema.AddField(del_initial_price)

		// add
		del_buy_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "awo9wzew",
			"name": "buy_date",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_buy_date)
		collection.Schema.AddField(del_buy_date)

		// add
		del_comment := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "uqeffpqw",
			"name": "comment",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_comment)
		collection.Schema.AddField(del_comment)

		// add
		del_latest_price := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "pjorq9gx",
			"name": "latest_price",
			"type": "number",
			"required": true,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), del_latest_price)
		collection.Schema.AddField(del_latest_price)

		// add
		del_latest_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "v4beuylm",
			"name": "latest_date",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_latest_date)
		collection.Schema.AddField(del_latest_date)

		// add
		del_sell_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "4af6dedi",
			"name": "sell_date",
			"type": "date",
			"required": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_sell_date)
		collection.Schema.AddField(del_sell_date)

		// remove
		collection.Schema.RemoveField("bmafibcb")

		// remove
		collection.Schema.RemoveField("tricsfrz")

		// remove
		collection.Schema.RemoveField("cfox5p1q")

		// remove
		collection.Schema.RemoveField("vuvec2d3")

		// remove
		collection.Schema.RemoveField("n1loy8sk")

		// remove
		collection.Schema.RemoveField("4jfbprxq")

		// remove
		collection.Schema.RemoveField("eg6v257j")

		// remove
		collection.Schema.RemoveField("7zmo9p5x")

		// remove
		collection.Schema.RemoveField("g68jzypr")

		return dao.SaveCollection(collection)
	})
}
