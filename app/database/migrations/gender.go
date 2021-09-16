package migrations

import (
	"go-scaffold/app/models"
	"log"

	"github.com/jinzhu/gorm"
)

func MigGender(db *gorm.DB) {
	log.Output(1, "--- Migrating Gender Table")
	db.AutoMigrate(&models.DBGender{})
	db.Delete(&models.DBGender{})
	log.Output(1, "--- Populating Gender Table")
	db.Create(&models.DBGender{
		ID:     0,
		Gender: "Pria",
	})
	db.Create(&models.DBGender{
		ID:     0,
		Gender: "Wanita",
	})
	db.Create(&models.DBGender{
		ID:     0,
		Gender: "Lainnya",
	})
}
