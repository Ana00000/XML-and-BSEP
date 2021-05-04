package dto

import "github.com/google/uuid"

type RegisteredUserFollowersDTO struct {
	RegisteredUserId uuid.UUID `json:"registered_user_id"`
	FollowerUserId uuid.UUID `json:"follower_user_id"`
}
