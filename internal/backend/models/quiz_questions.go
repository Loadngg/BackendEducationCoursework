package models

//go:generate reform

// QuizQuestions represents a row in quiz_questions table.
//
//reform:quiz_questions
type QuizQuestions struct {
	ID              int32  `reform:"id,pk" json:"id"`
	Question        string `reform:"question" json:"question"`
	CorrectAnswerID *int32 `reform:"correct_answer_id" json:"correct_answer_id"`
	QuizID          *int32 `reform:"quiz_id" json:"quiz_id"`
}
