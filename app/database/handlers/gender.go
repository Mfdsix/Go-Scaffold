package handlers

import (
	"go-scaffold/app/models"

	"github.com/jinzhu/gorm"
)

func GetAllGenders(db *gorm.DB) []models.DBGender {
	var genders []models.DBGender
	db.Find(&genders)
	return genders
}
