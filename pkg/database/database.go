package database

import (
	"database/sql"
	"fmt"
	"log"
	"monolithic-go-app/pkg/config"

	_ "github.com/lib/pq"
)

func InitDB(databaseCfg config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		databaseCfg.Host,
		databaseCfg.Port,
		databaseCfg.User,
		databaseCfg.Password,
		databaseCfg.Name,
		databaseCfg.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
	return db, nil
}
