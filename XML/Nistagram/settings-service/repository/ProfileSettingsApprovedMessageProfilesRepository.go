package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"fmt"
	"gorm.io/gorm"
)

type ProfileSettingsApprovedMessageProfilesRepository struct {
	Database * gorm.DB
}

func (repo * ProfileSettingsApprovedMessageProfilesRepository) CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles *model.ProfileSettingsApprovedMessageProfiles) error {
	result := repo.Database.Create(profileSettingsApprovedMessageProfiles)
	fmt.Print(result)
	return nil
}
