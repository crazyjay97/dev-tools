package db

import (
	"code-generator/load"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(load.Config.Db.Dialect, load.Config.Db.Url)
	if nil != err {
		fmt.Println(err)
		panic("Database Connect Error")
	}
	if 0 < load.Config.Db.MaxIdleCons {
		db.DB().SetMaxIdleConns(load.Config.Db.MaxIdleCons)
	}
	if 0 < load.Config.Db.MaxOpenCon {
		db.DB().SetMaxOpenConns(load.Config.Db.MaxOpenCon)
	}
	db.LogMode(load.Config.Db.PrintLog)
	DB = db
}
