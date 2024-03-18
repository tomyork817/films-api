package vk_films

import (
	"errors"
	"strings"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

type Film struct {
	Id          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Date        JsonDate `json:"date" db:"date"`
	Rating      float32  `json:"rating" db:"rating"`
}

type CreateFilmInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Date        JsonDate `json:"date"`
	Rating      float32  `json:"rating"`
	ActorsId    []int    `json:"actorsId"`
}

func (i CreateFilmInput) Validate() error {
	if i.Name == "" || i.Date.IsEmpty() || i.Rating == 0 {
		return errors.New("not all required fields are filled in")
	}
	if i.Rating > 10 || i.Rating < 0 || len(i.Description) > 1000 || len(i.Name) > 50 {
		return errors.New("invalid values")
	}
	return nil
}

type UpdateFilmInput struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Date        *JsonDate `json:"date"`
	Rating      *float32  `json:"rating"`
	ActorsId    []int     `json:"actorsId"`
}

func (i UpdateFilmInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Date == nil && i.Rating == nil && i.ActorsId == nil {
		return errors.New("no values to update")
	}
	if i.Rating != nil && (*i.Rating > 10 || *i.Rating < 0) {
		return errors.New("invalid rating value")
	}
	if i.Description != nil && (len(*i.Description) > 1000) {
		return errors.New("invalid rating value")
	}
	if i.Name != nil && (len(*i.Name) < 1 || len(*i.Name) > 50) {
		return errors.New("invalid rating value")
	}
	return nil
}

type FilmsActors struct {
	Id      int `json:"id"`
	FilmId  int `json:"filmId"`
	ActorId int `json:"actorId"`
}

type JsonDate time.Time

type Sex string

const (
	FEMALE Sex = "female"
	MALE   Sex = "male"
)

type Actor struct {
	Id       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Birthday JsonDate `json:"birthday" db:"birthday"`
	Sex      Sex      `json:"sex" db:"sex"`
}

func (a Actor) Validate() error {
	if a.Name == "" || a.Sex == "" {
		return errors.New("not all required fields are filled in")
	}
	if a.Sex != MALE && a.Sex != FEMALE || len(a.Name) > 50 {
		return errors.New("invalid values")
	}
	return nil
}

type UpdateActorInput struct {
	Name     *string   `json:"name"`
	Birthday *JsonDate `json:"birthday"`
	Sex      *Sex      `json:"sex"`
	FilmsId  []int     `json:"filmsId"`
}

func (a UpdateActorInput) Validate() error {
	if a.Name == nil && a.Sex == nil && a.Birthday == nil && a.FilmsId == nil {
		return errors.New("no values to update")
	}
	if *a.Sex != MALE && *a.Sex != FEMALE || len(*a.Name) > 50 {
		return errors.New("invalid values")
	}
	return nil
}

type GetActorOutput struct {
	Id       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	Birthday JsonDate `json:"birthday" db:"birthday"`
	Sex      Sex      `json:"sex" db:"sex"`
	Films    []Film   `json:"films"`
}

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j *JsonDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(*j).Format(dateLayout) + "\""), nil
}

func (j *JsonDate) Format() string {
	t := time.Time(*j)
	return t.Format(dateLayout)
}

func (j *JsonDate) IsEmpty() bool {
	return time.Time(*j).IsZero()
}
