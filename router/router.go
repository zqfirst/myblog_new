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
	r.LoadHTMLGlob("app/templates/admin/*.tmpl")
	r.LoadHTMLGlob("app/templates/admin/*/*.tmpl")
	//r.LoadHTMLGlob("app/templates/frontend/*.tmpl")
	r.LoadHTMLGlob("app/templates/frontend/*/*.tmpl")
	r.StaticFile("/favicon.ico","static/favicon.ico")

	admin := r.Group("/admin")
	{
		admin.GET("/login", login.Login)
	}
	return r
}
