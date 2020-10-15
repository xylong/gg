package main

import (
	"gg/src/common"
	"gg/src/db"
	"gg/src/handler"
	_ "gg/src/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.New()
	r.GET("/", handler.GetUser)
	r.Use(common.ErrorHandler())
	r.POST("/users", handler.UserAdd)
	r.GET("/users", handler.UserList)
	r.GET("/users/:id", handler.UserInfo)
	r.Run()
}
