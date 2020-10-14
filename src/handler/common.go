package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type JsonResult struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

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

type ResultFunc func(msg string, code int, data interface{})

func OK(ctx *gin.Context) ResultFunc {
	return func(msg string, code int, data interface{}) {
		r := ResultPool.Get().(*JsonResult)
		defer ResultPool.Put(r)

		r.Msg = msg
		r.Code = code
		r.Data = data
		ctx.JSON(http.StatusOK, r)
	}
}
