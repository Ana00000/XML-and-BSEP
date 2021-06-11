package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsMutedProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsMutedProfilesRepository) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	result := repo.Database.Create(profileSettingsMutedProfiles)
	fmt.Print(result)
	return nil
}
