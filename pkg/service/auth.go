package service

import (
	"crypto/sha1"
	"fmt"
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/repository"
)

const salt = "auwnfhxtwl193620k27n1ms"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user vkfilms.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	hash.Sum([]byte(salt))

	return fmt.Sprintf("%x", hash)
}
