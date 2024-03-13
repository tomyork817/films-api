package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
