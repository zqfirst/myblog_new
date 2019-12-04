package main

import (
	"fmt"
	"myblog/app/models"
	"myblog/router"
	"myblog/utils"
)

func main() {

	defer resolve()

	//init Config
	utils.InitConfig()

	//init db
	models.InitDb()

	//init router
	r := router.InitRouter()

	//r.Run(":9999")
	r.Run(fmt.Sprintf(":%d", utils.GlobalConfig.Http.Port))
}

func resolve() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
