package task

import (
	"asidikfauzi/go-gin-intikom/domain"
	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetTasks(c *gin.Context)
	ShowTask(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type TaskDomain struct {
	TaskService domain.TaskService `inject:"task_service"`
}
