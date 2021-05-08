package service

import (
	"../model"
	"../repository"
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
