package service

import "github.com/bitbox228/vk-films-api/pkg/repository"

type Authorization interface {
}

type Actor interface {
}

type Film interface {
}

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
