package model

import "time"

// User model
type (
	Users struct {
		ID        uint       `gorm:"primaryKey"`
		Name      string     `gorm:"type:varchar(255)"`
		Email     string     `gorm:"type:varchar(255);unique"`
		Password  string     `gorm:"type:varchar(255)"`
		CreatedAt time.Time  `gorm:"autoCreateTime"`
		UpdatedAt *time.Time `gorm:"default:null"`
		DeletedAt *time.Time `gorm:"default:null"`
		Tasks     []Tasks    `gorm:"foreignKey:UserID"`
	}

	UserDataGoogle struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}

	ReqLogin struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	ReqUser struct {
		Name            string `json:"name" binding:"required,max=255"`
		Email           string `json:"email" binding:"required,email,max=255"`
		Password        string `json:"password" binding:"required,min=8,max=255"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
		OldPassword     string `json:"old_password,omitempty"`
	}

	GetUser struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
