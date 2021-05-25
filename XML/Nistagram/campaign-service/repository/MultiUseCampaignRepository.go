package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"fmt"
	"gorm.io/gorm"
)

type MultiUseCampaignRepository struct {
	Database * gorm.DB
}

func (repo * MultiUseCampaignRepository) CreateMultiUseCampaign(multiUseCampaign *model.MultiUseCampaign) error {
	result := repo.Database.Create(multiUseCampaign)
	fmt.Print(result)
	return nil
}