package dto

import "github.com/google/uuid"

type ClassicUserCampaignsDTO struct {
	ClassicUserId uuid.UUID `json:"classic_user_id"`
	CampaignId uuid.UUID `json:"campaign_id"`
}
