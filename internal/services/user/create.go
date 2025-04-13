package user

import (
	"fmt"
	"github.com/ejquintans/go-l/internal/domain"
	"log"
	"time"
)

func (s Service) Create(user domain.User) (id interface{}, err error) {
	// Carga el CreationTime
	user.Created_Data = time.Now().UTC()

	// Guardar
	userId, nil := s.Repo.Insert(user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("Error al crear usuario: %w", err)
	}

	// Responder con el id del recurso creado
	return userId, nil
}
