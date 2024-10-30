package models

//go:generate reform

// Answers represents a row in answers table.
//
//reform:answers
type Answers struct {
	ID         int32  `reform:"id,pk"`
	Text       string `reform:"text"`
	QuestionID *int32 `reform:"question_id"`
}
