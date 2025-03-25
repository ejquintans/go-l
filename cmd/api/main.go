package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type User struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Image        string    `json:"image"`
	Created_Data time.Time `json:"-"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ginEngine := gin.Default()

	ginEngine.POST("/users", CreatePlayer)

	log.Println(ginEngine.Run(":8001"))
}

func CreatePlayer(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Carga el CreationTime
	user.Created_Data = time.Now().UTC()

	// Establece un tiempo de espera para permitir que el proceso de conexión se cancele si tarda demasiado
	// Crea un contexto(ctx) que muere a los 10 seg
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Se genera la URL para la conexion con la base
	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// Abrir conexion a la BD
	client, err := sql.Open("pgx", dns)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer client.Close()

	// Verificar conexión
	err = client.PingContext(ctx)
	if err != nil {
		fmt.Println("No se pudo conectar a PostgreSQL:", err)
		return
	}
	log.Println("Conexion exitosa")
	query := `INSERT INTO users (name, surname, password, email, role, image, created_data) 
		          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var userId int64
	err = client.QueryRowContext(ctx, query, user.Name, user.Surname, user.Password, user.Email, user.Role, user.Image, user.Created_Data).Scan(&userId)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		c.JSON(500, gin.H{"error": "Error al insertar usuario"})
		return
	}

	c.JSON(200, gin.H{"user_id": userId})
}
