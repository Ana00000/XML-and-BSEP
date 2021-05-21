package dto

import "github.com/google/uuid"

type ClassicUserCloseFriendsDTO struct {
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	CloseFriendUserId uuid.UUID `json:"close_friend_user_id"`
}
