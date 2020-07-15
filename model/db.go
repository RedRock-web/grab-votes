// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:24
package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	db, err := gorm.Open("mysql", "mysql", "root:mima@tcp(127.0.0."+
		"1:3306)/grab_votes?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		log.Panicf("Panic while connecting the gorm. Error: %s", err)
	}
	if db.HasTable(&Movie{}) {
		db.AutoMigrate(&Movie{})
	} else {
		db.CreateTable(&Movie{})
	}

	if db.HasTable(&Order{}) {
		db.AutoMigrate(&Order{})
	} else {
		db.CreateTable(&Order{})
	}

	DB = db
}
