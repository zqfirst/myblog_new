package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login.html", gin.H{})
}
