package main

import (
	"github.com/ejquintans/go-l/cmd/api/handlers/user"
	"github.com/ejquintans/go-l/internal/repositories/postgres"
	userPostgres "github.com/ejquintans/go-l/internal/repositories/postgres/user"
	userService "github.com/ejquintans/go-l/internal/services/user"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ginEngine := gin.Default()
	dbURI := postgres.GetDBURI()

	client, err := postgres.ConnectCLient(dbURI)
	if err != nil {
		log.Fatal(err.Error())

	}

	userRepo := userPostgres.Repository{
		Client: client,
	}

	userSrv := userService.Service{
		Repo: userRepo,
	}

	userHandler := user.UserHandler{
		UserService: userSrv,
	}

	ginEngine.POST("/users", userHandler.CreateUser)

	log.Println(ginEngine.Run(":8080"))
}
