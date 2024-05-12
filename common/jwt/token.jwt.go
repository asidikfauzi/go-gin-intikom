package jwt

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"asidikfauzi/go-gin-intikom/common/middleware"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetToken(email string) (token string, err error) {
	jwtKey := []byte(helper.GetEnv("JWT_KEY"))

	expTime := time.Now().Add(360 * time.Hour)
	claims := &middleware.JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    helper.GetEnv("ISSUER"),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenAlgo.SignedString(jwtKey)
	if err != nil {
		return
	}

	return
}
