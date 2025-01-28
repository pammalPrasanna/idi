package infra

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
)

const (
	defaultMaxOpenConns    = 25
	defaultMaxIdleConns    = 25
	defaultConnMaxIdleTime = 5 * time.Minute
	defaultConnMaxLifetime = 2 * time.Hour
)

var (
	dsn             = os.Getenv("DB_DSN")
	maxOpenConns    = os.Getenv("DB_MAX_OPEN_CONNS")
	maxIdleConns    = os.Getenv("DB_MAX_IDLE_CONNS")
	connMaxIdleTime = os.Getenv("DB_CONN_MAX_IDLE_TIME")
	connMaxLifetime = os.Getenv("DB_CONN_MAX_LIFETIME")

	instance *sql.DB
)

func Postgres() (*sql.DB, error) {
	if instance != nil {
		return instance, nil
	}
	conn, err := sql.Open("psql", dsn)
	if err != nil {
		return nil, err
	}

	if maxOpenConns != "" {
		if i, err := strconv.Atoi(maxOpenConns); err == nil {
			conn.SetMaxOpenConns(i)
		} else {
			return nil, err
		}
	} else {
		conn.SetMaxOpenConns(defaultMaxOpenConns)
	}

	if maxIdleConns != "" {
		if i, err := strconv.Atoi(maxIdleConns); err == nil {
			conn.SetMaxIdleConns(i)
		} else {
			return nil, err
		}
	} else {
		conn.SetMaxIdleConns(defaultMaxIdleConns)
	}

	if connMaxIdleTime != "" {
		if i, err := strconv.Atoi(connMaxIdleTime); err == nil {
			conn.SetConnMaxIdleTime(time.Duration(i * int(time.Minute)))
		} else {
			return nil, err
		}
	} else {
		conn.SetConnMaxIdleTime(defaultMaxIdleConns)
	}

	if connMaxLifetime != "" {
		if i, err := strconv.Atoi(connMaxLifetime); err == nil {
			conn.SetConnMaxLifetime(time.Duration(i * int(time.Minute)))
		} else {
			return nil, err
		}
	} else {
		conn.SetConnMaxLifetime(defaultConnMaxLifetime)
	}

	instance = conn
	return instance, nil
}
