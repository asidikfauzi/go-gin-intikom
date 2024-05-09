package service

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	"reflect"
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

func (s *UserService) ShowUser(c *gin.Context, id int, startTime time.Time) (user model.GetUser, err error) {

	getUser, err := s.userPg.FindById(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	user.ID = getUser.ID
	user.Name = getUser.Name
	user.Email = getUser.Email

	return
}

func (s *UserService) UpdateUser(c *gin.Context, id int, startTime time.Time) (message string, err error) {
	var (
		req  model.ReqUser
		user model.Users
	)

	getUser, err := s.userPg.FindById(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
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

	if req.Password != "" && req.OldPassword != "" {
		err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(req.OldPassword))
		if err != nil {
			helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.PasswordIncorrect}, startTime)
			return
		}

		if !helper.ErrorValidatingPassword(c, req.Password, req.ConfirmPassword) {
			err = errors.New("Password not valid")
			return
		}

		hashPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if errHash != nil {
			helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: errHash.Error()}, startTime)
			return
		}

		user.Password = string(hashPassword)
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	timeNow := time.Now()
	user.ID = getUser.ID
	user.UpdatedAt = &timeNow

	fmt.Println(user)

	err = s.userPg.Update(user)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	return helper.SuccessUpdatedData, nil
}
