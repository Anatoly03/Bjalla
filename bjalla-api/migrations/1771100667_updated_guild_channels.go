package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3028263738")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@collection.guild_members:auth.user ?= @request.auth.id && @collection.guild_members:auth.guild ?= guild",
			"viewRule": "@collection.guild_members:auth.user ?= @request.auth.id && @collection.guild_members:auth.guild ?= guild"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3028263738")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@collection.guild_members.guild ?= guild && @collection.guild_members.user = @request.auth.id",
			"viewRule": "@collection.guild_members.guild ?= guild && @collection.guild_members.user = @request.auth.id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
