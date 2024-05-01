package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"log"
)

type sqliteDB struct {
	client *bun.DB
	conn   *bun.Conn
	ctx    context.Context
	driver database.Driver
}

func NewTestDatabase() Persistence {
	log.Println("Setting Test DB connection")
	var database sqliteDB
	database.client = getSQLiteClient()
	database.ctx = context.Background()
	err := database.Connect()
	if err != nil {
		log.Fatalf("Test DB connection with error: %v\n", err)
	}
	driver, err := sqlite.WithInstance(database.client.DB, &sqlite.Config{})
	if err != nil {
		log.Fatalf("Test DB connection with error: %v\n", err)
	}
	database.driver = driver
	return database
}

func getSQLiteClient() *bun.DB {
	sqlite, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	// Create a Bun db on top of it.
	return bun.NewDB(sqlite, sqlitedialect.New())
}

func (d sqliteDB) Connect() error {
	conn, err := d.client.Conn(d.Context())

	if err != nil {
		return err
	}

	d.conn = &conn

	return nil
}

func (d sqliteDB) IsAlive() error {
	if d.conn == nil {
		return errors.New("db connection has not been set up")
	}
	return nil
}

func (d sqliteDB) PingContext() error {
	return d.client.PingContext(d.ctx)
}

func (d sqliteDB) Select() *bun.SelectQuery {
	return d.client.NewSelect()
}

func (d sqliteDB) Insert() *bun.InsertQuery {
	return d.client.NewInsert()
}

func (d sqliteDB) Update() *bun.UpdateQuery {
	return d.client.NewUpdate()
}

func (d sqliteDB) Context() context.Context {
	return d.ctx
}

func (d sqliteDB) Raw(q string, args ...interface{}) *bun.RawQuery {
	return d.client.NewRaw(q, args)
}

func (d sqliteDB) GetDriver() database.Driver {
	return d.driver
}
