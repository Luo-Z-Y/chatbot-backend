package main

import (
	"backend/internal/configs"
	"backend/internal/database"
	"flag"
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

const migrationFilePath = "./migrations"

func main() {
	var dirFlag = flag.Bool("down", false, "undo migration. omit for normal migration (upwards)")
	var countFlag = flag.Int("count", 0, "number of database migrations to run. omit (or 0) for no limit")
	var testDbFlag = flag.Bool("testdb", false, "run test database migrations. omit to run on development db")
	flag.Parse()

	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if *testDbFlag {
		database.SetupDb(cfg.GetTestDatabaseConfig())
	} else {
		database.SetupDb(cfg.GetDatabaseConfig())
	}

	db, err := database.GetDb().DB()
	if err != nil {
		log.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: migrationFilePath,
	}

	direction := migrate.Up
	if *dirFlag {
		direction = migrate.Down
	}

	count, err := migrate.ExecMax(db, "postgres", migrations, direction, *countFlag)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("applied %d migrations\n", count)
}
