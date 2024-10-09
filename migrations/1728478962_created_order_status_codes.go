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
			"id": "93h7enu78ajtk7w",
			"created": "2024-10-09 13:02:42.714Z",
			"updated": "2024-10-09 13:02:42.714Z",
			"name": "order_status_codes",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zgwsoimk",
					"name": "statusCode",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "3jemdow9",
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
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_YhmYEHX` + "`" + ` ON ` + "`" + `order_status_codes` + "`" + ` (` + "`" + `statusCode` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\"",
			"viewRule": "@request.auth.id != \"\"",
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

		collection, err := dao.FindCollectionByNameOrId("93h7enu78ajtk7w")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
