package service

import (
	"../model"
	"../repository"
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
