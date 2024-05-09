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

	ReqRegister struct {
		Name            string    `json:"name" binding:"required"`
		Email           string    `json:"email" binding:"required,email"`
		Password        string    `json:"password" binding:"required,min=8"`
		ConfirmPassword string    `json:"confirm_password" binding:"required"`
		CreatedAt       time.Time `json:"created_at"`
	}

	GetUser struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
