package user

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (u *UserDomain) GetUsers(c *gin.Context) {
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
		queryParams.Page = limitInt
	}
	if orderBy := c.Query("orderBy"); orderBy != "" {
		queryParams.OrderBy = orderBy
	}
	if direction := c.Query("direction"); direction != "" {
		queryParams.Direction = direction
	}

	users, paginate, err := u.UserService.GetUsers(c, queryParams, startTime)
	if err != nil {
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: helper.SuccessGetData}, users, paginate, startTime)
	return
}

func (u *UserDomain) ShowUser(c *gin.Context) {
	startTime := time.Now()

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	user, err := u.UserService.ShowUser(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: helper.SuccessGetData}, user, startTime)
	return
}

func (u *UserDomain) UpdateUser(c *gin.Context) {
	startTime := time.Now()

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := u.UserService.UpdateUser(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: res}, startTime)
	return
}

func (u *UserDomain) DeleteUser(c *gin.Context) {
	startTime := time.Now()
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := u.UserService.DeleteUser(c, id, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: res}, startTime)
	return
}
