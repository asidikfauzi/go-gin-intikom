package user

import (
	"asidikfauzi/go-gin-intikom/domain"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(c *gin.Context)
	ShowUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserDomain struct {
	UserService domain.UserService `inject:"user_service"`
}
