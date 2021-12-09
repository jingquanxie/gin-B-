package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据校验
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		// 如果名称没有值，默认给10位随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		//判断手机号是否存在
		if isTelephoneExits(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在了"})
			return
		}
		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		// 返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

// 生成一个随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 定一个结构体（module)
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

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
	db.AutoMigrate(&User{})
	return db

}

// 检验手机号是否存在
func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
