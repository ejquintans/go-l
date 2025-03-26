package main

import (
	"github.com/ejquintans/go-l/cmd/api/handlers/user"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ginEngine := gin.Default()

	ginEngine.POST("/users", user.CreatePlayer)

	log.Println(ginEngine.Run(":8001"))
}
