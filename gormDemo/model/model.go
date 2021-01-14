package model

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	config, err := goconfig.LoadConfigFile(`./gormDemo/config/database.ini`)
	if err != nil {
		panic(err)
	}
	mysql, err := config.GetSection("mysql")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysql["username"], mysql["password"], mysql["address"], mysql["port"], mysql["database"], mysql["charset"])
	Db, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	Db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	Db.LogMode(true)       //打印sql语句
	//开启连接池
	Db.DB().SetMaxIdleConns(100)   //最大空闲连接
	Db.DB().SetMaxOpenConns(10000) //最大连接数
	Db.DB().SetConnMaxLifetime(30) //最大生存时间(s)
	Db = Db.Unscoped()
}
