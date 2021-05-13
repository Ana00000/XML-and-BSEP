package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
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

func (service * ProfileSettingsService) FindAllProfileSettings() []model.ProfileSettings{
	profileSettings := service.Repo.FindAllProfileSettings()
	if profileSettings != nil {
		return profileSettings
	}
	return nil
}

func (service * ProfileSettingsService) FindAllProfileSettingsForPublicUsers() []uuid.UUID{
	profileSettings := service.Repo.FindAllProfileSettingsForPublicUsers()
	if profileSettings != nil {
		return profileSettings
	}
	return nil
}