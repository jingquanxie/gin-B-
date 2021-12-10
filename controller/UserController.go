package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"thsit.com/ginessential/model"
	"thsit.com/ginessential/util"
)

func Register(c *gin.Context) {

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
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if model.IsTelephoneExits(telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在了"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	model.SaveUser(newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
