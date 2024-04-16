package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tolazhewa/todo-backend-go/api"
	"github.com/tolazhewa/todo-backend-go/config"
	"github.com/tolazhewa/todo-backend-go/db"
)

// @title           Todo Backend
// @version         1.0
// @description     Backend for a todo application written in golang

// @license.name  MIT
// @license.url   https://github.com/tolazhewa/todo-backend-go/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	config.InitAppConfig()
	db.InitDB()
	api.Run()
}
