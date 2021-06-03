package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type TokenManager interface {
	NewJWT(userId string) (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty secret key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}