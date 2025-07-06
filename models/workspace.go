package models

type Workspace struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Monitor string `json:"monitor"`
}
