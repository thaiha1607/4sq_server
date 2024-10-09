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
			"id": "bd0d5gtacvvnnys",
			"created": "2024-10-09 13:02:42.712Z",
			"updated": "2024-10-09 13:02:42.712Z",
			"name": "conversations",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "modqcq1b",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 256,
						"pattern": ""
					}
				}
			],
			"indexes": [],
			"listRule": "participants_via_conversationId.userId ?= @request.auth.id || conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"viewRule": "participants_via_conversationId.userId ?= @request.auth.id || conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"deleteRule": "conversation_admins_via_conversationId.userId ?= @request.auth.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("bd0d5gtacvvnnys")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
