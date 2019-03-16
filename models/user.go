package models

type User struct {
	Id              uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Fullname        string `json:"fullname"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	Picture         string `json:"picture"`
	Comments []Comment `json:"comments,omitempty"`
}
