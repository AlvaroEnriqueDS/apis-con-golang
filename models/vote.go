package models

type Vote struct {
	Id        uint `json:"id"`
	CommentID uint `json:"commentId"`
	UserID    uint `json:"userId"`
	Value     bool `json:"value"`
}
