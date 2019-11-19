package main

import (
	"fmt"
	"owner/router"
)

func main() {
	fmt.Println("aaa")
	r := router.InitRouter()
	r.Run(":9899")
}

