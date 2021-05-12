package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsBlockedProfilesService struct {
	Repo * repository.ProfileSettingsBlockedProfilesRepository
}

func (service * ProfileSettingsBlockedProfilesService) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	err := service.Repo.CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles)
	if err != nil {
		return err
	}
	return nil
}
