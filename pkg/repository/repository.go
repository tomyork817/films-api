package repository

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user vkfilms.User) (int, error)
	GetUser(username, password string) (vkfilms.User, error)
}

type Actor interface {
	Create(actor vkfilms.Actor) (int, error)
	Delete(id int) error
	Update(id int, input vkfilms.UpdateActorInput) error
	GetAll() ([]vkfilms.GetActorOutput, error)
}

type Film interface {
	Create(film vkfilms.CreateFilmInput) (int, error)
	Delete(id int) error
	Update(id int, input vkfilms.UpdateFilmInput) error
	GetAllSorted(sort vkfilms.Sort) ([]vkfilms.Film, error)
	GetSearch(search vkfilms.Search) ([]vkfilms.Film, error)
}

type Repository struct {
	Authorization
	Actor
	Film
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Actor:         NewActorPostgres(db),
		Film:          NewFilmPostgres(db),
	}
}
