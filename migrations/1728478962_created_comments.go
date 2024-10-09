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
			"id": "yy58rkdg4k4kjs5",
			"created": "2024-10-09 13:02:42.712Z",
			"updated": "2024-10-09 13:02:42.712Z",
			"name": "comments",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "nekxjsfj",
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
					"id": "gjmj5uuz",
					"name": "content",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 512,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "jp5oc0o4",
					"name": "productId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "gpvjyk1s8lipcqo",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ucczphj7",
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_KAmzy3O` + "`" + ` ON ` + "`" + `comments` + "`" + ` (\n  ` + "`" + `productId` + "`" + `,\n  ` + "`" + `userId` + "`" + `\n)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != \"\" && @request.auth.role = 'customer'",
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("yy58rkdg4k4kjs5")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
