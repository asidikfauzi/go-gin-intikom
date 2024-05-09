package middleware

import (
	"asidikfauzi/go-gin-intikom/common/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type JwtClaim struct {
	Email string
	jwt.RegisteredClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		authHeader := c.GetHeader("Authorization")

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.Unauthorized}, startTime)
			return
		}
		tokenString := parts[1]
		claims := &JwtClaim{}

		jwtKey := []byte(helper.GetEnv("JWT_KEY"))

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.UnauthorizedInvalid}, startTime)
				return
			case jwt.ValidationErrorExpired:
				helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.UnauthorizedExpired}, startTime)
				return
			default:
				helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.Unauthorized}, startTime)
				return
			}
		}
		if !token.Valid {
			helper.ResponseAPI(c, false, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), map[string]interface{}{helper.Error: helper.UnauthorizedExpired}, startTime)
			return
		}

		c.Next()
	}
}
