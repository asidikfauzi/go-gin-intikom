package helper

import (
	"asidikfauzi/go-gin-intikom/common/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func ValidateTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Field '" + fe.Field() + "' is required"
	case "max":
		return fe.Field() + " must be less than or equal to " + fe.Param()
	case "min":
		return fe.Field() + " must be greater than or equal to " + fe.Param()
	case "email":
		return "Format email not valid"
	}

	return "Unknown error"
}

func ValidatePassword(password string) (err error) {
	var hasUpper, hasLower, hasNumber bool

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		}
	}

	if !hasUpper {
		err = fmt.Errorf("Password must contain at least one uppercase letter")
	} else if !hasLower {
		err = fmt.Errorf("Password must contain at least one lowercase letter")
	} else if !hasNumber {
		err = fmt.Errorf("Password must contain at least one digit")
	}

	return
}

func ErrorPassword(c *gin.Context, field, message string, statusCode int) {
	errorMessage := map[string]string{field: message}
	ResponseAPI(c, false, statusCode, http.StatusText(statusCode), map[string]interface{}{Error: errorMessage}, time.Now())
}

func ErrorValidatingPassword(c *gin.Context, password, confirmPassword string) bool {
	err := ValidatePassword(password)
	if err != nil {
		log.Error(err)
		ErrorPassword(c, "password", err.Error(), http.StatusUnprocessableEntity)
		return false
	}

	if password != confirmPassword {
		log.Error(err)
		ErrorPassword(c, "password", "Password does not match", http.StatusUnprocessableEntity)
		return false
	}

	return true
}
