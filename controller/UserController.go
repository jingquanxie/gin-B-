package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"thsit.com/ginessential/model"
	"thsit.com/ginessential/util"
)

func Login(c *gin.Context) {

	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	//数据校验
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//判断手机号是否存在
	user := model.FindUserByPhone(telephone)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 422, "msg": "密码校验错误"})
		return

	}
	// 发放token
	token := "1111"
	//
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": gin.H{"token": token}})

}
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
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 422, "msg": "密码加密错误"})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	model.SaveUser(newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
