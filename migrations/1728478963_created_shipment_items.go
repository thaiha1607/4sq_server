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
			"id": "st23val23amwoi5",
			"created": "2024-10-09 13:02:42.716Z",
			"updated": "2024-10-09 13:02:42.716Z",
			"name": "shipment_items",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "h6ihqpu6",
					"name": "qty",
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
					"id": "5i06dmud",
					"name": "rollQty",
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
					"id": "l9vhuhge",
					"name": "shipmentId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "0fvjhfcs9ig8nd8",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "im7wnzr1",
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
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_1xHG7KZ` + "`" + ` ON ` + "`" + `shipment_items` + "`" + ` (\n  ` + "`" + `shipmentId` + "`" + `,\n  ` + "`" + `orderItemId` + "`" + `\n)"
			],
			"listRule": "shipmentId.orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"viewRule": "shipmentId.orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.shipmentId.orderId = @request.data.orderItemId.orderId",
			"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.shipmentId = shipmentId && @request.data.orderItemId = orderItemId",
			"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("st23val23amwoi5")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
