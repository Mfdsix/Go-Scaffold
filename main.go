package main

import (
	"go-scaffold/app"
)

func main() {
	a := app.New()
	a.Run(":8010")
}
