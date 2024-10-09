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
			"id": "7uhp3ed4afuq9uc",
			"created": "2024-10-09 13:02:42.711Z",
			"updated": "2024-10-09 13:02:42.711Z",
			"name": "addresses",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "rbv1jzdm",
					"name": "line1",
					"type": "text",
					"required": true,
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
					"id": "ncmw5m13",
					"name": "line2",
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
					"id": "e9jojbr1",
					"name": "city",
					"type": "text",
					"required": true,
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
					"id": "anqbruir",
					"name": "state",
					"type": "text",
					"required": true,
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
					"id": "sul3iaab",
					"name": "country",
					"type": "text",
					"required": true,
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
					"id": "pc4cgzj6",
					"name": "zipOrPostcode",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 256,
						"pattern": ""
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\"",
			"viewRule": "@request.auth.id != \"\"",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": null,
			"deleteRule": "@request.auth.id != \"\"",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("7uhp3ed4afuq9uc")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
