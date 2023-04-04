package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListRespose[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Err     = -1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

// OKWithList 泛型结构体
func OKWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListRespose[any]{
		List:  list,
		Count: count,
	}, c)
}

func OKBase(c *gin.Context) {
	Result(Success, map[string]any{}, "成功！", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Err, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Err, map[string]any{}, msg, c)
}

func FailWithCode(code ErrCode, c *gin.Context) {
	msg, ok := ErrMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Err, map[string]any{}, "未知错误", c)
}
