package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func Migrate(sourceURL, dbName string) error {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		sourceURL,
		dbName,
		driver,
	)
	if err := m.Up(); err != nil {
		log.Fatal(err)
		return err
	}

	logrus.Infof("Migrate Done")
	return nil
}
