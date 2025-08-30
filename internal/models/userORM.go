package models

type User struct {
	ID          string `json:"user_id" gorm:"primaryKey"`
	Name        string `json:"name,omitempty"`
	Information string `json:"information,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
}

type UserUpdate struct {
	Name        string `json:"name,omitempty"`
	Information string `json:"information,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}
