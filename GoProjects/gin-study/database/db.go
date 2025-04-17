package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UvsDB *gorm.DB

func InitDB() {
	initUserDB()
	initPtDB()
}

func initUserDB() {
	username := ""
	password := ""
	host := ""
	port := 3306
	dbname := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
	var err error
	UvsDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败:" + err.Error())
	}
}

var PtDB *gorm.DB

func initPtDB() {
	username := ""
	password := ""
	host := ""
	port := 3306
	dbname := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
	var err error
	PtDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败:" + err.Error())
	}
}
