package vk_films

type Sort struct {
	Type  SortType
	Order SortOrder
}

type SortType string

const (
	NAME   SortType = "name"
	RATING SortType = "rating"
	DATE   SortType = "date"
)

type SortOrder string

const (
	DESC SortOrder = "desc"
	ASC  SortOrder = "asc"
)

type Search struct {
	Type     SearchType
	Fragment string
}

type SearchType string

const (
	FILM  SearchType = "film"
	ACTOR SearchType = "actor"
)
