package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type AuthPostgres struct {
	db *reform.DB
}

func NewAuthPostgres(db *reform.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.Users) (int32, error) {
	if err := a.db.Insert(&user); err != nil {
		return 0, err
	}

	return user.ID, nil
}
