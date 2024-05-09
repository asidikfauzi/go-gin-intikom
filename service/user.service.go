package service

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
)

type UserService struct {
	userPg domain.UserPostgres
}

func NewUserService(up domain.UserPostgres) domain.UserService {
	return &UserService{
		userPg: up,
	}
}

func (s *UserService) GetUsers(c *gin.Context, param model.ParamPaginate, startTime time.Time) (users []model.GetUser, paginate model.ResponsePaginate, err error) {
	offset := helper.GetOffset(param.Page, param.Limit)
	param.Offset = offset

	var totalData int64
	users, totalData, err = s.userPg.GetAll(param)
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
