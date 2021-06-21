package repository

import (
	"fmt"
	"github.com/google/uuid"
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

func (repo *ProfileSettingsPostNotificationsProfilesRepository) FindAllProfileSettings() []model.ProfileSettingsPostNotificationsProfiles{
	var profileSettings []model.ProfileSettingsPostNotificationsProfiles
	repo.Database.Select("*").Find(&profileSettings)
	return profileSettings

}

func (repo *ProfileSettingsPostNotificationsProfilesRepository) FindAllProfileSettingsForPostNotifications(postNotificationUserId uuid.UUID) []model.ProfileSettingsPostNotificationsProfiles {

	var allPostNotifications = repo.FindAllProfileSettings()
	var profileSettingsList []model.ProfileSettingsPostNotificationsProfiles
	for i := 0; i < len(allPostNotifications); i++ {
			if allPostNotifications[i].PostNotificationsProfileId == postNotificationUserId {
				profileSettingsList = append(profileSettingsList, allPostNotifications[i])
			}
	}
	return profileSettingsList
}

