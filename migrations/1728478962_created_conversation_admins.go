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
			"id": "eblguzfwjtr5n6s",
			"created": "2024-10-09 13:02:42.712Z",
			"updated": "2024-10-09 13:02:42.712Z",
			"name": "conversation_admins",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zlmqq2nq",
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
					"id": "5rzut65t",
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_PAsZnx1` + "`" + ` ON ` + "`" + `conversation_admins` + "`" + ` (\n  ` + "`" + `conversationId` + "`" + `,\n  ` + "`" + `userId` + "`" + `\n)"
			],
			"listRule": "conversationId.participants_via_conversationId.userId ?= @request.auth.id || conversationId.conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"viewRule": "conversationId.participants_via_conversationId.userId ?= @request.auth.id || conversationId.conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"createRule": "@request.data.conversationId.conversation_admins_via_conversationId.id:length < 1 || @request.auth.conversation_admins_via_userId.conversationId ?= @request.data.conversationId",
			"updateRule": null,
			"deleteRule": "userId = @request.auth.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("eblguzfwjtr5n6s")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
