package config

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetUpConfig() *oauth2.Config {

	conf := &oauth2.Config{
		ClientID:     helper.GetEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: helper.GetEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  helper.GetEnv("APP_HOST") + "/api/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}
