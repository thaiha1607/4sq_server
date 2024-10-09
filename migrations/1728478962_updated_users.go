package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\"")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\"")

		collection.CreateRule = types.Pointer("@request.auth.role = \"manager\" || @request.data.role = \"customer\"")

		collection.UpdateRule = types.Pointer("(id = @request.auth.id && @request.data.role = role) || @request.auth.role = \"manager\"")

		collection.DeleteRule = types.Pointer("id = @request.auth.id || @request.auth.role = \"manager\"")

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_wb5EYuS`+"`"+` ON `+"`"+`users`+"`"+` (`+"`"+`role`+"`"+`) WHERE `+"`"+`role`+"`"+` = 'staff'"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("users_avatar")

		// add
		new_avatarUrl := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), new_avatarUrl); err != nil {
			return err
		}
		collection.Schema.AddField(new_avatarUrl)

		// add
		new_phone := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), new_phone); err != nil {
			return err
		}
		collection.Schema.AddField(new_phone)

		// add
		new_role := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), new_role); err != nil {
			return err
		}
		collection.Schema.AddField(new_role)

		// update
		edit_name := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_name); err != nil {
			return err
		}
		collection.Schema.AddField(edit_name)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id")

		collection.ViewRule = types.Pointer("id = @request.auth.id")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("id = @request.auth.id")

		collection.DeleteRule = types.Pointer("id = @request.auth.id")

		if err := json.Unmarshal([]byte(`[]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		del_avatar := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_avatar",
			"name": "avatar",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"image/jpeg",
					"image/png",
					"image/svg+xml",
					"image/gif",
					"image/webp"
				],
				"thumbs": null,
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), del_avatar); err != nil {
			return err
		}
		collection.Schema.AddField(del_avatar)

		// remove
		collection.Schema.RemoveField("qv1rgzqs")

		// remove
		collection.Schema.RemoveField("l5r0tgky")

		// remove
		collection.Schema.RemoveField("upj9eszj")

		// update
		edit_name := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_name",
			"name": "name",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_name); err != nil {
			return err
		}
		collection.Schema.AddField(edit_name)

		return dao.SaveCollection(collection)
	})
}
