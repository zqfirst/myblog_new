package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"owner/app/utils"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login", gin.H{})
}

type LoginForm struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"username" json:"password" binding:"required"`
}

func VLogin(c *gin.Context) {
	params := LoginForm{}
	if err := c.ShouldBind(&params); err != nil {
		utils.ReturnError(c, utils.ERROR_LOGIN_PARAMS, nil)
	}



	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "onion", //这个是value要自己生成？？规则自定就可以？
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "登录成功",
	})
}
