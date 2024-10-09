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
			"id": "fpzmqrtemgkwcs6",
			"created": "2024-10-09 13:02:42.716Z",
			"updated": "2024-10-09 13:02:42.716Z",
			"name": "tags",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "vfnvroid",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 64,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_ZcKxUtS` + "`" + ` ON ` + "`" + `tags` + "`" + ` (` + "`" + `name` + "`" + `)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.role = \"manager\"",
			"updateRule": null,
			"deleteRule": "@request.auth.role = \"manager\"",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fpzmqrtemgkwcs6")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
