package dto

import "github.com/google/uuid"

type FollowRequestForUserDTO struct {
	ID uuid.UUID `json:"id"`
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	FollowerUserId uuid.UUID `json:"follower_user_id"`
	FollowRequestStatus string `json:"follow_request_status"`
}
