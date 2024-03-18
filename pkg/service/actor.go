package service

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/repository"
)

type ActorService struct {
	repo repository.Actor
}

func NewActorService(repo repository.Actor) *ActorService {
	return &ActorService{repo: repo}
}

func (s *ActorService) Create(actor vkfilms.Actor) (int, error) {
	if err := actor.Validate(); err != nil {
		return 0, err
	}
	return s.repo.Create(actor)
}

func (s *ActorService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ActorService) Update(id int, input vkfilms.UpdateActorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}

func (s *ActorService) GetAll() ([]vkfilms.GetActorOutput, error) {
	return s.repo.GetAll()
}
