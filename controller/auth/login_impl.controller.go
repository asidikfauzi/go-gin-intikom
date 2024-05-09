package auth

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"
)

func (a *AuthDomain) Login(c *gin.Context) {
	startTime := time.Now()

	var req model.ReqLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		if vErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for i, vError := range vErrors {
				field := reflect.TypeOf(req).Field(i).Tag.Get("json")
				errors[field] = vError.Tag()
			}
			helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), map[string]interface{}{helper.Error: errors}, startTime)
			return
		}
		helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	token, err := a.AuthService.Login(c, req, startTime)
	if err != nil {
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{"success": helper.SuccessLogin}, token, startTime)
	return
}
