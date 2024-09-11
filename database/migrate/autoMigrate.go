// Package migrate to migrate the schema
package migrate

import (
	"fmt"

	"liftapp/model"

	"liftapp/config"
	gdatabase "liftapp/database"
)

// Load all the models
type auth model.Auth
type twoFA model.TwoFA
type twoFABackup model.TwoFABackup
type tempEmail model.TempEmail

// DropAllTables - careful! It will drop all the tables!
func DropAllTables() error {
	db := gdatabase.GetDB()

	if err := db.Migrator().DropTable(
		&tempEmail{},
		&twoFABackup{},
		&twoFA{},
		&auth{},
	); err != nil {
		return err
	}

	fmt.Println("old tables are deleted!")
	return nil
}

// StartMigration - automatically migrate all the tables
//
// - Only create tables with missing columns and missing indexes
// - Will not change/delete any existing columns and their types
func StartMigration(configure config.Configuration) error {
	db := gdatabase.GetDB()

	if err := db.AutoMigrate(
		&auth{},
		&twoFA{},
		&twoFABackup{},
		&tempEmail{},
	); err != nil {
		return err
	}

	fmt.Println("new tables are  migrated successfully!")
	return nil
}
