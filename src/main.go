package main

import (
	"gg/src/common"
	"gg/src/model/user"
	"gg/src/result"
	"gg/src/test"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		u := user.NewUser(user.WithId(1)).Mutate(user.WithName("jj"))
		c.JSON(200, u)
	})
	r.Use(common.ErrorHandler())
	r.POST("/user", func(c *gin.Context) {
		u := user.NewUser()
		result.Result(c.ShouldBind(u)).Unwrap()
		c.JSON(200, gin.H{"msg": "ok", "data": u})
	})
	r.POST("/test", func(c *gin.Context) {
		u := user.NewUser()
		result.Result(c.ShouldBind(u)).Unwrap()
		info := result.Result(test.GetInfo(u.ID)).Unwrap()
		c.JSON(200, info)
	})
	r.Run()
}
