package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsRejectedMessageProfilesService struct {
	Repo * repository.ProfileSettingsRejectedMessageProfilesRepository
}

func (service * ProfileSettingsRejectedMessageProfilesService) CreateProfileSettingsRejectedMessageProfiles(profileSettingsRejectedMessageProfiles *model.ProfileSettingsRejectedMessageProfiles) error {
	err := service.Repo.CreateProfileSettingsRejectedMessageProfiles(profileSettingsRejectedMessageProfiles)
	if err != nil {
		return err
	}
	return nil
}
