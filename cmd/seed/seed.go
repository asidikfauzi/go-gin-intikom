package main

import (
	"asidikfauzi/go-gin-intikom/common/database"
	"asidikfauzi/go-gin-intikom/model"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	InitMoviesSeed(db)

	fmt.Println("Successfully Seeder")
}

func InitMoviesSeed(db *gorm.DB) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password123!"), 10)
	if err != nil {
		panic("Error hash password: " + err.Error())
	}

	user := model.Users{
		Name:      "Ach Sidik Fauzi",
		Email:     "asidikfauzi@gmail.com",
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
	}

	var existingUser model.Users
	if err = db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic("Error while querying user: " + err.Error())
		}

		if err = db.Create(&user).Error; err != nil {
			panic("Error creating user: " + err.Error())
		}
	}

}
