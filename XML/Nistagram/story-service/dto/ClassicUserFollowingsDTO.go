package dto

import (
	"github.com/google/uuid"
)

type ClassicUserFollowingsDTO struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	FollowingUserId uuid.UUID `json:"following_user_id"`
}
