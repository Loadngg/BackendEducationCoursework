package service

import (
	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

type Authorization interface {
	CreateUser(user models.Users) (int32, error)
	GenerateToken(login, password string) (string, error)
}

type Lectures interface {
}

type Quiz interface {
}

type Service struct {
	Authorization
	Lectures
	Quiz
}

func New(db *database.Database, salt string, signingKey string) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization, salt, signingKey),
	}
}
