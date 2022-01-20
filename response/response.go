package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func Response(ctx *gin.Context, httpsStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpsStatus, gin.H{"code": code, "data": data, "msg": msg})

}

func Success(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": data, "msg": msg})
}
func Fail(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": 400, "data": data, "msg": msg})
}
