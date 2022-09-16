package repo_test

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
)

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=221292",
			"MYSQL_DATABASE=token_service",
		},
		Platform: "linux/x86_64",
	}

	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	time.Sleep(time.Second * 60)
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	var db *sql.DB
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:221292@(127.0.0.1:%s)/token_service?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

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

func SetupDB() (*dockertest.Resource, *dockertest.Pool, *sql.DB) {
	// Step 1: create database mySQL by running a mysql container
	resource, pool := createContainer()
	// Step 2: connect to a database
	db := connectDB(resource, pool)
	// step 3: Run migrations to create tables
	if err := migrateDB(db); err != nil {
		logrus.Fatalf("could not migrate db")
	}
	return resource, pool, db
}
