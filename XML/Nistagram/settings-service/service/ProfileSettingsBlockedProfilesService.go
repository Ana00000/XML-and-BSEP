package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsBlockedProfilesService struct {
	Repo * repository.ProfileSettingsBlockedProfilesRepository
}

func (service * ProfileSettingsBlockedProfilesService) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	service.Repo.CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles)
	return nil
}
