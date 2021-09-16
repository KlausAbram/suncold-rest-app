package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
)

const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	AgentId int `json:"agent_id"`
}

type AuthCase struct {
	storage storage.Authorisation
}

func NewAuthCase(storage *storage.Authorisation) *AuthCase {
	return &AuthCase{storage: *storage}
}

func (ac *AuthCase) CreateAgent(agent models.Agent) (int, error) {
	agent.Password = HashPassword(agent.Password)
	return ac.storage.CreateAgent(agent)

}

func (ac *AuthCase) CreateJWT(username, password string) (string, error) {
	id, err := ac.storage.GetAgent(username, HashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func (ac *AuthCase) ParseJWT(tokenJWT string) (int, error) {
	tokenFromJWT, err := jwt.ParseWithClaims(tokenJWT, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := tokenFromJWT.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims*")
	}

	return claims.AgentId, nil
}

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("HASH_SALT"))))
}
