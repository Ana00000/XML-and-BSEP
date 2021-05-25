package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
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