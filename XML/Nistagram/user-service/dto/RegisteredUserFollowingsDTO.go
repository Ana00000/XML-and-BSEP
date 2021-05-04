package dto

import "github.com/google/uuid"

type RegisteredUserFollowingsDTO struct {
	RegisteredUserId uuid.UUID `json:"registered_user_id"`
	FollowingUserId uuid.UUID `json:"following_user_id"`
}
