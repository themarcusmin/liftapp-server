// Package database handles connections to different
// types of databases
package database

import (
	"database/sql"
	"fmt"

	"liftapp/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Import SQLite3 database driver
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
)

// RecordNotFound record not found error message
const RecordNotFound string = "record not found"

// dbClient variable to access gorm
var dbClient *gorm.DB

var sqlDB *sql.DB
var err error

// InitDB - function to initialize db
func InitDB() *gorm.DB {
	var db = dbClient

	configureDB := config.GetConfig().Database.RDBMS

	driver := configureDB.Env.Driver
	database := configureDB.Access.DbName

	switch driver {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open(database), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.WithError(err).Panic("panic code: 155")
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful!")
		}

	default:
		log.Fatal("The driver " + driver + " is not implemented yet")
	}

	dbClient = db

	return dbClient
}

// GetDB - get a connection
func GetDB() *gorm.DB {
	return dbClient
}
