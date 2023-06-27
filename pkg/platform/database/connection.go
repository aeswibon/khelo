package db

import (
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DatabaseClient struct to define the database
type DatabaseClient struct {
	client *gorm.DB
}

var databaseClient *DatabaseClient

// OpenDBConnection func for opening database connection.
func OpenDBConnection() error {
	// Define Database connection variables.
	var (
		err error
		db  *gorm.DB
	)

	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "mysql":
		db, err = MysqlConnection()
	default:
		db, err = PostgreSQLConnection()
	}

	if err != nil {
		log.Panic(err)
	}

	if sql, err := db.DB(); sql.Ping() != nil || err != nil {
		log.Panic("error: database connection failed")
	}

	// Set the global database client.
	databaseClient = &DatabaseClient{
		client: db,
	}
	return nil
}

// CloseDBConnection func for closing database connection.
func CloseDBConnection() error {
	// Close the database connection.
	db, err := databaseClient.client.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetDBClient function to return the global db
func GetDBClient() *DatabaseClient {
	if databaseClient == nil {
		log.Fatal("error: database client is not initialized")
	}
	return databaseClient
}

// Migrate func to migrate the database.
func (d *DatabaseClient) Migrate(models ...interface{}) error {
	return d.client.AutoMigrate(models...)
}

// Create func to create a new record in the database.
//
//	// Create a new user record
//	db.GetDBClient().Create(&User{})
func (d *DatabaseClient) Create(model interface{}) error {
	return d.client.Create(model).Error
}

// Update func to update a record in the database.
//
//	// Update user record based on external id
//	db.GetDBClient().Update(&User{}, "external_id = ?", id)
func (d *DatabaseClient) Update(model interface{}, query ...interface{}) error {
	query = append(query, "deleted = ?", false)
	return d.client.First(model, query).Save(model).Error
}

// Delete func to soft delete a record in the database
//
//	// Delete user record based on external id
//	db.GetDBClient().Delete(&User{}, "external_id = ?", id)
func (d *DatabaseClient) Delete(model interface{}, query ...interface{}) error {
	query = append(query, "deleted = ?", false)
	return d.client.First(model, query).UpdateColumn("deleted", true).Error
}

// Find func to find a record in the database.
//
//	// Find user record based on external id
//	db.GetDBClient().Find(&User{}, "external_id = ?", id)
func (d *DatabaseClient) Find(model interface{}, query ...interface{}) *gorm.DB {
	query = append(query, "deleted = ?", false)
	return d.client.Find(model, query)
}

// First func to find a record in the database.
//
//	// Find user record based on external id
//	db.GetDBClient().First(&User{}, "external_id = ?", id)
func (d *DatabaseClient) First(model interface{}, query ...interface{}) error {
	query = append(query, "deleted = ?", false)
	return d.client.First(model, query).Error
}

// Select func to select a record in the database.
//
//	// Select user record based on external id
//	db.GetDBClient().Select(&User{}, "external_id = ?", id)
func (d *DatabaseClient) Select(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.client.Select(query, args)
}

// Preload func to preload a record in the database.
//
//	// get all users, and preload all non-cancelled orders
//	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (d *DatabaseClient) Preload(model interface{}, query string, args ...interface{}) *gorm.DB {
	return d.client.Preload(query, args)
}

// Where func to find a record in the database.
//
//	// Find user record based on external id
//	db.GetDBClient().Where(&User{}, "external_id = ?", id)
func (d *DatabaseClient) Where(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.client.Where(query, args)
}

// Limit func to limit the number of records returned.
//
//	// Limit the number of records returned
//	db.Limit(2).First(&user)
func (d *DatabaseClient) Limit(limit int) *gorm.DB {
	d.client.Statement.AddClause(clause.Limit{Limit: &limit})
	return d.client
}

// Offset func to offset the number of records returned.
//
//	// Offset the number of records returned
//	db.Offset(2).First(&user)
func (d *DatabaseClient) Offset(offset int) *gorm.DB {
	d.client.Statement.AddClause(clause.Limit{Offset: offset})
	return d.client
}

// Model func to set the model for the query.
//
//	// Set the model for the query
//	db.Model(&User{}).Where("name = ?", "jinzhu").Update("name", "hello")
func (d *DatabaseClient) Model(model interface{}) *gorm.DB {
	return d.client.Model(model)
}
