package models

type Agent struct {
	Id        int    `json:"-"`
	Name      string `json:"name" binding:"required"`
	AgentName string `json:"agent_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
