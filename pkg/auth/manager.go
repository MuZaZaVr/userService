package auth

import (
	"fmt"
	"github.com/MuZaZaVr/notesService/pkg/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

type TokenManager interface {
	NewJWT(userId string) (string, error)
	UserIdentity(handler http.Handler) http.Handler
	Parse(accessToken string) (string, error)
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

func (m *Manager) UserIdentity(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)

		if header == "" {
			middleware.JSONError(w, http.StatusUnauthorized ,errors.New("invalid auth header"))
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			middleware.JSONError(w, http.StatusUnauthorized, errors.New("invalid auth header"))
			return
		}

		_, err := m.Parse(headerParts[1])
		if err != nil {
			middleware.JSONError(w, http.StatusUnauthorized, err)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	subClaims, ok := token.Claims.(jwt.MapClaims)["sub"]
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return subClaims.(string), nil
}

