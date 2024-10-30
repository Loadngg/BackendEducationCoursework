package models

//go:generate reform

// Lectures represents a row in lectures table.
//
//reform:lectures
type Lectures struct {
	ID    int32  `reform:"id,pk" json:"id"`
	Title string `reform:"title" json:"title"`
	Text  string `reform:"text" json:"text"`
}
