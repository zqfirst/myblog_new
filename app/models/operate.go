package models

import "fmt"

type operate struct {
	baseDao
}

func (this operate) QueryAll() {
}

func (this operate) QueryOne() {
	res , e := this.r.QueryString("select * from mg_user where id=?", 1)
	if e != nil{
		panic(e)
	}

	fmt.Println(res)
}

