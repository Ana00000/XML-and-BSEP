package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Database * gorm.DB
}

func (repo * CampaignRepository) CreateCampaign(campaign *model.Campaign) error {
	result := repo.Database.Create(campaign)
	fmt.Print(result)
	return nil
}