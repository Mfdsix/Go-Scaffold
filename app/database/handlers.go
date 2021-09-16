package database

import (
	"go-scaffold/app/database/handlers"
	"go-scaffold/app/models"
)

func (d *DB) GetAllGenders() []models.DBGender {
	return handlers.GetAllGenders(d.db)
}
