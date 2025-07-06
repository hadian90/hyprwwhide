package models

type Window struct {
	Address   string    `json:"address"`
	Title     string    `json:"title"`
	Class     string    `json:"class"`
	Workspace Workspace `json:"workspace"`
}
