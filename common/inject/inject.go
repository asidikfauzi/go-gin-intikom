package inject

import (
	"asidikfauzi/go-gin-intikom/common/database"
	"asidikfauzi/go-gin-intikom/repository/postgres"
	"asidikfauzi/go-gin-intikom/route"
	"asidikfauzi/go-gin-intikom/service"
	"github.com/facebookgo/inject"
	"log"
)

type InjectData struct {
	Routes *route.RouteService
}

func DependencyInjection(liq InjectData) {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// POSTGRES
	userPostgres := postgres.NewUserPostgres(db)

	// SERVICES
	authService := service.NewAuthService(userPostgres)
	userService := service.NewUserService(userPostgres)

	dependencies := []*inject.Object{
		{Value: authService, Name: "auth_service"},
		{Value: userService, Name: "user_service"},
	}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
			&inject.Object{Value: liq.Routes.AuthController, Name: "auth_controller"},
			&inject.Object{Value: liq.Routes.UserController, Name: "user_controller"},
		)
	}

	// DEPENDENCY INJECTION
	var g inject.Graph
	if err = g.Provide(dependencies...); err != nil {
		log.Fatal("Failed Inject Dependencies", err)
	}

	if err = g.Populate(); err != nil {
		log.Fatal("Failed Populate Inject Dependencies", err)
	}

}
