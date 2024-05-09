package domain

import (
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	UserPostgres interface {
		GetAll(param model.ParamPaginate) (users []model.GetUser, count int64, err error)
		FindById(id int) (user model.GetUser, err error)
		FindByEmail(email string) (user model.Users, err error)
		EmailExists(email string) bool
		Create(user *model.Users) error
	}

	AuthService interface {
		Login(c *gin.Context, req model.ReqLogin, startTime time.Time) (res string, err error)
		Register(c *gin.Context, req model.ReqRegister, startTime time.Time) (message string, err error)
	}

	UserService interface {
		GetUsers(c *gin.Context, param model.ParamPaginate, startTime time.Time) (users []model.GetUser, paginate model.ResponsePaginate, err error)
	}
)
