package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS_RETURN     int = 200
	ERROR_LOGIN_PARAMS int = 40001
	ERROR_LOGIN_CHECKED int = 40002
)

var msgMap = map[int]string{
	SUCCESS_RETURN:     "请求成功",
	ERROR_LOGIN_PARAMS: "参数有误",
	ERROR_LOGIN_CHECKED: "登录验证失败",
}

//result format
type ReturnModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//success json
func ReturnSuccess(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, ReturnModel{SUCCESS_RETURN, msgMap[SUCCESS_RETURN], data})

}

//error json
func ReturnError(c *gin.Context, code int, data interface{}) {

	errMsg, ok := msgMap[code];
	if !ok {
		errMsg = "unknown error"
	}

	c.JSON(http.StatusOK, ReturnModel{code, errMsg, data})
}

//view render
func Render(c *gin.Context, tmpl string, data interface{}) {
	c.HTML(http.StatusOK, tmpl, data)
}
