package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

func ConnectCLient(dbURI string) (client *pgxpool.Pool, err error) {
	// Establece un tiempo de espera para permitir que el proceso de conexión se cancele si tarda demasiado
	// Crea un contexto(ctx) que muere a los 10 seg
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		log.Println("Error al parsear la configuración de la base de datos:", err)
		return nil, err
	}

	client, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return nil, err
	}

	// Verificar conexión
	err = client.Ping(ctx)
	if err != nil {
		log.Println("No se pudo conectar a PostgreSQL:", err)
		client.Close()
		return nil, err
	}

	log.Println("Conexión exitosa a PostgreSQL")
	return client, nil
}

func GetDBURI() string {
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Println("DATABASE_URL no está definida. Verificá la configuración.")
	}
	fmt.Println(uri)
	return uri
}
