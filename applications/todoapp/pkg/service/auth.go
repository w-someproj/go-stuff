package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/repository"
	"time"
)

const (
	salt       = "jbhyntgmf"
	signingKey = "kgbmhfnmhokjkgtmsvfdg" // random bytes for signing token (also used for decoding tokens)
	tokenTTL   = 7 * 24 * time.Hour      // 7 days
)

type AuthService struct {
	repos repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

// send user structure to lower layer - repository
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	// get user from db
	user, err := s.repos.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// claims - json object with different fields
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing Method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims don`t have right type")
	}

	return claims.UserId, nil
}

// can`t save passwords open in db
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	// good practice - add to password set of random symbols before hashing
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
