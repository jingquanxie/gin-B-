package main

import (
	"github.com/gin-gonic/gin"
	"thsit.com/ginessential/controller"
)

func CollectRoute(r *gin.Engine){
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
}
