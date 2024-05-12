package auth

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/log"
	"asidikfauzi/go-gin-intikom/config"
	"asidikfauzi/go-gin-intikom/model"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func (a *AuthDomain) GoogleLogin(c *gin.Context) {

	googleConf := config.SetUpConfig()
	url := googleConf.AuthCodeURL(helper.GetEnv("GOOGLE_STATE"))

	c.Redirect(http.StatusSeeOther, url)
}

func (a *AuthDomain) GoogleCallback(c *gin.Context) {

	startTime := time.Now()

	state := c.Query("state")
	code := c.Query("code")

	if state != helper.GetEnv("GOOGLE_STATE") {
		helper.ResponseAPI(c, false, http.StatusForbidden, http.StatusText(http.StatusForbidden), map[string]interface{}{helper.Error: "State doesnt match."}, startTime)
		return
	}

	googleConf := config.SetUpConfig()
	tokenGoogle, err := googleConf.Exchange(context.Background(), code)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tokenGoogle.AccessToken)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	userData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	var user model.UserDataGoogle
	err = json.Unmarshal(userData, &user)
	if err != nil {
		log.Error(err)
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), map[string]interface{}{helper.Error: err.Error()}, startTime)
		return
	}

	token, err := a.AuthService.LoginGoogle(c, user, startTime)
	if err != nil {
		log.Error(err)
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), map[string]interface{}{helper.Success: helper.SuccessLogin}, token, startTime)
	return

}
