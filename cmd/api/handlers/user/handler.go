package user

import portsid "github.com/ejquintans/go-l/internal/ports"

type UserHandler struct {
	UserService portsid.UserService
}
