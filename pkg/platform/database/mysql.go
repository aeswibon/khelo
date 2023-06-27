package db

import (
	"fmt"

	"github.com/cp-Coder/khelo/pkg/utils"
	"gorm.io/driver/mysql" // load driver for Mysql
	"gorm.io/gorm"
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*gorm.DB, error) {
	// Build Mysql connection URL.
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	// Define database connection for Mysql.
	db, err := gorm.Open(mysql.Open(mysqlConnURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	return db, nil
}
