package validator

import (
	"github.com/go-playground/validator/v10"
)

func init() {
	// 注册自定义验证
	registerValidation("UserName", UserName("required,min=2,max=5").toFunc())
}

type UserName string

func (u UserName) toFunc() validator.Func {
	validatorError["UserName"] = "用户名在2-5位之间"
	return func(fl validator.FieldLevel) bool {
		name, ok := fl.Field().Interface().(string)
		if ok {
			return u.validate(name)
		}
		return false
	}
}

func (u UserName) validate(v string) bool {
	if err := val.Var(v, string(u)); err != nil {
		return false
	}
	return true
}

// UserName 验证用户名必须大于1个字符
//var UserName validator.Func = func(fl validator.FieldLevel) bool {
//	name, ok := fl.Field().Interface().(string)
//	return ok && len(name) > 1
//}
