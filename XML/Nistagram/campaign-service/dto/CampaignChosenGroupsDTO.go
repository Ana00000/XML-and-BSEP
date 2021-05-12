package dto

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/google/uuid"
)

type CampaignChosenGroupDTO struct {
	CampaignId uuid.UUID `json:"campaign_id"`
	RegisteredUserCategory model.RegisteredUserCategory `json:"registered_user_category"`
}