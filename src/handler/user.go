package handler

import (
	"gg/src/model/user"
	"gg/src/result"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	u := user.NewUser(user.WithId(1)).Mutate(user.WithName("jj"))
	ctx.JSON(200, u)
}

func UserAdd(ctx *gin.Context) {
	u := user.NewUser()
	result.Result(ctx.ShouldBind(u)).Unwrap()
	//info := result.Result(test.GetInfo(u.ID)).Unwrap()

	if u.ID > 10 {
		R(ctx)("ok", 0, u)(OK)
	} else {
		R(ctx)("fail", 10001, nil)(Error)
	}
}
