package user

import (
	"context"
	"fmt"
	"github.com/ejquintans/go-l/internal/domain"
	"log"
)

func (r Repository) Insert(user domain.User) (id interface{}, err error) {

	//--------------------------------------------------------
	query := `INSERT INTO users (name, surname, password, email, role, image, created_data) 
		          VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var userId int64
	err = r.Client.QueryRow(context.Background(), query, user.Name, user.Surname, user.Password, user.Email, user.Role, user.Image, user.Created_Data).Scan(&userId)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return nil, fmt.Errorf("Error al insertar usuario: %w", err)
	}

	return userId, nil
}
