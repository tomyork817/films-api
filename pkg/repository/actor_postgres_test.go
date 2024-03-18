package repository

import (
	"database/sql"
	"errors"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"
)

func TestActorPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		item vkfilms.Actor
	}
	type mockBehavior func(args args, id int)

	parse, _ := time.Parse(time.DateOnly, "2000-01-01")
	var tests = []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "ok",
			input: args{
				item: vkfilms.Actor{
					Name:     "Test Name",
					Sex:      vkfilms.MALE,
					Birthday: vkfilms.JsonDate(parse),
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO actors").
					WithArgs(args.item.Name, args.item.Sex, args.item.Birthday.Format()).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "empty fields",
			input: args{
				item: vkfilms.Actor{
					Name:     "",
					Sex:      vkfilms.MALE,
					Birthday: vkfilms.JsonDate(parse),
				},
			},
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("error"))
				mock.ExpectQuery("INSERT INTO actors").
					WithArgs(args.item.Name, args.item.Sex, args.item.Birthday.Format()).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.want)

			got, err := r.Create(testCase.input.item)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestActorPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		id int
	}

	var tests = []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "ok",
			input: args{
				id: 1,
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM actors WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "not found",
			input: args{
				id: 1,
			},
			mock: func() {
				mock.ExpectExec("DELETE FROM actors WHERE (.+)").
					WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			err := r.Delete(testCase.input.id)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestActorPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	type args struct {
		id    int
		input vkfilms.UpdateActorInput
	}

	name := "test"
	sex := vkfilms.FEMALE
	parse, _ := time.Parse(time.DateOnly, "2000-01-01")
	birthday := vkfilms.JsonDate(parse)
	var tests = []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "ok",
			input: args{
				id: 1,
				input: vkfilms.UpdateActorInput{
					Name:     &name,
					Sex:      &sex,
					Birthday: &birthday,
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec("UPDATE actors SET (.+) WHERE (.+)").
					WithArgs(name, sex, birthday.Format(), 1).WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "no args",
			input: args{
				id: 1,
				input: vkfilms.UpdateActorInput{
					Name:     nil,
					Sex:      nil,
					Birthday: nil,
					FilmsId:  nil,
				},
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectCommit()
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			err := r.Update(testCase.input.id, testCase.input.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestActorPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewActorPostgres(db)

	firstDate, _ := time.Parse(time.DateOnly, "2000-01-01")
	secondDate, _ := time.Parse(time.DateOnly, "2010-01-01")
	var tests = []struct {
		name    string
		mock    func()
		want    []vkfilms.GetActorOutput
		wantErr bool
	}{
		{
			name: "ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "birthday", "sex"}).
					AddRow(1, "name1", firstDate, vkfilms.MALE).
					AddRow(2, "name2", secondDate, vkfilms.FEMALE)

				mock.ExpectQuery("SELECT (.+) FROM actors").
					WithArgs().WillReturnRows(rows)
				rows = sqlmock.NewRows([]string{})
				mock.ExpectQuery("SELECT (.+) FROM actors a JOIN films_actors fa (.+) JOIN films f (.+)").
					WithArgs().WillReturnRows(rows)
				mock.ExpectQuery("SELECT (.+) FROM actors a JOIN films_actors fa (.+) JOIN films f (.+)").
					WithArgs().WillReturnRows(rows)
			},
			want: []vkfilms.GetActorOutput{
				{1, "name1", vkfilms.JsonDate(firstDate), vkfilms.MALE, nil},
				{2, "name2", vkfilms.JsonDate(secondDate), vkfilms.FEMALE, nil},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetAll()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
