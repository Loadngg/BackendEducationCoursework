package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/reform.v1/dialects/postgresql"

	"gopkg.in/reform.v1"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSL      string
}

func New(cfg Config) (*reform.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSL,
	)

	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unnable to connect to database: %w", err)
	}

	dbLogger := log.New(os.Stderr, "SQL: ", log.Flags())
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(dbLogger.Printf))
	return db, nil
}
