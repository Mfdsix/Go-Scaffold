package database

import (
	"fmt"
	"go-scaffold/app/helpers"
)

var (
	APP_DB_USERNAME   string = helpers.GetEnvVal("APP_DB_USERNAME")
	APP_DB_PASSWORD   string = helpers.GetEnvVal("APP_DB_PASSWORD")
	APP_DB_NAME       string = helpers.GetEnvVal("APP_DB_NAME")
	APP_DB_HOST       string = helpers.GetEnvVal("APP_DB_HOST")
	APP_DB_PORT       string = helpers.GetEnvVal("APP_DB_PORT")
	CONNECTION_STRING string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", APP_DB_USERNAME, APP_DB_PASSWORD, APP_DB_HOST, APP_DB_PORT, APP_DB_NAME)
)
