package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (repo *ClassicUserCampaignsRepository) FindById(id string) *model.ClassicUserCampaigns {
	campaign := &model.ClassicUserCampaigns{}
	if repo.Database.First(&campaign, "id = ?", id).RowsAffected == 0{
		return nil
	}
	return campaign
}

func (repo * ClassicUserCampaignsRepository) FindAllCampaignsForUser(userId uuid.UUID) []model.ClassicUserCampaigns{
	var campaigns []model.ClassicUserCampaigns
	repo.Database.Select("id = ?", userId).Find(&campaigns)
	return campaigns
}

