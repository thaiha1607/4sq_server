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
			"id": "jzxryofwy4tmpha",
			"created": "2024-10-09 13:02:42.713Z",
			"updated": "2024-10-09 13:02:42.713Z",
			"name": "invoice_line_items",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "t0bsxsim",
					"name": "invoiceId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "mbgvq8b9yf5i2d1",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "4xzbexgo",
					"name": "orderItemId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "nqclqjsjbs7e523",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "srijo8n1",
					"name": "unitPrice",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "wmuffpzt",
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
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_gS5MiNG` + "`" + ` ON ` + "`" + `invoice_line_items` + "`" + ` (\n  ` + "`" + `invoiceId` + "`" + `,\n  ` + "`" + `orderItemId` + "`" + `\n)"
			],
			"listRule": "invoiceId.orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"viewRule": "invoiceId.orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"createRule": "(@request.auth.role = \"manager\" || (@request.auth.id != '' && @request.data.invoiceId.type = 'pro_forma')) && @request.data.invoiceId.orderId = @request.data.orderItemId.orderId",
			"updateRule": "@request.auth.role = \"manager\" && @request.data.invoiceId = invoiceId && @request.data.orderItemId = orderItemId",
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

		collection, err := dao.FindCollectionByNameOrId("jzxryofwy4tmpha")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
