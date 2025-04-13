package user

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	Client *pgxpool.Pool
}
