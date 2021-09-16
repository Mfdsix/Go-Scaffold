package handlers

import (
	"go-scaffold/app/database"
	"go-scaffold/app/helpers"
	"log"
	"net/http"
)

func GetAllGenders(db database.IDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Output(1, "-- Request: Get All Genders")
		genders := db.GetAllGenders()

		if len(genders) == 0 {
			helpers.HttpOops(w, "Genders Not Found", http.StatusNotFound)
			return
		}

		helpers.HttpOk(w, genders, http.StatusOK)
	}
}
