package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsApprovedMessageProfilesService struct {
	Repo * repository.ProfileSettingsApprovedMessageProfilesRepository
}

func (service * ProfileSettingsApprovedMessageProfilesService) CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles *model.ProfileSettingsApprovedMessageProfiles) error {
	service.Repo.CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles)
	return nil
}
