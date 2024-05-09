package service

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math"
	"net/http"
	"reflect"
	"time"
)

type TaskService struct {
	taskPg domain.TaskPostgres
	userPg domain.UserPostgres
}

func NewTaskService(tp domain.TaskPostgres, up domain.UserPostgres) domain.TaskService {
	return &TaskService{
		taskPg: tp,
		userPg: up,
	}
}

func (s *TaskService) GetTasks(c *gin.Context, param model.ParamPaginate, startTime time.Time) (tasks []model.GetTask, paginate model.ResponsePaginate, err error) {
	offset := helper.GetOffset(param.Page, param.Limit)
	param.Offset = offset

	var totalData int64
	tasks, totalData, err = s.taskPg.GetAll(param)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	totalPages := 1
	if param.Limit > 0 {
		totalPages = int(math.Ceil(float64(totalData) / float64(param.Limit)))
	}

	paginate = model.ResponsePaginate{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return
}

func (s *TaskService) ShowTask(c *gin.Context, id int, startTime time.Time) (task model.GetTask, err error) {

	getTask, err := s.taskPg.FindById(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	getUser, err := s.userPg.FindById(int(getTask.UserID))
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	user := model.GetUser{
		ID:    getUser.ID,
		Name:  getUser.Name,
		Email: getUser.Email,
	}

	task.ID = getTask.ID
	task.UserID = getTask.UserID
	task.Title = getTask.Title
	task.Description = getTask.Description
	task.Status = getTask.Status
	task.User = user

	return
}

func (s *TaskService) CreateTask(c *gin.Context, req model.ReqTask, startTime time.Time) (message string, err error) {

	existsTitle := s.taskPg.TitleExists(req.Title)
	if existsTitle {
		err = fmt.Errorf(helper.AlreadyExists, "Title")
		errMessage := make(map[string]string)
		errMessage["title"] = err.Error()
		helper.ResponseAPI(c, false, http.StatusConflict, http.StatusText(http.StatusConflict), map[string]interface{}{helper.Error: errMessage}, startTime)
		return
	}

	_, err = s.userPg.FindById(int(req.UserID))
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	task := model.Tasks{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
	}

	err = s.taskPg.Create(&task)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	return helper.SuccessCreatedData, nil
}

func (s *TaskService) UpdateTask(c *gin.Context, id int, startTime time.Time) (message string, err error) {
	var (
		req  model.ReqTask
		task model.Tasks
	)

	getTask, err := s.taskPg.FindById(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: "Task " + err.Error()}, startTime)
		return
	}

	_, err = s.userPg.FindById(int(getTask.UserID))
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: "User " + err.Error()}, startTime)
		return
	}

	var ve validator.ValidationErrors
	if err = c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &ve) {
			out := make(map[string]string, len(ve))
			for i, fe := range ve {
				if fe.Tag() == "required" {
					continue
				}
				field := reflect.TypeOf(req).Field(i).Tag.Get("json")
				out[field] = helper.ValidateTag(fe)
			}
			if len(out) > 0 {
				helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), map[string]interface{}{helper.Error: out}, startTime)
				return
			}
		}
	}

	if req.Title != "" {
		task.Title = req.Title
	}

	if req.UserID != 0 {
		task.UserID = req.UserID
	}

	if req.Description != "" {
		task.Description = req.Description
	}

	if req.Status != "" {
		task.Status = req.Status
	}

	timeNow := time.Now()
	task.ID = getTask.ID
	task.UpdatedAt = &timeNow

	err = s.taskPg.Update(task)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	return helper.SuccessUpdatedData, nil
}

func (s *TaskService) DeleteTask(c *gin.Context, id int, startTime time.Time) (message string, err error) {
	var (
		task model.Tasks
	)

	getTask, err := s.taskPg.FindById(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	timeNow := time.Now()
	task.ID = getTask.ID
	task.DeletedAt = &timeNow

	err = s.taskPg.Delete(task)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	return helper.SuccessDeletedData, nil
}
