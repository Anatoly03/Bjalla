package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2620268817")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.id=user.id ||\n// compare if auth user and row user have mutual guild\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members.user ?= id &&\n  @collection.guild_members:auth.guild ?= @collection.guild_members.guild)",
			"viewRule": "@request.auth.id=user.id ||\n// compare if auth user and row user have mutual guild\n(@collection.guild_members:auth.user ?= @request.auth.id &&\n  @collection.guild_members.user ?= id &&\n  @collection.guild_members:auth.guild ?= @collection.guild_members.guild)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2620268817")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.id=user.id",
			"viewRule": "@request.auth.id=user.id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
