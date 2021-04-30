package model

type PostICR struct {
	InappropriateContentRequest
	PostId string `json: "postId" gorm:"not null"`
}