package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"owner/app/utils"
	"xorm.io/core"
)

type baseDao struct {
	r *xorm.Engine
	w *xorm.Engine
	rDb *core.DB
	wDb *core.DB
}

var (
	dao baseDao
)

//db init
func InitDb() {

	//init
	dao = baseDao{}

	//read
	dao.r = engine(utils.GlobalConfig.Mysql.Rdsn)
	dao.rDb = dao.r.DB()

	fmt.Println(dao.r)

	//write
	dao.w = engine(utils.GlobalConfig.Mysql.Wdsn)
	dao.wDb = dao.w.DB()
}

func engine(dsn string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetTableMapper(core.SnakeMapper{})
	engine.SetMaxOpenConns(utils.GlobalConfig.Mysql.MaxOpenConnections)
	engine.SetMaxIdleConns(utils.GlobalConfig.Mysql.MaxIdleConnections)
	return engine
}
