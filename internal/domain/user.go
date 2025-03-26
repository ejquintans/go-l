package domain

import "time"

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
