package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"fmt"
	"gorm.io/gorm"
)

type ProfileSettingsBlockedProfilesRepository struct {
	Database * gorm.DB
}

func (repo * ProfileSettingsBlockedProfilesRepository) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	result := repo.Database.Create(profileSettingsBlockedProfiles)
	fmt.Print(result)
	return nil
}
