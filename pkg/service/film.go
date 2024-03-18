package service

import (
	"errors"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) Create(film vkfilms.CreateFilmInput) (int, error) {
	if err := film.Validate(); err != nil {
		return 0, err
	}
	return s.repo.Create(film)
}

func (s *FilmService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *FilmService) Update(id int, input vkfilms.UpdateFilmInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}

func (s *FilmService) GetAllSorted(sort vkfilms.Sort) ([]vkfilms.Film, error) {
	if sort.Type != vkfilms.NAME && sort.Type != vkfilms.DATE && sort.Type != vkfilms.RATING {
		sort.Type = vkfilms.RATING
	}
	if sort.Order != vkfilms.ASC && sort.Order != vkfilms.DESC {
		sort.Order = vkfilms.DESC
	}
	return s.repo.GetAllSorted(sort)
}

func (s *FilmService) GetSearch(search vkfilms.Search) ([]vkfilms.Film, error) {
	if search.Type != vkfilms.ACTOR && search.Type != vkfilms.FILM {
		return nil, errors.New("no search type")
	}
	if search.Fragment == "" {
		return nil, errors.New("no fragment to search")
	}
	return s.repo.GetSearch(search)
}
