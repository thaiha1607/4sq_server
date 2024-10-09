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
			"id": "kgkfehwran8zgs8",
			"created": "2024-10-09 13:02:42.717Z",
			"updated": "2024-10-09 13:02:42.717Z",
			"name": "transaction_history",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "xy2z6b67",
					"name": "entityType",
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
					"id": "ubiiz0ft",
					"name": "entityId",
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
					"id": "cprxxoqe",
					"name": "statusCodeId",
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
					"id": "hycktg06",
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
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_kcKC4Vm` + "`" + ` ON ` + "`" + `transaction_history` + "`" + ` (\n  ` + "`" + `entityType` + "`" + `,\n  ` + "`" + `entityId` + "`" + `\n)"
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("kgkfehwran8zgs8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
