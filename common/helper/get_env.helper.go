package helper

import (
	"fmt"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	var appConfig map[string]string
	var err error
	appConfig, err = godotenv.Read()

	if err != nil {
		fmt.Println("Error reading .env file")
	}

	value, ok := appConfig[key]
	if !ok {
		fmt.Printf("Environment variable %s not found\n", key)
		return ""
	}
	return value
}
