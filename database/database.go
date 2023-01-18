package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Printf("Applied %d migrations!\n", n)
}
