package base

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtPayload struct {
	ID   string
	Role string
}

func GenerateAuthToken(id string, jwtKey string, role string) (tokenString string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": id,
			"role":    role,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err = token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("generating JWT key: %w", err)
	}

	return tokenString, nil
}

func VerifyAuthToken(tokenString, jwtKey string) (*JwtPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	payload := new(JwtPayload)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		payload.ID = fmt.Sprint(claims["user_id"])
		payload.Role = fmt.Sprint(claims["role"])
	}

	return payload, nil
}
