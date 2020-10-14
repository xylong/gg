package handler

import (
	"gg/src/model/user"
	"gg/src/result"
	"gg/src/test"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	u := user.NewUser(user.WithId(1)).Mutate(user.WithName("jj"))
	ctx.JSON(200, u)
}

func UserAdd(ctx *gin.Context) {
	u := user.NewUser()
	result.Result(ctx.ShouldBind(u)).Unwrap()
	info := result.Result(test.GetInfo(u.ID)).Unwrap()
	OK(ctx)("ok", 0, info)
}
