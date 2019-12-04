package login

import (
	"myblog/app/services"
	"myblog/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login", gin.H{})
}

type LoginForm struct {
	Username string `from:"username" form:"username" binding:"required"`
	Password string `from:"password" form:"password" binding:"required"`
}

func VLogin(c *gin.Context) {
	var params LoginForm

	if err := c.ShouldBind(&params);err != nil {
		utils.ReturnError(c, utils.ERROR_LOGIN_PARAMS, nil)
		return
	}

	if login :=services.CheckAdmin(params.Username, params.Password); !login{
		utils.ReturnError(c, utils.ERROR_LOGIN_CHECKED, nil)
		return
	}

	utils.SetCookie(c, utils.GlobalConfig.Http.Cookie.TokenKey, utils.Md5(params.Username), "/", 10,true)
	utils.SetCookie(c, "username", params.Username, "/", 10,true)

	data := map[string]string{"url":"/"}
	utils.ReturnSuccess(c,data)
}
