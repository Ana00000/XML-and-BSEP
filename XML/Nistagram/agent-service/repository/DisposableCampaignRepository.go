package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type DisposableCampaignRepository struct {
	Database * gorm.DB
}

func (repo * DisposableCampaignRepository) CreateDisposableCampaign(disposableCampaign *model.DisposableCampaign) error {
	result := repo.Database.Create(disposableCampaign)
	fmt.Print(result)
	return nil
}