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
			"id": "bg69cxt80nkfws1",
			"created": "2024-10-09 13:02:42.715Z",
			"updated": "2024-10-09 13:02:42.715Z",
			"name": "product_quantity_history",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jfgew7gh",
					"name": "categoryId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "upkshg4h89ndt95",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "jwg5tfn2",
					"name": "amountOfChange",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": true
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.role = \"manager\"",
			"viewRule": "@request.auth.role = \"manager\"",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("bg69cxt80nkfws1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
