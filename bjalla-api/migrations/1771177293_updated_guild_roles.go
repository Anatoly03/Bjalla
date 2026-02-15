package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2000738042")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "// user is already admin in guild of this role\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members:auth.guild ?= guild &&\n  @collection.guild_members:auth.roles.is_admin ?= true &&\n  @collection.guild_members:auth.roles.guild ?= guild)",
			"deleteRule": "// user is already admin in guild of this role\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members:auth.guild ?= guild &&\n  @collection.guild_members:auth.roles.is_admin ?= true &&\n  @collection.guild_members:auth.roles.guild ?= guild)",
			"updateRule": "// user is already admin in guild of this role\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members:auth.guild ?= guild &&\n  @collection.guild_members:auth.roles.is_admin ?= true &&\n  @collection.guild_members:auth.roles.guild ?= guild)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2000738042")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "// TODO: only admins and permitted users\n1=1",
			"deleteRule": "// TODO: only admins and permitted users\n1=1",
			"updateRule": "// TODO: only admins and permitted users\n1=1"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
