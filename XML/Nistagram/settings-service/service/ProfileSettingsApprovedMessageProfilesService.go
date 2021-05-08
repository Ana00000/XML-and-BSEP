package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsApprovedMessageProfilesService struct {
	Repo * repository.ProfileSettingsApprovedMessageProfilesRepository
}

func (service * ProfileSettingsApprovedMessageProfilesService) CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles *model.ProfileSettingsApprovedMessageProfiles) error {
	err := service.Repo.CreateProfileSettingsApprovedMessageProfiles(profileSettingsApprovedMessageProfiles)
	if err != nil {
		return err
	}
	return nil
}
