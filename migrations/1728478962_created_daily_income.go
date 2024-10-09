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
			"id": "90dgpouuugdj3g8",
			"created": "2024-10-09 13:02:42.712Z",
			"updated": "2024-10-09 13:02:42.712Z",
			"name": "daily_income",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "obi8scyl",
					"name": "amountOfChange",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.role = \"manager\"",
			"viewRule": "@request.auth.role = \"manager\"",
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

		collection, err := dao.FindCollectionByNameOrId("90dgpouuugdj3g8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
