package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsStoryNotificationsProfilesService struct {
	Repo *repository.ProfileSettingsStoryNotificationsProfilesRepository
}

func (service *ProfileSettingsStoryNotificationsProfilesService) CreateProfileSettingsStoryNotificationsProfiles(profileSettingsStoryNotificationsProfiles *model.ProfileSettingsStoryNotificationsProfiles) error {
	err := service.Repo.CreateProfileSettingsStoryNotificationsProfiles(profileSettingsStoryNotificationsProfiles)
	if err != nil {
		return err
	}
	return nil
}


