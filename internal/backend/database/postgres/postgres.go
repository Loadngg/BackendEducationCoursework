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
		"sslmode=%s host=%s port=%v user=%s password=%s dbname=%s",
		cfg.SSL,
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)

	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unnable to connect to database: %w", err)
	}

	dbLogger := log.New(os.Stderr, "SQL: ", log.Flags())
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(dbLogger.Printf))
	return db, nil
}
