package dto

import "github.com/google/uuid"

type ClassicUserFollowersDTO struct {
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	FollowerUserId uuid.UUID `json:"follower_user_id"`
}
