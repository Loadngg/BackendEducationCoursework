package service

import (
	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

type QuizzesService struct {
	db database.Quizzes
}

func NewQuizzesService(db database.Quizzes) *QuizzesService {
	return &QuizzesService{db: db}
}

func (s *QuizzesService) GetAll() ([]models.Quiz, error) {
	return s.db.GetAll()
}

func (s *QuizzesService) GetById(id string) (*models.Quiz, error) {
	return s.db.GetById(id)
}

func (s *QuizzesService) GetAllQuestions(id string) ([]models.QuizQuestions, error) {
	return s.db.GetAllQuestions(id)
}

func (s *QuizzesService) GetQuestionById(id string) (*models.QuizQuestions, error) {
	return s.db.GetQuestionById(id)
}

func (s *QuizzesService) GetAllAnswers(questionId string) ([]models.Answers, error) {
	return s.db.GetAllAnswers(questionId)
}

func (s *QuizzesService) CreateAnswer(input models.UserAnswers) (int32, error) {
	return s.db.CreateAnswer(input)
}
