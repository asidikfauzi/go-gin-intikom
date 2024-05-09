package route

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/controller/auth"
	"github.com/gin-gonic/gin"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	AuthController auth.AuthController `inject:"auth_controller"`
}

func InitPackage() *RouteService {
	return &RouteService{
		AuthController: &auth.AuthDomain{},
	}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	api := router.Group("/api")
	{
		prefix := api.Group("/v1")
		{
			auth := prefix.Group("/auth")
			{
				auth.POST("/login", r.AuthController.Login)
			}
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
