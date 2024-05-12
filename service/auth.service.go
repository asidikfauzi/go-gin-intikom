package service

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/jwt"
	"asidikfauzi/go-gin-intikom/common/log"
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService struct {
	userPg domain.UserPostgres
}

func NewAuthService(up domain.UserPostgres) domain.AuthService {
	return &AuthService{
		userPg: up,
	}
}

func (s *AuthService) Login(c *gin.Context, req model.ReqLogin, startTime time.Time) (token string, err error) {
	user, err := s.userPg.FindByEmail(req.Email)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.PasswordIncorrect}, startTime)
		return
	}

	token, err = jwt.GetToken(req.Email)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.KeyInvalid}, startTime)
		return
	}

	return
}

func (s *AuthService) Register(c *gin.Context, req model.ReqUser, startTime time.Time) (message string, err error) {

	existsUser := s.userPg.EmailExists(req.Email)
	if existsUser {
		err = fmt.Errorf(helper.AlreadyExists, "Email")
		errMessage := make(map[string]string)
		errMessage["email"] = err.Error()
		helper.ResponseAPI(c, false, http.StatusConflict, http.StatusText(http.StatusConflict), map[string]interface{}{helper.Error: errMessage}, startTime)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	user := model.Users{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
	}

	err = s.userPg.Create(&user)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	return helper.SuccessCreatedData, nil
}

func (s *AuthService) LoginGoogle(c *gin.Context, req model.UserDataGoogle, startTime time.Time) (token string, err error) {

	existsUser := s.userPg.EmailExists(req.Email)
	if !existsUser {
		passwordDefaultGoogle := helper.GetEnv("PASSWORD_DEFAULT_GOOGLE")
		hashPassword, errHash := bcrypt.GenerateFromPassword([]byte(passwordDefaultGoogle), 10)
		if errHash != nil {
			log.Error(err)
			helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
			return
		}

		user := model.Users{
			Name:      req.Name,
			Email:     req.Email,
			Password:  string(hashPassword),
			CreatedAt: time.Now(),
		}

		err = s.userPg.Create(&user)
		if err != nil {
			log.Error(err)
			helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
			return
		}
	}

	token, err = jwt.GetToken(req.Email)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.KeyInvalid}, startTime)
		return
	}

	return
}
