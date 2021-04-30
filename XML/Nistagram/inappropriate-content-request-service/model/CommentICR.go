package model

type CommentICR struct {
	InappropriateContentRequest
	CommentId string `json: "commentId" gorm:"not null"`
}