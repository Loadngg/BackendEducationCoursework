package models

//go:generate reform

// Quiz represents a row in quiz table.
//
//reform:quiz
type Quiz struct {
	ID    int32  `reform:"id,pk" json:"id"`
	Title string `reform:"title" json:"title"`
}
