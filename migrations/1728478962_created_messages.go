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
			"id": "3ryn3wvszi2ny9q",
			"created": "2024-10-09 13:02:42.714Z",
			"updated": "2024-10-09 13:02:42.714Z",
			"name": "messages",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zxx9i3u1",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"text",
							"image",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "ljeykjul",
					"name": "content",
					"type": "text",
					"required": true,
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
					"id": "l30xjrgx",
					"name": "participantId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "nhuyrnedwhw2jm8",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "l3nskghf",
					"name": "recipientIds",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "nhuyrnedwhw2jm8",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 32,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.participants_via_userId.conversationId ?= participantId.conversationId",
			"viewRule": "@request.auth.participants_via_userId.conversationId ?= participantId.conversationId",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "participantId.userId = @request.auth.id || @request.auth.conversation_admins_via_userId.conversationId ?= participantId.conversationId",
			"deleteRule": "participantId.userId = @request.auth.id || @request.auth.conversation_admins_via_userId.conversationId ?= participantId.conversationId",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3ryn3wvszi2ny9q")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
