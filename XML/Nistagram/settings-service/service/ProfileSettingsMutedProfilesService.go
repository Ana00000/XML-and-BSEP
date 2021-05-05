package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsMutedProfilesService struct {
	Repo * repository.ProfileSettingsMutedProfilesRepository
}

func (service * ProfileSettingsMutedProfilesService) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	service.Repo.CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles)
	return nil
}
