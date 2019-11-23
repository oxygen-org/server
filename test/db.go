/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:06:26
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"time"

	conf "github.com/oxygen-org/server/config"

	"github.com/jinzhu/gorm"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB gorm DB 对象
var DB = &gorm.DB{}

func init() {
	fmt.Println("db init")

	linkStr := conf.CONFIG.Get("SQL_URI").MustString()
	fmt.Println("**************")
	fmt.Println(linkStr)
	fmt.Println("**************")

	db, error := gorm.Open("mysql", linkStr)
	if error != nil {
		fmt.Println(error)
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	DB = db
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	db.SingularTable(true)
}

func main() {
	var d Role
	DB.First(&d)
	fmt.Println(d)
}
