package vk_films

import "errors"

type UserRole string

const (
	USER  UserRole = "user"
	ADMIN UserRole = "admin"
)

type User struct {
	Id       int      `json:"-" db:"id"`
	Name     string   `json:"username"`
	Role     UserRole `json:"role" db:"user_role"`
	Password string   `json:"password"`
}

func (u User) ValidateSignIn() error {
	if u.Name == "" || u.Password == "" || (u.Role != ADMIN && u.Role != USER) {
		return errors.New("not all required fields are filled in")
	}
	return nil
}

func (u User) ValidateSignUp() error {
	if u.Name == "" || u.Password == "" {
		return errors.New("not all required fields are filled in")
	}
	return nil
}

type SignInUserInput struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
