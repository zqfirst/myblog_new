package services

import (
	"myblog/app/models"
)

//check admin password
func CheckAdmin(admin, password string) bool {
	r := models.User.GetUserByUserName(admin)
	return r.Password == password
}

