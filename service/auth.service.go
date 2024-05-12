package service

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/middleware"
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.PasswordIncorrect}, startTime)
		return
	}

	jwtKey := []byte(helper.GetEnv("JWT_KEY"))

	expTime := time.Now().Add(360 * time.Hour)
	claims := &middleware.JwtClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    helper.GetEnv("ISSUER"),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenAlgo.SignedString(jwtKey)
	if err != nil {
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
			helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), map[string]interface{}{helper.Error: err.Error()}, startTime)
			return
		}
	}

	jwtKey := []byte(helper.GetEnv("JWT_KEY"))

	expTime := time.Now().Add(360 * time.Hour)
	claims := &middleware.JwtClaim{
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    helper.GetEnv("ISSUER"),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenAlgo.SignedString(jwtKey)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.KeyInvalid}, startTime)
		return
	}

	return
}
