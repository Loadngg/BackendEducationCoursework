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

type Quizzes interface {
	GetAll() ([]models.Quiz, error)
	GetById(id string) (*models.Quiz, error)
	GetAllQuestions(id string) ([]models.QuizQuestions, error)
	GetQuestionById(id string) (*models.QuizQuestions, error)
	GetAllAnswers(questionId string) ([]models.Answers, error)
	CreateAnswer(input models.UserAnswers) (int32, error)
}

type Results interface {
	GetAll() ([]models.UserQuiz, error)
}

type Service struct {
	Authorization
	Lectures
	Quizzes
	Results
}

func New(db *database.Database, salt string, signingKey string) *Service {
	return &Service{
		Authorization: NewAuthService(db.Authorization, salt, signingKey),
		Lectures:      NewLecturesService(db.Lectures),
		Quizzes:       NewQuizzesService(db.Quizzes),
		Results:       NewResultsService(db.Results),
	}
}
