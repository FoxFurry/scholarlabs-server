package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrator_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func NewDB(dbHost, dbPort, dbUser, dbPass, dbName string) (*sqlx.DB, error) {
	mysqlcfg := mysql.Config{
		User:            dbUser,
		Passwd:          dbPass,
		Addr:            fmt.Sprintf("%s:%s", dbHost, dbPort),
		Net:             "tcp",
		DBName:          dbName,
		MultiStatements: true,
		ParseTime:       true,
	}

	db, err := sqlx.Open("mysql", mysqlcfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sqlx.DB, path_to_migrations string) error {
	migrator_driver, err := migrator_mysql.WithInstance(db.DB, &migrator_mysql.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://"+path_to_migrations,
		"mysql",
		migrator_driver)
	if err != nil {
		return err
	}

	return migrator.Up()
}
