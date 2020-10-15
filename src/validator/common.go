package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	val            *validator.Validate
	validatorError map[string]string
)

func init() {
	validatorError = make(map[string]string)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		val = v
	} else {
		log.Fatal("error validator")
	}
}

// registerValidation 注册自定义验证
func registerValidation(tag string, fn validator.Func) {
	if err := val.RegisterValidation(tag, fn); err != nil {
		log.Fatal(err)
	}
}

func CheckErrors(errors error) {
	if errs, ok := errors.(validator.ValidationErrors); ok {
		for _, err := range errs {
			if v, exits := validatorError[err.Tag()]; exits {
				panic(v)
			}
		}
	}
}
