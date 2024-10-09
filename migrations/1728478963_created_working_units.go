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
			"id": "cr8iizw10m7kh1d",
			"created": "2024-10-09 13:02:42.717Z",
			"updated": "2024-10-09 13:02:42.717Z",
			"name": "working_units",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "o9gekwse",
					"name": "name",
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
					"id": "xn47i6tj",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"warehouse",
							"office",
							"delivery",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "ry1rm3ou",
					"name": "imageUrl",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "jn427lzh",
					"name": "addressId",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "7uhp3ed4afuq9uc",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"createRule": "@request.auth.role = \"manager\"",
			"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"deleteRule": "@request.auth.role = \"manager\"",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("cr8iizw10m7kh1d")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
