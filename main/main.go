package main

import (
	"fmt"
	"owner/app/models"
	"owner/app/utils"
	"owner/router"
)

func main() {

	defer recover()

	//init router
	r := router.InitRouter()

	//init Config
	utils.InitConfig()

	//init db
	models.InitDb()

	//r.Run(":9999")
	r.Run(fmt.Sprintf(":%d", utils.GlobalConfig.Http.Port))
}

func recover()  {
	
}