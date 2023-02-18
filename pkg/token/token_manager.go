package token

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	NewJWT(userID int32) (string, error)
	Parse(accessToken string) (*int64, error)
	NewRefreshToken() (string, error)
}

type tokenManager struct {
	jwtSecret string
}

func NewTokenManager(jwtSecret string) TokenManager {
	return &tokenManager{jwtSecret: jwtSecret}
}

func (t *tokenManager) NewJWT(userID int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject: string(userID),
	})

	return token.SignedString([]byte(t.jwtSecret))
}

func (t *tokenManager) Parse(accessToken string) (*int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(t.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("cannot get claims from token")
	}
	atoi, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("cannot convert str to int: %v", err)
	}
	id := int64(atoi)
	return &id, nil
}

func (t *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
