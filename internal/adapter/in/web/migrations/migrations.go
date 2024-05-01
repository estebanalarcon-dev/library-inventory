package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"log"
)

func Migrate(driver database.Driver) {
	m, err := migrate.NewWithDatabaseInstance("file://migrations",
		"sqlite3", driver)
	if err != nil {
		fmt.Println("migrate:", err)
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println("Up: ", err)
		log.Fatal(err)
	}
}
