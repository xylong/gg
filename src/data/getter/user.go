package getter

import (
	"fmt"
	"gg/src/db"
	"gg/src/model/user"
	"gg/src/result"
)

type IUserGetter interface {
	List() []*user.User
	Show(id int) *result.ErrorResult
}

type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

func (g *UserGetterImpl) List() (users []*user.User) {
	db.Orm.Find(&users)
	return
}

func (g *UserGetterImpl) Show(id int) *result.ErrorResult {
	u := user.NewUser()
	r := db.Orm.First(u, id)
	if r.Error != nil || r.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found user:%d", id))
	}
	return result.Result(u, nil)
}
