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
			"id": "nqclqjsjbs7e523",
			"created": "2024-10-09 13:02:42.714Z",
			"updated": "2024-10-09 13:02:42.714Z",
			"name": "order_items",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "iispyjhp",
					"name": "orderedQty",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "moj8swah",
					"name": "receivedQty",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "hygeyvnx",
					"name": "shippedQty",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "brpue4wf",
					"name": "assignedQty",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "rdcou6v1",
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
					"id": "doi97njz",
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
					"id": "7yeasgnz",
					"name": "orderId",
					"type": "relation",
					"required": true,
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
					"id": "netqqwib",
					"name": "productCategoryId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "upkshg4h89ndt95",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_LGPiKG4` + "`" + ` ON ` + "`" + `order_items` + "`" + ` (\n  ` + "`" + `orderId` + "`" + `,\n  ` + "`" + `productCategoryId` + "`" + `\n)"
			],
			"listRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"viewRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"createRule": "@request.auth.id != \"\" && (@request.data.assignedQty:isset = false || @request.data.assignedQty = 0) && (@request.data.shippedQty:isset = false || @request.data.shippedQty = 0) && (@request.data.receivedQty:isset = false || @request.data.receivedQty = 0)",
			"updateRule": "((orderId.creatorId = @request.auth.id && @request.data.assignedQty:isset = false && @request.data.shippedQty:isset = false && @request.data.receivedQty:isset = false) || (@request.auth.id != \"\" && @request.auth.role != \"customer\")) && @request.data.orderId = orderId && @request.data.productCategoryId = productCategoryId",
			"deleteRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nqclqjsjbs7e523")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
