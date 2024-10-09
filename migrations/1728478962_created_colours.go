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
			"id": "3o2zbir67gxkcni",
			"created": "2024-10-09 13:02:42.712Z",
			"updated": "2024-10-09 13:02:42.712Z",
			"name": "colours",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "xflmf0yq",
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
					"id": "2kpyfojn",
					"name": "hexCode",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("3o2zbir67gxkcni")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
