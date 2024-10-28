package service

import "backend/internal/backend/database"

type Service struct {
}

func New(repository *database.Database) *Service {
	return &Service{}
}
