package user

import "github.com/ejquintans/go-l/internal/ports"

type UserHandler struct {
	UserService ports.UserService
}
