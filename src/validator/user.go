package validator

import (
	"github.com/go-playground/validator/v10"
	"log"
)

func init() {
	if err := val.RegisterValidation("UserName", UserName); err != nil {
		log.Fatal("validator UserName error ")
	}
}

// UserName 验证用户名必须大于1个字符
var UserName validator.Func = func(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	return ok && len(name) > 1
}
