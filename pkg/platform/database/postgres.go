package db

import (
	"fmt"

	"github.com/cp-Coder/khelo/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*gorm.DB, error) {
	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	// Define database connection for PostgreSQL.
	db, err := gorm.Open(postgres.Open(postgresConnURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	return db, nil
}
