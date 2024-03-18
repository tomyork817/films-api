package repository

import (
	"fmt"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type ActorPostgres struct {
	db *sqlx.DB
}

func NewActorPostgres(db *sqlx.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (r *ActorPostgres) Create(actor vkfilms.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, sex, birthday) values ($1, $2, $3) RETURNING id", actorsTable)

	row := r.db.QueryRow(query, actor.Name, actor.Sex, actor.Birthday.Format())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ActorPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", actorsTable)

	_, err := r.db.Exec(query, id)

	return err
}

func (r *ActorPostgres) Update(id int, input vkfilms.UpdateActorInput) error {
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
	if input.Sex != nil {
		setValues = append(setValues, fmt.Sprintf("sex=$%d", argId))
		args = append(args, *input.Sex)
		argId++
	}
	if input.Birthday != nil {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, input.Birthday.Format())
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", actorsTable, setQuery, argId)
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
	if input.FilmsId != nil {
		for _, filmId := range input.FilmsId {
			_, err = tx.Exec(createFilmsActorsQuery, filmId, id)
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

func (r *ActorPostgres) GetAll() ([]vkfilms.GetActorOutput, error) {
	var actors []vkfilms.GetActorOutput

	selectActorsQuery := fmt.Sprintf("SELECT * FROM %s", actorsTable)
	if err := r.db.Select(&actors, selectActorsQuery); err != nil {
		return nil, err
	}

	selectFilmsQuery := fmt.Sprintf(`SELECT  f.id, f.name, f.description, f.rating, f.date FROM %s a 
    JOIN %s fa ON a.id = fa.actor_id JOIN %s f ON fa.film_id = f.id WHERE a.id = $1;`, actorsTable, filmsActorsTable, filmsTable)
	for i, actor := range actors {
		var films []vkfilms.Film
		if err := r.db.Select(&films, selectFilmsQuery, actor.Id); err != nil {
			return nil, err
		}
		actors[i].Films = films
	}

	return actors, nil
}
