package models

// UsersAccount
type Accounts struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	EMail  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
