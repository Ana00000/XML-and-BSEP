package service

import (
	"../model"
	"../repository"
)

type ProfileSettingsService struct {
	Repo * repository.ProfileSettingsRepository
}

func (service * ProfileSettingsService) CreateProfileSettings(profileSettings *model.ProfileSettings) error {
	service.Repo.CreateProfileSettings(profileSettings)
	return nil
}