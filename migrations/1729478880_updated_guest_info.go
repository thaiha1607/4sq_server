package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("68wsl8pf0ab44u1")
		if err != nil {
			return err
		}

		// update
		edit_phone := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
				"pattern": "^\\+(?:[0-9] ?){6,14}[0-9]$"
			}
		}`), edit_phone); err != nil {
			return err
		}
		collection.Schema.AddField(edit_phone)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("68wsl8pf0ab44u1")
		if err != nil {
			return err
		}

		// update
		edit_phone := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_phone); err != nil {
			return err
		}
		collection.Schema.AddField(edit_phone)

		return dao.SaveCollection(collection)
	})
}
