package repository

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user vkfilms.User) (int, error)
}

type Actor interface {
}

type Film interface {
}

type Repository struct {
	Authorization
	Actor
	Film
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
