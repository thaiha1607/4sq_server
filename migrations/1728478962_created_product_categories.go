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
			"id": "upkshg4h89ndt95",
			"created": "2024-10-09 13:02:42.715Z",
			"updated": "2024-10-09 13:02:42.715Z",
			"name": "product_categories",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "vtp20sii",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 256,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "6apt3tel",
					"name": "productId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "gpvjyk1s8lipcqo",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "iehyx3oh",
					"name": "colourId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "3o2zbir67gxkcni",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_Xa76LkK` + "`" + ` ON ` + "`" + `product_categories` + "`" + ` (\n  ` + "`" + `productId` + "`" + `,\n  ` + "`" + `colourId` + "`" + `\n)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
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

		collection, err := dao.FindCollectionByNameOrId("upkshg4h89ndt95")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
