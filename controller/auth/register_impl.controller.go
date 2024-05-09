package auth

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"
)

func (a *AuthDomain) Register(c *gin.Context) {
	startTime := time.Now()

	var (
		req model.ReqUser
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

	if !helper.ErrorValidatingPassword(c, req.Password, req.ConfirmPassword) {
		return
	}

	res, err := a.AuthService.Register(c, req, startTime)
	if err != nil {
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, http.StatusText(http.StatusCreated), map[string]interface{}{helper.Success: res}, startTime)
	return
}
