package dto

import "github.com/google/uuid"

type ClassicUserFollowingsDTO struct {
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	FollowingUserId uuid.UUID `json:"following_user_id"`
}

