package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type CampaignChosenGroupRepository struct {
	Database * gorm.DB
}

func (repo * CampaignChosenGroupRepository) CreateCampaignChosenGroup(campaignChosenGroup *model.CampaignChosenGroup) error {
	result := repo.Database.Create(campaignChosenGroup)
	fmt.Print(result)
	return nil
}
