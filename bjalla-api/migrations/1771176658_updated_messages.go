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
			"createRule": "author = @request.auth.id &&\n// only if @request.auth is member in the channels' guild\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= channel.guild",
			"deleteRule": "@request.auth.id = author ||\n// user is admin in guild of this message\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members:auth.guild ?= channel.guild &&\n  @collection.guild_members:auth.roles.is_admin ?= true &&\n  @collection.guild_members:auth.roles.guild ?= channel.guild)",
			"listRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= channel.guild",
			"viewRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= channel.guild"
		}`), &collection); err != nil {
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
			"createRule": "author = @request.auth.id &&\n// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild",
			"deleteRule": "@request.auth.id = author ||\n// user is admin in guild of this message\n(@collection.guild_channels.channel ?= channel &&\n  @collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members:auth.guild ?= @collection.guild_channels.guild &&\n  @collection.guild_members:auth.roles.is_admin ?= true &&\n  @collection.guild_members:auth.roles.guild ?= @collection.guild_channels.guild)",
			"listRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild",
			"viewRule": "// only if @request.auth is member in the channels' guild\n@collection.guild_channels.channel ?= channel &&\n@collection.guild_members:auth.user ?= @request.auth.id &&\n@collection.guild_members:auth.guild ?= @collection.guild_channels.guild"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
