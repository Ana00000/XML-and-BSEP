package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsBlockedProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsBlockedProfilesRepository) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	result := repo.Database.Create(profileSettingsBlockedProfiles)
	fmt.Print(result)
	return nil
}
