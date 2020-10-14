package main

import (
	"gg/src/common"
	"gg/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", handler.GetUser)
	r.Use(common.ErrorHandler())
	r.POST("/user", handler.UserAdd)
	r.Run()
}
