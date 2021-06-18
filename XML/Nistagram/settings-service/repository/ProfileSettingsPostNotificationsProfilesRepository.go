package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsPostNotificationsProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsPostNotificationsProfilesRepository) CreateProfileSettingsPostNotificationsProfiles(profileSettingsPostNotificationsProfiles *model.ProfileSettingsPostNotificationsProfiles) error {
	result := repo.Database.Create(profileSettingsPostNotificationsProfiles)
	fmt.Print(result)
	return nil
}

