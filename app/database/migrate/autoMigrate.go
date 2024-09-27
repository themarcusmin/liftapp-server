// Package migrate to migrate the schema for the example application
package migrate

import (
	"fmt"
	"os"

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
type format model.Format
type muscle model.Muscle
type exerciseMuscle model.ExerciseMuscle
type program model.Program
type programDay model.ProgramDay
type programExercise model.ProgramExercise
type programEntry model.ProgramEntry
type log model.Log
type logExercise model.LogExercise
type logEntry model.LogEntry

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
		&muscle{},
		&exerciseMuscle{},
		&program{},
		&programDay{},
		&programExercise{},
		&programEntry{},
		&log{},
		&logExercise{},
		&logEntry{},
		&format{},
	); err != nil {
		return err
	}

	fmt.Println("app-level: old tables are deleted!")
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
		&muscle{},
		&exerciseMuscle{},
		&program{},
		&programDay{},
		&programExercise{},
		&programEntry{},
		&log{},
		&logExercise{},
		&logEntry{},
		&format{},
	); err != nil {
		return err
	}

	fmt.Println("app-level: new tables are  migrated successfully!")
	return nil
}

// PopulateTables
// - Using insert queries from raw sql file
func PopulateTables() error {
	dir, err := os.Getwd()
	buf, err := os.ReadFile(dir + "/app/database/raw/starter.sql")
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	sqlString := string(buf)
	db := gdatabase.GetDB()
	err = db.Exec(sqlString).Error
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}

	fmt.Println("app-level: tables have been populated successfully!")
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
