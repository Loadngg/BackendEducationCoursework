package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type ResultsPostgres struct {
	db *reform.DB
}

func NewResultsPostgres(db *reform.DB) *ResultsPostgres {
	return &ResultsPostgres{db: db}
}

func (r *ResultsPostgres) GetAll() ([]models.UserQuiz, error) {
	return nil, nil
}
