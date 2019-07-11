package db

import (
	"code-generator/init"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(init.Config.Db.Dialect, init.Config.Db.Url)
	if nil != err {
		fmt.Println(err)
		panic("Database Connect Error")
	}
	//设置最大空闲连接
	db.DB().SetMaxIdleConns(10)
	//设置最大打开连接
	db.DB().SetMaxOpenConns(100)
	//log
	db.LogMode(true)
	DB = db
}
