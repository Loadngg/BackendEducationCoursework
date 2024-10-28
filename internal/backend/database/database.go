package database

import "gopkg.in/reform.v1"

type Database struct {
}

func New(db *reform.DB) *Database {
	return &Database{}
}
