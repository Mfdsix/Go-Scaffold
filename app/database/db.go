package database

import (
	"go-scaffold/app/database/migrations"
	"go-scaffold/app/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IDB interface {
	Open() error
	Migrate() error
	Close() error
	// genders
	GetAllGenders() []models.DBGender
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Open() error {
	log.Output(1, "-- Connecting to DB")
	log.Output(1, CONNECTION_STRING)
	db, err := gorm.Open("postgres", CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Set to `true` and GORM will print out all DB queries.
	db.LogMode(false)
	d.db = db
	log.Output(1, "Connection Opened")
	return nil
}

func (d *DB) Migrate() error {
	// Automatically create tables based on models
	migrations.MigGender(d.db)
	return nil
}

func (d *DB) Close() error {
	d.db.Close()
	return nil
}
