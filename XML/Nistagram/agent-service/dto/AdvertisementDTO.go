package dto

import "github.com/google/uuid"

type AdvertisementDTO struct {
	AdvertisementContentId string `json:"advertisement_content_id"`
	CampaignRefer uuid.UUID `json:"campaign_refer"`
}
