package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
)

type AuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthStorage {
	return &AuthStorage{db: db}
}

func (as *AuthStorage) CreateAgent(agent models.Agent) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, agent_name, password_hash) values ($1, $2, $3) RETURNING id", agentsTable)
	row := as.db.QueryRow(query, agent.Name, agent.AgentName, agent.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (as *AuthStorage) GetAgent(agentname, password string) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE agent_name=$1 AND password_hash=$2", agentsTable)
	err := as.db.Get(&id, query, agentname, password)

	return id, err
}
