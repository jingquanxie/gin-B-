package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"thsit.com/ginessential/common"
	"thsit.com/ginessential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取header
		tokenString := c.GetHeader("Authorization")
		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足，token为空或格式错误"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足，解密失败"})
			c.Abort()
			return
		}
		userId := claims.UserId
		user := model.FindUserByID(userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足,找不到用户"})
			c.Abort()
			return
		}
		//用户存在将用户信息写入上下文
		c.Set("user",user)
		c.Next()


	}
}
