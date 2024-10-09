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
			"id": "dmlx8wh39pu41u8",
			"created": "2024-10-09 13:02:42.717Z",
			"updated": "2024-10-09 13:02:42.717Z",
			"name": "warehouse_assignments",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "3v5lcqbt",
					"name": "status",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"pending",
							"assigned",
							"in_progress",
							"completed",
							"cancelled",
							"failed",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "gvim9fk7",
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
				},
				{
					"system": false,
					"id": "uk9hplvt",
					"name": "staffId",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "v6f615blwqyoe4d",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "rji4r1bl",
					"name": "internalOrderId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "g0y7nfa8ommcv1h",
						"cascadeDelete": true,
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("dmlx8wh39pu41u8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
