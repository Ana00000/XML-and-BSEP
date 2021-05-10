package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
"fmt"
"gorm.io/gorm"
)

type RegisteredUserCampaignsRepository struct {
	Database * gorm.DB
}

func (repo * RegisteredUserCampaignsRepository) CreateRegisteredUserCampaigns(registeredUserCampaigns *model.RegisteredUserCampaigns) error {
	result := repo.Database.Create(registeredUserCampaigns)
	fmt.Print(result)
	return nil
}
