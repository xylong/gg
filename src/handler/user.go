package handler

import (
	"gg/src/data/getter"
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

	u.Mutate(user.WithId(1))
	if u.ID > 0 {
		R(ctx)("ok", 0, u)(OK)
	} else {
		R(ctx)("fail", 10001, nil)(Error)
	}
}

func UserList(ctx *gin.Context) {
	R(ctx)("ok", 0, getter.UserGetter.List())(OK)
}

func UserInfo(ctx *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required,gt=0"`
	}{}
	result.Result(ctx.ShouldBindUri(id)).Unwrap()
	R(ctx)("ok", 0, getter.UserGetter.Show(id.ID).Unwrap())(OK)
}
