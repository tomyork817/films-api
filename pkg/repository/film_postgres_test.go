package repository

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"
)

func TestFilmPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilmPostgres(db)

	type args struct {
		item vkfilms.CreateFilmInput
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
				item: vkfilms.CreateFilmInput{
					Name:        "Test Name",
					Description: "Test Description",
					Rating:      2.5,
					Date:        vkfilms.JsonDate(parse),
					ActorsId:    nil,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO films").
					WithArgs(args.item.Name, args.item.Description, args.item.Date.Format(), args.item.Rating).WillReturnRows(rows)
				mock.ExpectCommit()
			},
			wantErr: false,
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

func TestFilmPostgres_GetAllSorted(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilmPostgres(db)

	sort := vkfilms.Sort{}
	firstDate, _ := time.Parse(time.DateOnly, "2000-01-01")
	secondDate, _ := time.Parse(time.DateOnly, "2010-01-01")
	var tests = []struct {
		name    string
		mock    func()
		want    []vkfilms.Film
		wantErr bool
	}{
		{
			name: "ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "date", "rating"}).
					AddRow(1, "name1", "d1", firstDate, 2.5).
					AddRow(2, "name2", "d2", secondDate, 5)

				mock.ExpectQuery("SELECT (.+) FROM films (.+)").
					WithArgs().WillReturnRows(rows)
			},
			want: []vkfilms.Film{
				{1, "name1", "d1", vkfilms.JsonDate(firstDate), 2.5},
				{2, "name2", "d2", vkfilms.JsonDate(secondDate), 5},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetAllSorted(sort)
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

func TestFilmPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilmPostgres(db)

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
				mock.ExpectExec("DELETE FROM films WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
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

func TestFilmPostgres_GetSearch(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilmPostgres(db)

	search := vkfilms.Search{
		Type: vkfilms.FILM,
	}
	firstDate, _ := time.Parse(time.DateOnly, "2000-01-01")
	secondDate, _ := time.Parse(time.DateOnly, "2010-01-01")
	var tests = []struct {
		name    string
		mock    func()
		want    []vkfilms.Film
		wantErr bool
	}{
		{
			name: "ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "date", "rating"}).
					AddRow(1, "name1", "d1", firstDate, 2.5).
					AddRow(2, "name2", "d2", secondDate, 5)

				mock.ExpectQuery("SELECT (.+) FROM films (.+)").
					WithArgs().WillReturnRows(rows)
			},
			want: []vkfilms.Film{
				{1, "name1", "d1", vkfilms.JsonDate(firstDate), 2.5},
				{2, "name2", "d2", vkfilms.JsonDate(secondDate), 5},
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetSearch(search)
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
