package main

import (
	"fmt"
	"log"
	"os"

	web "github.com/antare74/gopek/router"
	"github.com/antare74/gopek/router/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	httpConnect()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func httpConnect() {
	router := gin.Default()
	baseURL := baseURL()
	api.Routes(router)
	web.Routes(router)
	router.Run(baseURL)
	fmt.Println("Connecting to http://localhost:8080")
	return
}

func baseURL() string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("host:", host+":"+port)
	return host + ":" + port
}
