package route

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/middleware"
	"asidikfauzi/go-gin-intikom/controller/auth"
	"asidikfauzi/go-gin-intikom/controller/task"
	"asidikfauzi/go-gin-intikom/controller/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	AuthController auth.AuthController `inject:"auth_controller"`
	UserController user.UserController `inject:"user_controller"`
	TaskController task.TaskController `inject:"task_controller"`
}

func InitPackage() *RouteService {
	return &RouteService{
		AuthController: &auth.AuthDomain{},
		UserController: &user.UserDomain{},
		TaskController: &task.TaskDomain{},
	}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"loopback", "link-local", "unspecified"})

	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", r.AuthController.Login)
			auth.POST("/register", r.AuthController.Register)

			google := auth.Group("/google")
			{
				google.GET("/login", r.AuthController.GoogleLogin)
				google.GET("/callback", r.AuthController.GoogleCallback)
			}
		}

		prefix := api.Group("/v1")
		{

			users := prefix.Group("/users")
			users.Use(middleware.JWTMiddleware())
			{
				users.GET("", r.UserController.GetUsers)
				users.GET(":id", r.UserController.ShowUser)
				users.PUT(":id", r.UserController.UpdateUser)
				users.DELETE(":id", r.UserController.DeleteUser)
			}

			tasks := prefix.Group("/tasks")
			tasks.Use(middleware.JWTMiddleware())
			{
				tasks.GET("", r.TaskController.GetTasks)
				tasks.GET(":id", r.TaskController.ShowTask)
				tasks.POST("", r.TaskController.CreateTask)
				tasks.PUT(":id", r.TaskController.UpdateTask)
				tasks.DELETE(":id", r.TaskController.DeleteTask)
			}
		}

	}

	port := helper.GetEnv("APP_PORT")
	err := router.Run(":" + port)
	if err != nil {
		err = http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}

}
