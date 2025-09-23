package auth

import (
	"time"
	"umkm-api/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID uint, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.JWTExpireHour)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func (s *jwtService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
}
