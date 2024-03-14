package vk_films

type UserRole string

const (
	USER  UserRole = "user"
	ADMIN UserRole = "admin"
)

type User struct {
	Id       int      `json:"-"`
	Name     string   `json:"username"`
	Role     UserRole `json:"role"`
	Password string   `json:"password"`
}
