package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsPostNotificationsProfilesService struct {
	Repo *repository.ProfileSettingsPostNotificationsProfilesRepository
}

func (service *ProfileSettingsPostNotificationsProfilesService) CreateProfileSettingsPostNotificationsProfiles(profileSettingsPostNotificationsProfiles *model.ProfileSettingsPostNotificationsProfiles) error {
	err := service.Repo.CreateProfileSettingsPostNotificationsProfiles(profileSettingsPostNotificationsProfiles)
	if err != nil {
		return err
	}
	return nil
}

