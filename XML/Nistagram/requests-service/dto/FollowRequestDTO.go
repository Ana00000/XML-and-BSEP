package dto

import "github.com/google/uuid"

type FollowRequestDTO struct {
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	FollowerUserId uuid.UUID `json:"follower_user_id"`
	//status does not to be added trough the DTO because it will be set to PENDING
}
