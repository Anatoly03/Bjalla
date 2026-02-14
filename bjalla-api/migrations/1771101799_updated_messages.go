package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2605467279")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "author = @request.auth.id &&\n// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild",
			"deleteRule": "@request.auth.id = author",
			"listRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild",
			"updateRule": "@request.auth.id = author",
			"viewRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_3009067695",
			"hidden": false,
			"id": "relation2734263879",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "channel",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2605467279")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": null,
			"deleteRule": null,
			"listRule": null,
			"updateRule": null,
			"viewRule": ""
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2734263879")

		return app.Save(collection)
	})
}
