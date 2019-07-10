package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:zywl2018@tcp(192.168.1.82:3306)/basecloud?charset=utf8&parseTime=True&loc=Local")
	if nil != err {
		fmt.Println(err)
		panic("数据库连接异常")
	}
	//设置最大空闲连接
	db.DB().SetMaxIdleConns(10)
	//设置最大打开连接
	db.DB().SetMaxOpenConns(100)
	//log
	db.LogMode(true)
	DB = db
}
