package main

import (
	"main/internel/app"
	"main/internel/router"
)

func main() {

	r := router.NewRouter()
	ap := app.NewApp(r)
	ap.Start()

}
