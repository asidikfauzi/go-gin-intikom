package domain

import (
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	TaskPostgres interface {
		GetAll(param model.ParamPaginate) (users []model.GetTask, count int64, err error)
		FindById(id int) (user model.Tasks, err error)
		IdExists(id int) bool
		TitleExists(title string) bool
		Create(user *model.Tasks) error
		Update(user model.Tasks) error
		Delete(user model.Tasks) error
	}

	TaskService interface {
		GetTasks(c *gin.Context, param model.ParamPaginate, startTime time.Time) (users []model.GetTask, paginate model.ResponsePaginate, err error)
		ShowTask(c *gin.Context, id int, startTime time.Time) (user model.GetTask, err error)
		CreateTask(c *gin.Context, req model.ReqTask, startTime time.Time) (message string, err error)
		UpdateTask(c *gin.Context, id int, startTime time.Time) (message string, err error)
		DeleteTask(c *gin.Context, id int, startTime time.Time) (message string, err error)
	}
)
