package dto

type CommentICRDTO struct {
	Note string `json:"note"`
	UserId string `json:"userId"`
	CommentId string `json:"commentId"`
}