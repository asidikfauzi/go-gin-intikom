package route

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/middleware"
	"asidikfauzi/go-gin-intikom/controller/auth"
	"asidikfauzi/go-gin-intikom/controller/user"
	"github.com/gin-gonic/gin"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	AuthController auth.AuthController `inject:"auth_controller"`
	UserController user.UserController `inject:"user_controller"`
}

func InitPackage() *RouteService {
	return &RouteService{
		AuthController: &auth.AuthDomain{},
		UserController: &user.UserDomain{},
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
				auth.POST("/register", r.AuthController.Register)
			}

			users := prefix.Group("/users")
			users.Use(middleware.JWTMiddleware())
			{
				users.GET("", r.UserController.GetUsers)
				users.GET(":id", r.UserController.ShowUser)
				users.PUT(":id", r.UserController.UpdateUser)
				//users.DELETE(":id", r.UserController.DeleteUser)
			}
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
