package db

import (
	"code-generator/internal/app/base"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(base.Config.Db.Dialect, base.Config.Db.Url)
	if nil != err {
		fmt.Println(err)
		panic("Database Connect Error")
	}
	if 0 < base.Config.Db.MaxIdleCons {
		db.DB().SetMaxIdleConns(base.Config.Db.MaxIdleCons)
	}
	if 0 < base.Config.Db.MaxOpenCon {
		db.DB().SetMaxOpenConns(base.Config.Db.MaxOpenCon)
	}
	db.LogMode(base.Config.Db.PrintLog)
	DB = db
}
