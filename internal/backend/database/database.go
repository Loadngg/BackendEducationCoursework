package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type Authorization interface {
	CreateUser(user models.Users) (int32, error)
}

type Lectures interface {
}

type Quiz interface {
}

type Database struct {
	Authorization
	Lectures
	Quiz
}

func New(db *reform.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
	}
}
