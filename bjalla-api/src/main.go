package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	// Check if the program is in development mode.
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	_ = godotenv.Load(".env")
	app := pocketbase.New()

	// Register the migration command.
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
