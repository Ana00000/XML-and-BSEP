package dto

import "github.com/google/uuid"

type FollowRequestDTO struct {
	//status does not to be added trough the DTO because it will be set to PENDING
	ClassicUserId  uuid.UUID `json:"classic_user_id" validate:"required"`
	FollowerUserId uuid.UUID `json:"follower_user_id" validate:"required"`
}
