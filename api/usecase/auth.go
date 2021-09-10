package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/klaus-abram/suncold-restful-app/api/storage"
	"github.com/klaus-abram/suncold-restful-app/models"
	"os"
	"time"
)

const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	AgentId int `json:"agent_id"`
}

type AuthCase struct {
	storage storage.Authorisation
}

func NewAuthCase(storage storage.Authorisation) *AuthCase {
	return &AuthCase{storage: storage}
}

func (ac *AuthCase) CreateAgent(agent *models.Agent) (int, error) {
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

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("HASH_SALT"))))
}
