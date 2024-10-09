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
			"id": "sznmzh643y5pv5s",
			"created": "2024-10-09 13:02:42.715Z",
			"updated": "2024-10-09 13:02:42.715Z",
			"name": "product_images",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "cugfpyja",
					"name": "imageUrl",
					"type": "url",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "0xxipmnd",
					"name": "altText",
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
					"id": "lz2i7fzs",
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
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_NZuOJuB` + "`" + ` ON ` + "`" + `product_images` + "`" + ` (` + "`" + `productId` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("sznmzh643y5pv5s")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
