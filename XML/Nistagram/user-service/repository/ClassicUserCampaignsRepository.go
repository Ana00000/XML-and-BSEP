package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"gorm.io/gorm"
)

type ClassicUserCampaignsRepository struct {
	Database * gorm.DB
}

func (repo * ClassicUserCampaignsRepository) CreateClassicUserCampaigns(classicUserCampaigns *model.ClassicUserCampaigns) error {
	result := repo.Database.Create(classicUserCampaigns)
	fmt.Print(result)
	return nil
}
