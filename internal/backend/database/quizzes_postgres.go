package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type QuizzesPostgres struct {
	db *reform.DB
}

func NewQuizzesPostgres(db *reform.DB) *QuizzesPostgres {
	return &QuizzesPostgres{db: db}
}

func (a *QuizzesPostgres) GetAll() ([]models.Quiz, error) {
	var quizzesList []models.Quiz

	quizzes, err := a.db.SelectAllFrom(models.QuizTable, "")
	if err != nil {
		return nil, err
	}

	for _, q := range quizzes {
		item := *q.(*models.Quiz)
		quizzesList = append(quizzesList, item)
	}

	return quizzesList, nil
}

func (a *QuizzesPostgres) GetById(id string) (*models.Quiz, error) {
	var quiz models.Quiz

	err := a.db.SelectOneTo(&quiz, "WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (a *QuizzesPostgres) GetAllQuestions(id string) ([]models.QuizQuestions, error) {
	var questionsList []models.QuizQuestions

	questions, err := a.db.SelectAllFrom(models.QuizQuestionsTable, "WHERE quiz_id=$1", id)
	if err != nil {
		return nil, err
	}

	for _, q := range questions {
		item := *q.(*models.QuizQuestions)
		questionsList = append(questionsList, item)
	}

	return questionsList, nil
}

func (a *QuizzesPostgres) GetQuestionById(id string) (*models.QuizQuestions, error) {
	var question models.QuizQuestions

	err := a.db.SelectOneTo(&question, "WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (a *QuizzesPostgres) GetAllAnswers(id string) ([]models.Answers, error) {
	var answersList []models.Answers

	answers, err := a.db.SelectAllFrom(models.AnswersTable, "WHERE question_id=$1", id)
	if err != nil {
		return nil, err
	}

	for _, a := range answers {
		item := *a.(*models.Answers)
		answersList = append(answersList, item)
	}

	return answersList, nil
}

func (a *QuizzesPostgres) CreateAnswer(input models.UserAnswers) (int32, error) {
	if err := a.db.Insert(&input); err != nil {
		return 0, err
	}

	return input.ID, nil
}
