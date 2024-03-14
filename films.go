package vk_films

import "time"

type Film struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Rating      float32   `json:"rating"`
}

type FilmsActors struct {
	Id      int `json:"id"`
	FilmId  int `json:"filmId"`
	ActorId int `json:"actorId"`
}

type Sex string

const (
	FEMALE Sex = "female"
	MALE   Sex = "male"
)

type Actor struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Sex      Sex       `json:"sex"`
}
