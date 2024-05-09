package model

import "time"

// Task model
type Task struct {
	ID          uint       `gorm:"primaryKey"`
	UserID      uint       // Foreign key to users table
	Title       string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(50);default:pending"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"default:null"`
	DeletedAt   *time.Time `gorm:"default:null"`
}
