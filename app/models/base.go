package models

import (
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
	d baseDao
)

//db init
func InitDb() {

	//init
	d = baseDao{}

	//read
	d.r = engine(utils.GlobalConfig.Mysql.Rdsn)
	d.rDb = d.r.DB()

	//write
	d.w = engine(utils.GlobalConfig.Mysql.Wdsn)
	d.wDb = d.w.DB()
}

func engine(dsn string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(utils.GlobalConfig.Mysql.MaxOpenConnections)
	engine.SetMaxIdleConns(utils.GlobalConfig.Mysql.MaxIdleConnections)
	return engine
}
