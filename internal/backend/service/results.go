package service

import (
	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

type ResultsService struct {
	db database.Results
}

func NewResultsService(db database.Results) *ResultsService {
	return &ResultsService{db: db}
}

func (s *ResultsService) GetAll() ([]models.UserQuiz, error) {
	return s.db.GetAll()
}
