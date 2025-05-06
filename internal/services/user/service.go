package user

import (
	portsid "github.com/ejquintans/go-l/internal/ports"
)

type Service struct {
	Repo portsid.UserRepository
}
