package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//设置cookie
func SetCookie(c *gin.Context, key, value, domain string, expire int,
	httponly bool) {
	cookie := &http.Cookie{
		Name:     key,
		Value:    value, //这个是value要自己生成？？规则自定就可以？
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}

//获取cookie
func GetCookie(c *gin.Context, key string) (*http.Cookie, error) {
	return c.Request.Cookie(key)
}
