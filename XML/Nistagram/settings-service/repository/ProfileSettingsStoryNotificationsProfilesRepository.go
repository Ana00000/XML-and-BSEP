package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsStoryNotificationsProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsStoryNotificationsProfilesRepository) CreateProfileSettingsStoryNotificationsProfiles(profileSettingsStoryNotificationsProfiles *model.ProfileSettingsStoryNotificationsProfiles) error {
	result := repo.Database.Create(profileSettingsStoryNotificationsProfiles)
	fmt.Print(result)
	return nil
}

