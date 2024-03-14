package repository

import (
	"fmt"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user vkfilms.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, user_role) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
