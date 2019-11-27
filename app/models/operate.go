package models

import "fmt"

type operate struct {
	baseDao
}

func (this operate) QueryAll() {
}

func (this operate) QueryOne(id string, a interface{}) {
	res , e := dao.r.Query("select * from mg_user")
	if e != nil{
		panic(e)
	}

	fmt.Println(res)
}

