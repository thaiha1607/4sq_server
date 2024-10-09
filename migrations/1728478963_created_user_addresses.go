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
			"id": "4mc1mg16zw31tum",
			"created": "2024-10-09 13:02:42.717Z",
			"updated": "2024-10-09 13:02:42.717Z",
			"name": "user_addresses",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "hsz5thp5",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"home",
							"work",
							"billing",
							"shipping",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "fqkslwn0",
					"name": "friendlyName",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 64,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "3yjyqbnv",
					"name": "isDefault",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "vlr8r3cp",
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
					"id": "zkamkzp9",
					"name": "addressId",
					"type": "relation",
					"required": true,
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
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_iLfx8Tu` + "`" + ` ON ` + "`" + `user_addresses` + "`" + ` (\n  ` + "`" + `userId` + "`" + `,\n  ` + "`" + `addressId` + "`" + `\n)"
			],
			"listRule": "userId = @request.auth.id",
			"viewRule": "userId = @request.auth.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "userId = @request.auth.id",
			"deleteRule": "userId = @request.auth.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4mc1mg16zw31tum")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
