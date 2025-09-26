package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/su-de-sh/nestly/internal/config"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresConnection(cfg *config.Config) (*PostgresDB, error) {
	connectionString := cfg.GetDatabaseConnectionString()
	fmt.Println("Conection string", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

func (p *PostgresDB) HealthCheck() error {
	return p.DB.Ping()
}
