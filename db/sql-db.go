/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 21:47:55
 * @LastEditors: your name
 */
package db

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

	db, _ := gorm.Open("mysql", linkStr)
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
