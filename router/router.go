package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"owner/app/controllers/admin/login"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.OPTIONS("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})

	//load static file
	//r.LoadHTMLGlob("app/templates/*/*.tmpl")
	r.LoadHTMLGlob("app/templates/*/*/*.tmpl")

	//load static file
	r.StaticFile("/favicon.ico", "static/favicon.ico")
	r.Static("/static","./app/static")

	setAdminRouter(r)

	return r
}

//admin url
func setAdminRouter(r *gin.Engine) {
	admin := r.Group("/admin",checkAdminLogin())
	{
		admin.GET("/login", login.Login)
		admin.POST("/login", login.VLogin)
	}
}

//check login
func checkAdminLogin() gin.HandlerFunc{
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			if value == "onion" {
				c.Next()
				return
			}
		}
		if url := c.Request.URL.String(); url == "/admin/login" {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
