package user

import (
	"github.com/ejquintans/go-l/internal/ports"
)

type Service struct {
	Repo ports.UserRepository
}
