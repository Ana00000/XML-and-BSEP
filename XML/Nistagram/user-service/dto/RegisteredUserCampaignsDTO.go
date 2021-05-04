package dto

import "github.com/google/uuid"

type RegisteredUserCampaignsDTO struct {
	RegisteredUserId uuid.UUID `json:"registered_user_id"`
	CampaignId uuid.UUID `json:"campaign_id"`
}
