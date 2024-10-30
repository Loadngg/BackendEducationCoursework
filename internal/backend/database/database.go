package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type Authorization interface {
	CreateUser(user models.Users) (int32, error)
	GetUser(login, password string) (models.Users, error)
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

type Database struct {
	Authorization
	Lectures
	Quiz
	Results
}

func New(db *reform.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
		Lectures:      NewLecturesPostgres(db),
	}
}
