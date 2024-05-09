package task

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func (u *TaskDomain) GetTasks(c *gin.Context) {
	startTime := time.Now()

	queryParams := model.ParamPaginate{
		Page:      1,
		Limit:     0,
		OrderBy:   "id",
		Direction: "asc",
	}

	if search := c.Query("search"); search != "" {
		queryParams.Search = search
	}

	if page := c.Query("page"); page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
			return
		}
		queryParams.Page = pageInt

	}
	if limit := c.Query("limit"); limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
			return
		}
		queryParams.Limit = limitInt
	}
	if orderBy := c.Query("orderBy"); orderBy != "" {
		queryParams.OrderBy = orderBy
	}
	if direction := c.Query("direction"); direction != "" {
		queryParams.Direction = direction
	}

	tasks, paginate, err := u.TaskService.GetTasks(c, queryParams, startTime)
	if err != nil {
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: helper.SuccessGetData}, tasks, paginate, startTime)
	return
}

func (u *TaskDomain) ShowTask(c *gin.Context) {
	startTime := time.Now()

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	user, err := u.TaskService.ShowTask(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: helper.SuccessGetData}, user, startTime)
	return
}

func (a *TaskDomain) CreateTask(c *gin.Context) {
	startTime := time.Now()

	var (
		req model.ReqTask
		ve  validator.ValidationErrors
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		if errors.As(err, &ve) {
			out := make(map[string]string, len(ve))
			for i, fe := range ve {
				field := reflect.TypeOf(req).Field(i).Tag.Get("json")
				out[field] = helper.ValidateTag(fe)
			}
			helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), map[string]interface{}{helper.Error: out}, startTime)
			return
		}
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := a.TaskService.CreateTask(c, req, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, http.StatusText(http.StatusCreated), map[string]interface{}{helper.Success: res}, startTime)
	return
}

func (u *TaskDomain) UpdateTask(c *gin.Context) {
	startTime := time.Now()

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := u.TaskService.UpdateTask(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: res}, startTime)
	return
}

func (u *TaskDomain) DeleteTask(c *gin.Context) {
	startTime := time.Now()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := u.TaskService.DeleteTask(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: res}, startTime)
	return
}
