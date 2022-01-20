package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dchebakov/tracker/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
)

const databaseMigrationPath = "file://migrations/"
const databaseDriverName = "pgx"

func main() {
	cfg := config.New()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Name,
		cfg.Database.Pass,
	)
	db, err := sql.Open(databaseDriverName, dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error instantiating database: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(databaseMigrationPath, databaseDriverName, driver)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Migrating to the lastest revision...")
	err = m.Up()
	if err != nil {
		log.Fatalf("Failed to migrated: %s", err)
	}

	log.Println("Migrated succsefully")
}
