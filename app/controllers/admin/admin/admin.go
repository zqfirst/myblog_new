package admin

import (
	"github.com/gin-gonic/gin"
	"owner/app/models"
)

func AdminList(c *gin.Context) {
	models.User.GetUser("1")
}
