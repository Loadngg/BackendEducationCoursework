package service

import (
	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

type Authorization interface {
	CreateUser(user models.Users) (int32, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int32, error)
}

type Lectures interface {
	GetAll() ([]models.Lectures, error)
	GetById(id string) (*models.Lectures, error)
	GetAllMaterials(lectureId string) ([]models.LectureMaterials, error)
}

type Quiz interface {
}

type Results interface {
}

type Service struct {
	Authorization
	Lectures
	Quiz
	Results
}

func New(db *database.Database, salt string, signingKey string) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization, salt, signingKey),
		Lectures:      NewLecturesService(db.Lectures),
	}
}
