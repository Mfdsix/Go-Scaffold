package app

import (
	"go-scaffold/app/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.IDB
}

func New() *App {
	log.Output(1, "Launching App")
	a := &App{
		Router: mux.NewRouter(),
	}
	a.DB = &database.DB{}

	var err error
	err = a.DB.Open()
	check(err)
	// a.Migrate() if you wanna migrate on app launch

	a.initRoutes()
	return a
}

func (a *App) Migrate() {
	log.Output(1, "-- Migrating Database")
	a.DB.Migrate()
}

func (a *App) Run(addr string) {
	log.Output(1, "------------------------------------")
	log.Output(1, "App Running on "+addr)
	log.Output(1, "------------------------------------")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/genders", a.GetAllGenders()).Methods("GET")
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
