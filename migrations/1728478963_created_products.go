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
			"id": "gpvjyk1s8lipcqo",
			"created": "2024-10-09 13:02:42.715Z",
			"updated": "2024-10-09 13:02:42.715Z",
			"name": "products",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "x8n2fv7t",
					"name": "name",
					"type": "text",
					"required": true,
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
					"id": "jceepjq0",
					"name": "description",
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
					"id": "prnqsamb",
					"name": "expectedPrice",
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
					"id": "ii0yxz7n",
					"name": "provider",
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
					"id": "pnzgatb7",
					"name": "tagIds",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "fpzmqrtemgkwcs6",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 32,
						"displayFields": null
					}
				}
			],
			"indexes": [],
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

		collection, err := dao.FindCollectionByNameOrId("gpvjyk1s8lipcqo")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
