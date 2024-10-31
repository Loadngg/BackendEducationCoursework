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

type Database struct {
	Authorization
	Lectures
	Quizzes
	Results
}

func New(db *reform.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
		Lectures:      NewLecturesPostgres(db),
		Quizzes:       NewQuizzesPostgres(db),
		Results:       NewResultsPostgres(db),
	}
}
