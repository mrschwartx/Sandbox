package db

import (
	"errors"
	"fmt"

	"example.com/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // need this here for running migrate on testing.
	_ "github.com/golang-migrate/migrate/v4/source/file"       // need this here for running migrate on testing.
	"github.com/jmoiron/sqlx"
)

func NewSQLx(c config.SQL) *sqlx.DB {
	db, err := sqlx.Open("postgres", c.GetDataSource())
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(c.MaxOpenConn)
	db.SetMaxIdleConns(c.MaxIdleConn)

	return db
}

func PostgresMigrate(dir string, databaseUrl string, up bool) error {
	m, err := migrate.New(fmt.Sprintf("file://%s", dir), databaseUrl)
	if err != nil {
		return err
	}
	if up {
		err = m.Up()
	} else {
		err = m.Down()
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
