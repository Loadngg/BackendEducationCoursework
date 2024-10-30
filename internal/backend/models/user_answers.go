package models

//go:generate reform

// UserAnswers represents a row in user_answers table.
//
//reform:user_answers
type UserAnswers struct {
	ID         int32  `reform:"id,pk" json:"id"`
	UserID     *int32 `reform:"user_id" json:"user_id"`
	QuestionID *int32 `reform:"question_id" json:"question_id"`
	AnswerID   *int32 `reform:"answer_id" json:"answer_id"`
	IsCorrect  bool   `reform:"is_correct" json:"is_correct"`
}
