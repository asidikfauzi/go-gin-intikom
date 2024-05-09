package model

import "time"

// Task model
type (
	Tasks struct {
		ID          uint       `gorm:"primaryKey"`
		UserID      uint       // Foreign key to users table
		Title       string     `gorm:"type:varchar(255)"`
		Description string     `gorm:"type:text"`
		Status      string     `gorm:"type:varchar(50);default:pending"`
		CreatedAt   time.Time  `gorm:"autoCreateTime"`
		UpdatedAt   *time.Time `gorm:"default:null"`
		DeletedAt   *time.Time `gorm:"default:null"`
	}

	ReqTask struct {
		UserID      uint   `json:"user_id" binding:"required"`
		Title       string `json:"title" binding:"required,max=255"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" binding:"max=50"`
	}

	GetTask struct {
		ID          uint    `json:"id"`
		UserID      uint    `json:"user_id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Status      string  `json:"status"`
		User        GetUser `json:"user"`
	}
)
