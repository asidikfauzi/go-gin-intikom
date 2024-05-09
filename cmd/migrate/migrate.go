package main

import (
	"asidikfauzi/go-gin-intikom/common/database"
	"asidikfauzi/go-gin-intikom/model"
	"fmt"
)

func main() {
	_, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = database.DB.AutoMigrate(&model.Users{})
	if err != nil {
		panic("Error Migrate User")
	}

	err = database.DB.AutoMigrate(&model.Tasks{})
	if err != nil {
		panic("Error Migrate Tas")
	}

	fmt.Println("Successfully Migrate")
}
