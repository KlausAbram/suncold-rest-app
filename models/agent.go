package models

type Agent struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	AgentName string `json:"agent_name"`
	Password  string `json:"password"`
}
