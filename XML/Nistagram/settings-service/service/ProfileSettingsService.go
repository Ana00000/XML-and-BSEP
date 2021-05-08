package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsService struct {
	Repo * repository.ProfileSettingsRepository
}

func (service * ProfileSettingsService) CreateProfileSettings(profileSettings *model.ProfileSettings) error {
	err := service.Repo.CreateProfileSettings(profileSettings)
	if err != nil {
		return err
	}
	return nil
}