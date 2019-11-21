package main

import (
	"owner/app/models"
	"owner/router"
)

func main() {

	//init router
	r := router.InitRouter()

	//init db
	models.InitDb()

	r.Run(":9600")
}

