// Package migrate to migrate the schema for the example application
package migrate

import (
	"fmt"

	gconfig "liftapp/config"
	gdatabase "liftapp/database"
	gmodel "liftapp/model"

	"liftapp/app/database/model"
)

// Load all the models
type auth gmodel.Auth
type twoFA gmodel.TwoFA
type twoFABackup gmodel.TwoFABackup
type tempEmail gmodel.TempEmail
type user model.User
type exercise model.Exercise

// DropAllTables - careful! It will drop all the tables!
func DropAllTables() error {
	db := gdatabase.GetDB()

	if err := db.Migrator().DropTable(
		&user{},
		&tempEmail{},
		&twoFABackup{},
		&twoFA{},
		&auth{},
		&exercise{},
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
func StartMigration(configure gconfig.Configuration) error {
	db := gdatabase.GetDB()

	if err := db.AutoMigrate(
		&auth{},
		&twoFA{},
		&twoFABackup{},
		&tempEmail{},
		&user{},
		&exercise{},
	); err != nil {
		return err
	}

	fmt.Println("new tables are  migrated successfully!")
	return nil
}

// SetPkFk - manually set foreign key for MySQL and PostgreSQL
func SetPkFk() error {
	db := gdatabase.GetDB()

	if !db.Migrator().HasConstraint(&user{}, "Posts") {
		err := db.Migrator().CreateConstraint(&user{}, "Posts")
		if err != nil {
			return err
		}
	}

	return nil
}
