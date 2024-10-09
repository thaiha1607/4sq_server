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
			"id": "2stc36eglv3voqe",
			"created": "2024-10-09 13:02:42.714Z",
			"updated": "2024-10-09 13:02:42.714Z",
			"name": "orders",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "vfmtk9sh",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"sale",
							"return",
							"exchange",
							"transfer",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "z610saqw",
					"name": "priority",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 100,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "lvsuqylu",
					"name": "rating",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 5,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "brnqocq4",
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
					"id": "ff8dl790",
					"name": "otherInfo",
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
					"id": "uh4rmuha",
					"name": "rootOrderId",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "2stc36eglv3voqe",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "eumiqg6p",
					"name": "creatorId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "hy0l5eky",
					"name": "guestId",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "68wsl8pf0ab44u1",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "vgowsnxd",
					"name": "statusCodeId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "93h7enu78ajtk7w",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "tp4tozu9",
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
				"CREATE INDEX ` + "`" + `idx_3Q3JUsm` + "`" + ` ON ` + "`" + `orders` + "`" + ` (` + "`" + `creatorId` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_f4G2ru9` + "`" + ` ON ` + "`" + `orders` + "`" + ` (` + "`" + `guestId` + "`" + `)"
			],
			"listRule": "creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"viewRule": "creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "(creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")) && @request.data.creatorId = creatorId && @request.data.type = type",
			"deleteRule": "creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("2stc36eglv3voqe")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
