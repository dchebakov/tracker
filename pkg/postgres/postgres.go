package postgres

import (
	"fmt"

	"github.com/dchebakov/tracker/config"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const pgDriver = "pgx"

func NewDB(c *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Name,
		c.Database.Pass,
	)

	db, err := sqlx.Connect(pgDriver, dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.Database.MaxConnectionPool)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
