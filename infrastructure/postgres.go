package infrastructure

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"libraryInventory/internal/config"
	"log"
)

// Persistence is an adapter for the database ORM
type Persistence interface {
	PingContext() error
	IsAlive() error
	Connect() error
	Select() *bun.SelectQuery
	Insert() *bun.InsertQuery
	Update() *bun.UpdateQuery
	Context() context.Context
	Raw(q string, args ...interface{}) *bun.RawQuery
	GetDriver() database.Driver
}

type postgresDB struct {
	client *bun.DB
	conn   *bun.Conn
	ctx    context.Context
}

// NewBunMysqlClient returns a new bun connection and db client (connection is used for pinging)
func NewPersistence(conf *config.Postgres) Persistence {
	log.Println("Setting DB connection")
	var database postgresDB
	database.client = getPostgresClient(conf)
	database.ctx = context.Background()
	err := database.Connect()
	if err != nil {
		log.Fatalf("DB connection with error: %v\n", err)
	}
	log.Println("DB connection was set successfully")
	return &database
}

func getPostgresClient(config *config.Postgres) *bun.DB {
	sqlDB := sql.OpenDB(createConnector(config))
	return bun.NewDB(sqlDB, pgdialect.New())

}

func createConnector(config *config.Postgres) *pgdriver.Connector {
	return pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(config.Host+":"+config.Host),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(config.User),
		pgdriver.WithPassword(config.Password),
		pgdriver.WithDatabase(config.Name))
}

func (d *postgresDB) Connect() error {
	conn, err := d.client.Conn(d.Context())

	if err != nil {
		return err
	}

	d.conn = &conn

	return nil
}

func (d postgresDB) IsAlive() error {
	if d.conn == nil {
		return errors.New("db connection has not been set up")
	}
	return nil
}

func (d postgresDB) PingContext() error {
	return d.client.PingContext(d.ctx)
}

func (d postgresDB) Select() *bun.SelectQuery {
	return d.client.NewSelect()
}

func (d postgresDB) Insert() *bun.InsertQuery {
	return d.client.NewInsert()
}

func (d postgresDB) Update() *bun.UpdateQuery {
	return d.client.NewUpdate()
}

func (d postgresDB) Context() context.Context {
	return d.ctx
}

func (d postgresDB) Raw(q string, args ...interface{}) *bun.RawQuery {
	return d.client.NewRaw(q, args)
}

func (d postgresDB) GetDriver() database.Driver {
	return nil
}
