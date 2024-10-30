package service

import "backend/internal/backend/database"

type Authorization interface {
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

func New(repository *database.Database) *Service {
	return &Service{}
}
