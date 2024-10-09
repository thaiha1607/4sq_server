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
			"id": "nhuyrnedwhw2jm8",
			"created": "2024-10-09 13:02:42.715Z",
			"updated": "2024-10-09 13:02:42.715Z",
			"name": "participants",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ly2ky3cr",
					"name": "conversationId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "bd0d5gtacvvnnys",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "bwwxre69",
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
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_bJZA4H8` + "`" + ` ON ` + "`" + `participants` + "`" + ` (\n  ` + "`" + `conversationId` + "`" + `,\n  ` + "`" + `userId` + "`" + `\n)"
			],
			"listRule": "conversationId.participants_via_conversationId.userId ?= @request.auth.id || conversationId.conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"viewRule": "conversationId.participants_via_conversationId.userId ?= @request.auth.id || conversationId.conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": null,
			"deleteRule": "conversationId.participants_via_conversationId.userId ?= @request.auth.id || conversationId.conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nhuyrnedwhw2jm8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
