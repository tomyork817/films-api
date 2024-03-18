package service

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user vkfilms.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (vkfilms.UserRole, error)
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

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Actor:         NewActorService(repos.Actor),
		Film:          NewFilmService(repos.Film),
	}
}
