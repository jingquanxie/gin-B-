package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
// 初始 开启数据连接池
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database ,err: " + err.Error())
	}
	// 自动创建表-指定了一个表名
	//db.AutoMigrate(&model.User{})
	DB = db
	return db

}

func GetDb() *gorm.DB {
	return DB
}
