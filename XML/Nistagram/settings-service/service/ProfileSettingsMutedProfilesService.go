package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsMutedProfilesService struct {
	Repo * repository.ProfileSettingsMutedProfilesRepository
}

func (service * ProfileSettingsMutedProfilesService) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	err := service.Repo.CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles)
	if err != nil {
		return err
	}
	return nil
}
