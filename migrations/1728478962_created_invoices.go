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
			"id": "mbgvq8b9yf5i2d1",
			"created": "2024-10-09 13:02:42.713Z",
			"updated": "2024-10-09 13:02:42.713Z",
			"name": "invoices",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "0qohuax0",
					"name": "totalAmount",
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
					"id": "njtfqifi",
					"name": "paidAmount",
					"type": "number",
					"required": false,
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
					"id": "g1lqvenw",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"pro_forma",
							"regular",
							"past_due",
							"retainer",
							"interim",
							"timesheet",
							"final",
							"credit",
							"debit",
							"mixed",
							"commercial",
							"recurring",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "w4lktdsd",
					"name": "paymentMethod",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"cash",
							"eft",
							"gift_card",
							"credit_card",
							"debit_card",
							"prepaid_card",
							"check",
							"other"
						]
					}
				},
				{
					"system": false,
					"id": "grrtgg9g",
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
					"id": "tejl6v7q",
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
					"id": "0elcxjzq",
					"name": "statusCodeId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "ldiepkv7gnn4nrr",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "xguda16b",
					"name": "rootInvoiceId",
					"type": "relation",
					"required": false,
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
					"id": "ek3zcke2",
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
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_VlOtvkK` + "`" + ` ON ` + "`" + `invoices` + "`" + ` (` + "`" + `orderId` + "`" + `)"
			],
			"listRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"viewRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
			"createRule": "(@request.auth.role = \"manager\" && (@request.data.paidAmount:isset = false || @request.data.paidAmount <= totalAmount)) || (@request.auth.id != \"\" && @request.data.type = \"pro_forma\" && @request.data.paidAmount:isset = false)",
			"updateRule": "(@request.auth.staff_info_via_userId.role ?= \"delivery\" || @request.auth.role = \"manager\") && @request.data.orderId = orderId && (@request.data.paidAmount:isset = false || @request.data.paidAmount <= totalAmount)",
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

		collection, err := dao.FindCollectionByNameOrId("mbgvq8b9yf5i2d1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
