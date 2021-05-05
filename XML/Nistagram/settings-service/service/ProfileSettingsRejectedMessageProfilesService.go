package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsRejectedMessageProfilesService struct {
	Repo * repository.ProfileSettingsRejectedMessageProfilesRepository
}

func (service * ProfileSettingsRejectedMessageProfilesService) CreateProfileSettingsRejectedMessageProfiles(profileSettingsRejectedMessageProfiles *model.ProfileSettingsRejectedMessageProfiles) error {
	service.Repo.CreateProfileSettingsRejectedMessageProfiles(profileSettingsRejectedMessageProfiles)
	return nil
}
