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
		Tasks     []Task     `gorm:"foreignKey:UserID"`
	}

	ReqLogin struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	ReqUser struct {
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required,email"`
		Password        string `json:"password" binding:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
		OldPassword     string `json:"old_password,omitempty"`
	}

	GetUser struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
