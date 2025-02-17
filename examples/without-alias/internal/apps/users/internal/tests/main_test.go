package users_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"without-alias/internal/lib"

	"github.com/joho/godotenv"
)

var (
	dbConn            *sql.DB
	rootApp           lib.IApp
	INTEGRATION_TESTS bool
)

func setupTestDB(m *testing.M) (int, error) {
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		INTEGRATION_TESTS = false
		fmt.Println("skipping database setup")
		return m.Run(), nil
	} else {
		INTEGRATION_TESTS = true
	}

	var err error

	if err := os.Remove(dbfile); err != nil {
		log.Printf("unable to delete testdb with error: %s", err)
	}
	conn, err := Sqlite3Test()
	if err != nil {
		return -1, err
	}

	cmd := exec.Command("goose",
		"-s", "-dir=../../../../../migrations", "sqlite3", "testdb.db", "up",
	)

	if err := cmd.Run(); err != nil {
		log.Printf("unable to run migrations with error: %s", err)
	}

	dbConn = conn

	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("unable to close DB with error: '%s'", err)
		}
		if err := os.Remove(dbfile); err != nil {
			log.Printf("unable to delete testdb with error: %s", err)
		}
	}()

	return m.Run(), nil
}

func TestMain(m *testing.M) {
	godotenv.Load("../../../../../.env")
	ra, err := lib.Idi()
	if err != nil {
		log.Printf("unable to create root app: %s", err)
		os.Exit(-1)
	}
	rootApp = ra
	code, err := setupTestDB(m)
	if err != nil {
		log.Printf("unable to create test db connection: %s", err)
	}
	os.Exit(code)
}
