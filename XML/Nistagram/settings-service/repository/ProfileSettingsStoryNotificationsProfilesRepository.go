package repository

import (
	"fmt"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ProfileSettingsStoryNotificationsProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsStoryNotificationsProfilesRepository) CreateProfileSettingsStoryNotificationsProfiles(profileSettingsStoryNotificationsProfiles *model.ProfileSettingsStoryNotificationsProfiles) error {
	result := repo.Database.Create(profileSettingsStoryNotificationsProfiles)
	fmt.Print(result)
	return nil
}

func (repo *ProfileSettingsStoryNotificationsProfilesRepository) FindAllProfileSettings() []model.ProfileSettingsStoryNotificationsProfiles{
	var profileSettings []model.ProfileSettingsStoryNotificationsProfiles
	repo.Database.Select("*").Find(&profileSettings)
	return profileSettings
}

func (repo *ProfileSettingsStoryNotificationsProfilesRepository) FindAllProfileSettingsForStoryNotifications(storyNotificationUserId uuid.UUID) []model.ProfileSettingsStoryNotificationsProfiles {

	var allStoryNotifications = repo.FindAllProfileSettings()
	var profileSettingsList []model.ProfileSettingsStoryNotificationsProfiles
	for i := 0; i < len(allStoryNotifications); i++ {
		if allStoryNotifications[i].StoryNotificationsProfileId == storyNotificationUserId {
			profileSettingsList = append(profileSettingsList, allStoryNotifications[i])
		}
	}
	return profileSettingsList
}
