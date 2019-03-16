package models

type Comment struct {
	Id       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	ParentId uint      `json:"parent_id"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  int8      `json:"has_vote"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
