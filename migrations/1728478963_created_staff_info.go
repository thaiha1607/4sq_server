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
			"id": "v6f615blwqyoe4d",
			"created": "2024-10-09 13:02:42.716Z",
			"updated": "2024-10-09 13:02:42.716Z",
			"name": "staff_info",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "lpq224tk",
					"name": "statusCode",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"active",
							"inactive",
							"suspended",
							"terminated",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "thsgvtek",
					"name": "role",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"salesperson",
							"warehouse",
							"delivery",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "lbgtar1m",
					"name": "userId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "kbmzxgda",
					"name": "workingUnitId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "cr8iizw10m7kh1d",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_kikHSGl` + "`" + ` ON ` + "`" + `staff_info` + "`" + ` (` + "`" + `userId` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_r696cdI` + "`" + ` ON ` + "`" + `staff_info` + "`" + ` (` + "`" + `role` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"createRule": "@request.auth.role = \"manager\"",
			"updateRule": "@request.auth.role = \"manager\"",
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

		collection, err := dao.FindCollectionByNameOrId("v6f615blwqyoe4d")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
