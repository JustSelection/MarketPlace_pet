package models

type User struct {
	ID          string `json:"user_id"`
	Name        string `json:"name,omitempty"`
	Information string `json:"information,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
}
