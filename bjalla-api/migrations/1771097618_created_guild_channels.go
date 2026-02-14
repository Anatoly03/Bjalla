package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"cascadeDelete": false,
					"collectionId": "pbc_1835829336",
					"hidden": false,
					"id": "relation1967160747",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "guild",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
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
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_3028263738",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_yw1af0R0U1` + "`" + ` ON ` + "`" + `guild_channels` + "`" + ` (\n  ` + "`" + `guild` + "`" + `,\n  ` + "`" + `channel` + "`" + `\n)"
			],
			"listRule": "@collection.guild_members.guild ?= guild && @collection.guild_members.user = @request.auth.id",
			"name": "guild_channels",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": "@collection.guild_members.guild ?= guild && @collection.guild_members.user = @request.auth.id"
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3028263738")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
