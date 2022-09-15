package dbtest

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	dbMigrate "github.com/Thanh17b4/practice/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

const (
	expirationInSeconds = 120
	dbPoolSize          = 40
	dbPoolTimeout       = 20 * time.Second
	dbPoolMaxWait       = 1200 * time.Second
)

// TestUtil represents data to help running integration tests.
type TestUtil struct {
	DB         *sql.DB
	DBURL      string
	PostgresDB *sqlx.DB

	Log *logrus.Entry
}

// New creates a pointer to the instance of test-util
func New() *TestUtil {
	log := logrus.New()
	return &TestUtil{
		Log: logrus.NewEntry(log),
	}
}

func (util *TestUtil) bootstrapDB() error {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = dbPoolMaxWait
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
		return err
	}

	// run docker to create a database
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "5.7",
		Env:        []string{"MYSQL_ROOT_PASSWORD=221292", "MYSQL_DATABASE=token_service"},
	})
	if err != nil {
		log.Fatalf("could not init database: %s", err)
		return err
	}

	if err := resource.Expire(expirationInSeconds); err != nil {
		return err
	}

	if err := util.connectDB(pool, resource); err != nil {
		return err
	}

	return nil
}

func (util *TestUtil) connectDB(pool *dockertest.Pool, containerResource *dockertest.Resource) error {
	address := containerResource.GetPort("3306/tcp")
	dbURL := fmt.Sprintf("root:221292@(localhost:%s)/token_service?parseTime=true", address)
	util.DBURL = dbURL
	pool.MaxWait = dbPoolMaxWait
	return pool.Retry(func() error {
		var err error
		util.DB, err = sql.Open("mysql", fmt.Sprintf("root:221292@tcp(localhost:%s)/token_service?multiStatements=true", containerResource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}

		if err := util.DB.Ping(); err != nil {
			return err
		}

		return nil
	})
}

// InitDB initializes db.
func (util *TestUtil) InitDB() error {
	if err := util.bootstrapDB(); err != nil {
		return err
	}

	return nil
}

// SetupDB do creating schemas and populate data to database
func (util *TestUtil) SetupDB() error {
	if err := dbMigrate.Migrate("file://../../db/migrations", util.DBURL); err != nil {
		return err
	}

	return nil
}

func (util *TestUtil) CleanAndClose() {

}
