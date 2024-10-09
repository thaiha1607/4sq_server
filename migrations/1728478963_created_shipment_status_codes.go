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
			"id": "inophh2p3nyvq5c",
			"created": "2024-10-09 13:02:42.716Z",
			"updated": "2024-10-09 13:02:42.716Z",
			"name": "shipment_status_codes",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jqcy8b80",
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
					"id": "fgxdixiy",
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_tbnOBcE` + "`" + ` ON ` + "`" + `shipment_status_codes` + "`" + ` (` + "`" + `statusCode` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("inophh2p3nyvq5c")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
