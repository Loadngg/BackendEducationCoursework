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

func (a *AuthPostgres) GetUser(login, password string) (models.Users, error) {
	var user models.Users
	err := a.db.SelectOneTo(&user, "WHERE login=$1 AND password=$2", login, password)
	return user, err
}
