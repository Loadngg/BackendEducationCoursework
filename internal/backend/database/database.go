package database

import "gopkg.in/reform.v1"

type Authorization interface {
}

type Lectures interface {
}

type Quiz interface {
}

type Database struct {
	Authorization
	Lectures
	Quiz
}

func New(db *reform.DB) *Database {
	return &Database{}
}
