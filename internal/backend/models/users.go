package models

//go:generate reform

// Users represents a row in users table.
//
//reform:users
type Users struct {
	ID       int32  `reform:"id,pk" json:"id"`
	Login    string `reform:"login" json:"login" binding:"required"`
	Password string `reform:"password" json:"password" binding:"required"`
}
