package models

//go:generate reform

// UserQuiz represents a row in user_quiz table.
//
//reform:user_quiz
type UserQuiz struct {
	ID     int32  `reform:"id,pk" json:"id"`
	UserID *int32 `reform:"user_id" json:"user_id"`
	QuizID *int32 `reform:"quiz_id" json:"quiz_id"`
	Score  int32  `reform:"score" json:"score"`
}
