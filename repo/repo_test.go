package repo_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func migrateDB(db *sql.DB) error {
	logrus.Infof("Start migrate")

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	migrateOps, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"token_service",
		driver,
	)

	logrus.Infof("End init migrate")
	if err != nil {
		fmt.Println("could not migrateDB: ", err.Error())
		return err
	}

	return migrateOps.Steps(2)
}

func cleanUpDB(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS `schema_migrations`")
	if err != nil {
		logrus.Println("could not truncate schema_migrations table: ", err.Error())
	}
	_, err = db.Exec("DROP TABLE IF EXISTS `users`")
	if err != nil {
		logrus.Println("could not truncate users table: ", err.Error())
	}
}

func SetupDB() (*sql.DB, error) {
	// run this command to create mysql container: `docker run -d -p 3306:3306 --name mysql --platform linux/x86_64 --env MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=token_service mysql:5.7`
	dbURL := "root:12345@(127.0.0.1:3306)/token_service?parseTime=true"
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	// clean DB before starting new test case
	cleanUpDB(db)

	err = migrateDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
