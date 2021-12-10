package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"thsit.com/ginessential/common"
)

func main() {
	common.InitDB()
	db := common.GetDb()
	defer db.Close()

	r := gin.Default()
	CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}



