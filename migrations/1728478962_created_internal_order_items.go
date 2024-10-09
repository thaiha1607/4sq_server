package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "gsc7xs8i2q990pg",
			"created": "2024-10-09 13:02:42.713Z",
			"updated": "2024-10-09 13:02:42.713Z",
			"name": "internal_order_items",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "k33papdq",
					"name": "qty",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "6yjchw0j",
					"name": "rollQty",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "xhyfdg6v",
					"name": "note",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "ypmd9q9s",
					"name": "internalOrderId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "g0y7nfa8ommcv1h",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ad5rzbhz",
					"name": "orderItemId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "nqclqjsjbs7e523",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_3TMYpsm` + "`" + ` ON ` + "`" + `internal_order_items` + "`" + ` (\n  ` + "`" + `internalOrderId` + "`" + `,\n  ` + "`" + `orderItemId` + "`" + `\n)"
			],
			"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.internalOrderId.rootOrderId = @request.data.orderItemId.orderId",
			"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.internalOrderId = internalOrderId && @request.data.orderItemId = orderItemId",
			"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("gsc7xs8i2q990pg")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
