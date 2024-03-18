package repository

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	type args struct {
		item vkfilms.User
	}
	type mockBehavior func(args args, id int)

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
				item: vkfilms.User{
					Name:     "Test Name",
					Password: "Test Pass",
					Role:     vkfilms.ADMIN,
				},
			},
			want: 1,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(args.item.Name, args.item.Password, args.item.Role).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.want)

			got, err := r.CreateUser(testCase.input.item)
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

func TestAuthPostgres_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	type args struct {
		item vkfilms.User
	}
	type mockBehavior func(args args, userRole vkfilms.UserRole)

	var tests = []struct {
		name    string
		mock    mockBehavior
		input   args
		want    vkfilms.UserRole
		wantErr bool
	}{
		{
			name: "ok",
			input: args{
				item: vkfilms.User{
					Name:     "Test Name",
					Password: "Test Pass",
				},
			},
			want: vkfilms.ADMIN,
			mock: func(args args, userRole vkfilms.UserRole) {
				rows := sqlmock.NewRows([]string{"user_role"}).AddRow(userRole)
				mock.ExpectQuery("SELECT (.+) FROM users (.+)").
					WithArgs(args.item.Name, args.item.Password).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.want)

			got, err := r.GetUser(testCase.input.item.Name, testCase.input.item.Password)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got.Role)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
