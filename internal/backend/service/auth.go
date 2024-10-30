package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"backend/internal/backend/database"
	"backend/internal/backend/models"
	"github.com/dgrijalva/jwt-go"
)

var passwordSalt string
var signingTokenKey string

const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int32 `json:"user_id"`
}

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization, salt string, signingKey string) *AuthService {
	passwordSalt = salt
	signingTokenKey = signingKey
	return &AuthService{db: db}
}

func (s *AuthService) CreateUser(user models.Users) (int32, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.db.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.db.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingTokenKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}

func (s *AuthService) ParseToken(tokenStr string) (int32, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingTokenKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
