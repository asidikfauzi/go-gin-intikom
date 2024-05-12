package auth

import (
	"asidikfauzi/go-gin-intikom/domain"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
}

type AuthDomain struct {
	AuthService domain.AuthService `inject:"auth_service"`
}
