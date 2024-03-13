package vk_films

type UserRole string

const (
	USER  UserRole = "USER"
	ADMIN UserRole = "ADMIN"
)

type User struct {
	Id       int      `json:"-"`
	Name     string   `json:"username"`
	Role     UserRole `json:"role"`
	Password string   `json:"password"`
}
