package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		// Skip if the bot user already exists.
		if _, err := app.FindRecordById("_pb_users_auth_", "pbbotbjalla0000"); err == nil {
			return nil
		}

		// fetch `users` collection
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		record := core.NewRecord(collection)

		record.Id = "pbbotbjalla0000"
		record.Set("email", "bjalla@bjalla.eu")
		record.Set("name", "Bjalla")
		record.Set("verified", true)
		record.SetRandomPassword()

		return app.Save(record)
	}, func(app core.App) error {
		// delete the bot user in `users` collection
		record, err := app.FindRecordById("_pb_users_auth_", "pbbotbjalla0000")
		if err != nil {
			return nil
		}

		return app.Delete(record)
	})
}
