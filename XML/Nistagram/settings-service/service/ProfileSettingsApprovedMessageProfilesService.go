package service

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsApprovedMessageProfilesService struct {
	Repo *repository.ProfileSettingsApprovedMessageProfilesRepository
}

func (service *ProfileSettingsApprovedMessageProfilesService) CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles *model.ProfileSettingsApprovedMessageProfiles) error {
	err := service.Repo.CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles)
	if err != nil {
		return err
	}
	return nil
}
