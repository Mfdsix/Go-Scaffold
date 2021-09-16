package app

import (
	"go-scaffold/app/handlers"
	"net/http"
)

func (app *App) GetAllGenders() http.HandlerFunc {
	return handlers.GetAllGenders(app.DB)
}
