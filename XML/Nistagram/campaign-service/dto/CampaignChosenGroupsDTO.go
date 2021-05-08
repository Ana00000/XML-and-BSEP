package dto

import (
	"../model"
	"github.com/google/uuid"
)

type CampaignChosenGroupDTO struct {
	CampaignId uuid.UUID `json:"campaign_id"`
	UserCategoryValue model.UserCategory `json:"user_category_value"`
}