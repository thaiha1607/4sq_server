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
			"id": "1hln2do3ca9mwaz",
			"created": "2024-10-09 13:02:42.717Z",
			"updated": "2024-10-09 13:02:42.736Z",
			"name": "product_quantity_summary",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ry8hwknq",
					"name": "totalQty",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 1
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\"",
			"viewRule": "@request.auth.id != \"\"",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT categoryId as id, SUM(qty) AS totalQty\nFROM product_quantities\nGROUP BY categoryId;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("1hln2do3ca9mwaz")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
