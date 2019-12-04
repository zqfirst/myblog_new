package router

import (
	"github.com/gin-gonic/gin"
	ad "myblog/app/controllers/admin/admin"
	"myblog/app/controllers/admin/login"
	"myblog/utils"
	"net/http"
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
	r.Static("/static", "./app/static")

	setAdminRouter(r)

	return r
}

//admin url
func setAdminRouter(r *gin.Engine) {
	admin := r.Group("/admin", checkAdminLogin)
	{
		//login
		admin.GET("/login", login.Login)
		admin.POST("/login", login.VLogin)

		//admin list
		admin.GET("/list", ad.AdminList)
	}

	frontend := r.Group("/web")
	{
		frontend.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, ",%d", 22)
		})
	}
}

//check login
func checkAdminLogin(c *gin.Context) {
	var token,username string
	if cookie, err := utils.GetCookie(c, utils.GlobalConfig.Http.Cookie.
		TokenKey);
		err == nil {
		token = cookie.Value
	}

	if cookie, err := utils.GetCookie(c, "username");
		err == nil {
		username = cookie.Value
	}

	if token == utils.Md5(username){
		c.Next()
		return
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
