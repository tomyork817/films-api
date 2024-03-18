package repository

import (
	"fmt"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type FilmPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *FilmPostgres {
	return &FilmPostgres{db: db}
}

func (r *FilmPostgres) Create(film vkfilms.CreateFilmInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var filmId int
	createFilmQuery := fmt.Sprintf("INSERT INTO %s (name, description, date, rating) VALUES ($1, $2, $3, $4) RETURNING id", filmsTable)

	row := tx.QueryRow(createFilmQuery, film.Name, film.Description, film.Date.Format(), film.Rating)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createFilmsActorsQuery := fmt.Sprintf("INSERT INTO %s (film_id, actor_id) VALUES ($1, $2)", filmsActorsTable)
	for _, actorId := range film.ActorsId {
		_, err = tx.Exec(createFilmsActorsQuery, filmId, actorId)
		if err != nil {
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return filmId, tx.Commit()
}

func (r *FilmPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", filmsTable)

	_, err := r.db.Exec(query, id)

	return err
}

func (r *FilmPostgres) Update(id int, input vkfilms.UpdateFilmInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, input.Date.Format())
		argId++
	}
	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, input.Date.Format())
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", filmsTable, setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", setQuery)
	logrus.Debugf("args: %s", args)

	if argId != 1 {
		_, err = r.db.Exec(query, args...)
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	createFilmsActorsQuery := fmt.Sprintf("INSERT INTO %s (film_id, actor_id) VALUES ($1, $2)", filmsActorsTable)

	if input.ActorsId != nil {
		for _, actorId := range input.ActorsId {
			_, err = tx.Exec(createFilmsActorsQuery, id, actorId)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *FilmPostgres) GetAllSorted(sort vkfilms.Sort) ([]vkfilms.Film, error) {
	var films []vkfilms.Film

	selectFilmsQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY %s %s", filmsTable, sort.Type, sort.Order)
	if err := r.db.Select(&films, selectFilmsQuery); err != nil {
		return nil, err
	}

	return films, nil
}

func (r *FilmPostgres) GetSearch(search vkfilms.Search) ([]vkfilms.Film, error) {
	var films []vkfilms.Film
	var selectFilmsQuery string

	switch search.Type {
	case vkfilms.FILM:
		selectFilmsQuery = fmt.Sprintf("SELECT * FROM %s WHERE name LIKE '%%%s%%'", filmsTable, search.Fragment)
	case vkfilms.ACTOR:
		selectFilmsQuery = fmt.Sprintf(`SELECT DISTINCT f.id, f.name, f.description, f.date, f.rating FROM %s f 
                                                              JOIN %s fa ON f.id = fa.film_id 
                                                                  JOIN %s a ON a.id = fa.actor_id 
                                                                                   WHERE a.name LIKE '%%%s%%'`,
			filmsTable,
			filmsActorsTable,
			actorsTable,
			search.Fragment)

	}

	if err := r.db.Select(&films, selectFilmsQuery); err != nil {
		return nil, err
	}

	return films, nil
}
