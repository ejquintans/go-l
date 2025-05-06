package portsid

import "github.com/ejquintans/go-l/internal/domain"

type UserService interface {
	Create(user domain.User) (id interface{}, err error)
}

type UserRepository interface {
	Insert(user domain.User) (id interface{}, err error)
}
