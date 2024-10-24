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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2024-09-04 02:11:30.920Z",
				"updated": "2024-09-29 12:08:25.897Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "qv1rgzqs",
						"name": "avatarUrl",
						"type": "url",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "l5r0tgky",
						"name": "phone",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^\\+(?:[0-9] ?){6,14}[0-9]$"
						}
					},
					{
						"system": false,
						"id": "upj9eszj",
						"name": "role",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"customer",
								"staff",
								"manager"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_wb5EYuS` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `role` + "`" + `) WHERE ` + "`" + `role` + "`" + ` = 'staff'"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.role = \"manager\" || @request.data.role = \"customer\"",
				"updateRule": "(id = @request.auth.id && @request.data.role = role) || @request.auth.role = \"manager\"",
				"deleteRule": "id = @request.auth.id || @request.auth.role = \"manager\"",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "7uhp3ed4afuq9uc",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-26 11:55:35.931Z",
				"name": "addresses",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "rbv1jzdm",
						"name": "line1",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ncmw5m13",
						"name": "line2",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "e9jojbr1",
						"name": "city",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "anqbruir",
						"name": "state",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "sul3iaab",
						"name": "country",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pc4cgzj6",
						"name": "zipOrPostcode",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": null,
				"deleteRule": "@request.auth.id != \"\"",
				"options": {}
			},
			{
				"id": "3o2zbir67gxkcni",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-26 11:55:45.866Z",
				"name": "colours",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "xflmf0yq",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "2kpyfojn",
						"name": "hexCode",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "yy58rkdg4k4kjs5",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-29 14:09:16.199Z",
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
			},
			{
				"id": "eblguzfwjtr5n6s",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-22 13:14:29.502Z",
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
			},
			{
				"id": "bd0d5gtacvvnnys",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-26 12:02:37.544Z",
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
			},
			{
				"id": "gsc7xs8i2q990pg",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-26 11:58:59.840Z",
				"name": "internal_order_items",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "k33papdq",
						"name": "qty",
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
						"id": "6yjchw0j",
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
						"id": "xhyfdg6v",
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
						"id": "ypmd9q9s",
						"name": "internalOrderId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "g0y7nfa8ommcv1h",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ad5rzbhz",
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
					"CREATE UNIQUE INDEX ` + "`" + `idx_3TMYpsm` + "`" + ` ON ` + "`" + `internal_order_items` + "`" + ` (\n  ` + "`" + `internalOrderId` + "`" + `,\n  ` + "`" + `orderItemId` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.internalOrderId.rootOrderId = @request.data.orderItemId.orderId",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.internalOrderId = internalOrderId && @request.data.orderItemId = orderItemId",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "g0y7nfa8ommcv1h",
				"created": "2024-09-04 02:14:31.777Z",
				"updated": "2024-09-28 02:10:38.709Z",
				"name": "internal_orders",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "yqlybqb6",
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
						"id": "ogr0e7bq",
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
						"id": "ukra4tgz",
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
						"id": "uudmfia5",
						"name": "rootOrderId",
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
						"id": "wi7c7mns",
						"name": "shipmentId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "0fvjhfcs9ig8nd8",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "gvj94it2",
						"name": "srcWorkingUnitId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "cr8iizw10m7kh1d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "k4lbq406",
						"name": "dstWorkingUnitId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "cr8iizw10m7kh1d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.rootOrderId = rootOrderId && @request.data.type = type",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "ldiepkv7gnn4nrr",
				"created": "2024-09-04 02:14:31.778Z",
				"updated": "2024-09-26 11:59:30.127Z",
				"name": "invoice_status_codes",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "1buavidk",
						"name": "statusCode",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "axubqxgs",
						"name": "description",
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
					"CREATE UNIQUE INDEX ` + "`" + `idx_hl4uZdG` + "`" + ` ON ` + "`" + `invoice_status_codes` + "`" + ` (` + "`" + `statusCode` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "mbgvq8b9yf5i2d1",
				"created": "2024-09-04 02:14:31.778Z",
				"updated": "2024-09-29 10:00:17.276Z",
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
			},
			{
				"id": "3ryn3wvszi2ny9q",
				"created": "2024-09-04 02:14:31.778Z",
				"updated": "2024-09-22 13:14:29.506Z",
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
			},
			{
				"id": "nqclqjsjbs7e523",
				"created": "2024-09-04 02:14:31.778Z",
				"updated": "2024-09-29 07:47:35.722Z",
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
			},
			{
				"id": "93h7enu78ajtk7w",
				"created": "2024-09-04 02:14:31.778Z",
				"updated": "2024-09-22 13:14:29.506Z",
				"name": "order_status_codes",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zgwsoimk",
						"name": "statusCode",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3jemdow9",
						"name": "description",
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
					"CREATE UNIQUE INDEX ` + "`" + `idx_YhmYEHX` + "`" + ` ON ` + "`" + `order_status_codes` + "`" + ` (` + "`" + `statusCode` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "2stc36eglv3voqe",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-27 00:19:06.077Z",
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
			},
			{
				"id": "nhuyrnedwhw2jm8",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-22 13:14:29.506Z",
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
			},
			{
				"id": "upkshg4h89ndt95",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-26 12:05:21.610Z",
				"name": "product_categories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vtp20sii",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "6apt3tel",
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
						"id": "iehyx3oh",
						"name": "colourId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "3o2zbir67gxkcni",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_Xa76LkK` + "`" + ` ON ` + "`" + `product_categories` + "`" + ` (\n  ` + "`" + `productId` + "`" + `,\n  ` + "`" + `colourId` + "`" + `\n)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "sznmzh643y5pv5s",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-22 13:14:29.510Z",
				"name": "product_images",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cugfpyja",
						"name": "imageUrl",
						"type": "url",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "0xxipmnd",
						"name": "altText",
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
						"id": "lz2i7fzs",
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
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_NZuOJuB` + "`" + ` ON ` + "`" + `product_images` + "`" + ` (` + "`" + `productId` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "gomuxh946wr3qbj",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-29 01:37:01.890Z",
				"name": "product_quantities",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "24k4ot5g",
						"name": "priority",
						"type": "number",
						"required": true,
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
						"id": "drgo0oo1",
						"name": "qty",
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
						"id": "tisyu2hi",
						"name": "categoryId",
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
					},
					{
						"system": false,
						"id": "l4jcckup",
						"name": "workingUnitId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "cr8iizw10m7kh1d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_B2bqLvs` + "`" + ` ON ` + "`" + `product_quantities` + "`" + ` (\n  ` + "`" + `categoryId` + "`" + `,\n  ` + "`" + `workingUnitId` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "gpvjyk1s8lipcqo",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-22 13:14:29.510Z",
				"name": "products",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "x8n2fv7t",
						"name": "name",
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
						"id": "jceepjq0",
						"name": "description",
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
						"id": "prnqsamb",
						"name": "expectedPrice",
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
						"id": "ii0yxz7n",
						"name": "provider",
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
						"id": "pnzgatb7",
						"name": "tagIds",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "fpzmqrtemgkwcs6",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 32,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "6utk07e3r3pigtn",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-10-03 01:54:12.144Z",
				"name": "shipment_assignments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gcdvlyq8",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"pending",
								"assigned",
								"in_progress",
								"completed",
								"cancelled",
								"failed",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "mqtaqezd",
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
						"id": "mvjqphgb",
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
						"id": "15cwxflp",
						"name": "staffId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "v6f615blwqyoe4d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.role = \"manager\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.role = \"manager\"",
				"options": {}
			},
			{
				"id": "st23val23amwoi5",
				"created": "2024-09-04 02:14:31.779Z",
				"updated": "2024-09-27 00:28:40.845Z",
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
			},
			{
				"id": "inophh2p3nyvq5c",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-22 13:14:29.511Z",
				"name": "shipment_status_codes",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "jqcy8b80",
						"name": "statusCode",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fgxdixiy",
						"name": "description",
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
					"CREATE UNIQUE INDEX ` + "`" + `idx_tbnOBcE` + "`" + ` ON ` + "`" + `shipment_status_codes` + "`" + ` (` + "`" + `statusCode` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "0fvjhfcs9ig8nd8",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-27 00:29:05.146Z",
				"name": "shipments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ikmp1nmh",
						"name": "type",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"outbound",
								"inbound",
								"transfer",
								"return",
								"exchange",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "rigsg6vw",
						"name": "shipmentDate",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "v2vpuxyc",
						"name": "deliveryDate",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "b1slbcyj",
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
						"id": "fpekt4ye",
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
						"id": "lyollt8g",
						"name": "invoiceId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "mbgvq8b9yf5i2d1",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "cjylj0t6",
						"name": "statusCodeId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "inophh2p3nyvq5c",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
				"viewRule": "orderId.creatorId = @request.auth.id || (@request.auth.id != \"\" && @request.auth.role != \"customer\")",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.orderId = @request.data.invoiceId.orderId",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\" && @request.data.orderId = orderId && @request.data.invoiceId.orderId = orderId",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "v6f615blwqyoe4d",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-29 03:14:02.488Z",
				"name": "staff_info",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "lpq224tk",
						"name": "statusCode",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"active",
								"inactive",
								"suspended",
								"terminated",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "thsgvtek",
						"name": "role",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"salesperson",
								"warehouse",
								"delivery",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "lbgtar1m",
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
					},
					{
						"system": false,
						"id": "kbmzxgda",
						"name": "workingUnitId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "cr8iizw10m7kh1d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_kikHSGl` + "`" + ` ON ` + "`" + `staff_info` + "`" + ` (` + "`" + `userId` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_r696cdI` + "`" + ` ON ` + "`" + `staff_info` + "`" + ` (` + "`" + `role` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.role = \"manager\"",
				"updateRule": "@request.auth.role = \"manager\"",
				"deleteRule": "@request.auth.role = \"manager\"",
				"options": {}
			},
			{
				"id": "fpzmqrtemgkwcs6",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-26 12:08:10.081Z",
				"name": "tags",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vfnvroid",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 64,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_ZcKxUtS` + "`" + ` ON ` + "`" + `tags` + "`" + ` (` + "`" + `name` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.role = \"manager\"",
				"updateRule": null,
				"deleteRule": "@request.auth.role = \"manager\"",
				"options": {}
			},
			{
				"id": "4mc1mg16zw31tum",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-26 12:08:34.310Z",
				"name": "user_addresses",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "hsz5thp5",
						"name": "type",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"home",
								"work",
								"billing",
								"shipping",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "fqkslwn0",
						"name": "friendlyName",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 64,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3yjyqbnv",
						"name": "isDefault",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "vlr8r3cp",
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
					},
					{
						"system": false,
						"id": "zkamkzp9",
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
					"CREATE UNIQUE INDEX ` + "`" + `idx_iLfx8Tu` + "`" + ` ON ` + "`" + `user_addresses` + "`" + ` (\n  ` + "`" + `userId` + "`" + `,\n  ` + "`" + `addressId` + "`" + `\n)"
				],
				"listRule": "userId = @request.auth.id",
				"viewRule": "userId = @request.auth.id",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "userId = @request.auth.id",
				"deleteRule": "userId = @request.auth.id",
				"options": {}
			},
			{
				"id": "dmlx8wh39pu41u8",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-10-03 01:53:28.138Z",
				"name": "warehouse_assignments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "3v5lcqbt",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"pending",
								"assigned",
								"in_progress",
								"completed",
								"cancelled",
								"failed",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "gvim9fk7",
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
						"id": "uk9hplvt",
						"name": "staffId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "v6f615blwqyoe4d",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "rji4r1bl",
						"name": "internalOrderId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "g0y7nfa8ommcv1h",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.role = \"manager\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.role = \"manager\"",
				"options": {}
			},
			{
				"id": "cr8iizw10m7kh1d",
				"created": "2024-09-04 02:14:31.780Z",
				"updated": "2024-09-22 13:14:29.512Z",
				"name": "working_units",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "o9gekwse",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xn47i6tj",
						"name": "type",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"warehouse",
								"office",
								"delivery",
								"other"
							]
						}
					},
					{
						"system": false,
						"id": "ry1rm3ou",
						"name": "imageUrl",
						"type": "url",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "jn427lzh",
						"name": "addressId",
						"type": "relation",
						"required": false,
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
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.role = \"manager\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.role = \"manager\"",
				"options": {}
			},
			{
				"id": "1hln2do3ca9mwaz",
				"created": "2024-09-12 09:25:29.264Z",
				"updated": "2024-09-29 01:44:36.093Z",
				"name": "product_quantity_summary",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "t7aymr5p",
						"name": "totalQty",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT categoryId as id, SUM(qty) AS totalQty\nFROM product_quantities\nGROUP BY categoryId;"
				}
			},
			{
				"id": "kgkfehwran8zgs8",
				"created": "2024-09-19 11:20:14.375Z",
				"updated": "2024-09-22 13:14:29.512Z",
				"name": "transaction_history",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "xy2z6b67",
						"name": "entityType",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ubiiz0ft",
						"name": "entityId",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "cprxxoqe",
						"name": "statusCodeId",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "hycktg06",
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
					"CREATE INDEX ` + "`" + `idx_kcKC4Vm` + "`" + ` ON ` + "`" + `transaction_history` + "`" + ` (\n  ` + "`" + `entityType` + "`" + `,\n  ` + "`" + `entityId` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "jzxryofwy4tmpha",
				"created": "2024-09-23 09:07:34.130Z",
				"updated": "2024-09-29 13:27:10.688Z",
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
			},
			{
				"id": "68wsl8pf0ab44u1",
				"created": "2024-09-27 00:17:34.115Z",
				"updated": "2024-09-27 00:17:34.115Z",
				"name": "guest_info",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "yznhx5l7",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 256,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "kujsv6yr",
						"name": "email",
						"type": "email",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "f0ozkueb",
						"name": "phone",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^\\+[1-9]\\d{1,14}$"
						}
					},
					{
						"system": false,
						"id": "mrprgcue",
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
				"indexes": [],
				"listRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"viewRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role != \"customer\"",
				"options": {}
			},
			{
				"id": "bg69cxt80nkfws1",
				"created": "2024-09-29 09:14:09.393Z",
				"updated": "2024-09-29 09:14:09.393Z",
				"name": "product_quantity_history",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "jfgew7gh",
						"name": "categoryId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "upkshg4h89ndt95",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "jwg5tfn2",
						"name": "amountOfChange",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": true
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"manager\"",
				"viewRule": "@request.auth.role = \"manager\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "90dgpouuugdj3g8",
				"created": "2024-09-29 09:20:53.989Z",
				"updated": "2024-09-29 09:30:07.423Z",
				"name": "daily_income",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "obi8scyl",
						"name": "amountOfChange",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"manager\"",
				"viewRule": "@request.auth.role = \"manager\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
