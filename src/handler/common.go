package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// JsonResult 返回json
type JsonResult struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

// NewJsonResult 构造函数
func NewJsonResult(msg string, code int, data interface{}) *JsonResult {
	return &JsonResult{Msg: msg, Code: code, Data: data}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewJsonResult("", 0, nil)
		},
	}
}

// ResultFunc 结果处理函数
type ResultFunc func(msg string, code int, data interface{}) func(output Output)

// Output 输出
type Output func(ctx *gin.Context, any interface{})

func R(ctx *gin.Context) ResultFunc {
	return func(msg string, code int, data interface{}) func(output Output) {
		r := ResultPool.Get().(*JsonResult)
		defer ResultPool.Put(r)

		r.Msg = msg
		r.Code = code
		r.Data = data

		return func(output Output) {
			output(ctx, r)
		}
	}
}

func OK(ctx *gin.Context, any interface{}) {
	ctx.JSON(http.StatusOK, any)
}

func Error(ctx *gin.Context, any interface{}) {
	ctx.JSON(http.StatusBadRequest, any)
}

func OK2String(ctx *gin.Context, any interface{}) {
	ctx.String(http.StatusOK, fmt.Sprintf("%v", any))
}
