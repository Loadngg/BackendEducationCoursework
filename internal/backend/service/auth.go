package service

import (
	"crypto/sha1"
	"fmt"

	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

var passwordSalt string

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization, salt string) *AuthService {
	passwordSalt = salt
	return &AuthService{db: db}
}

func (s *AuthService) CreateUser(user models.Users) (int32, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.db.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
