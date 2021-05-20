package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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

func (service *ProfileSettingsService) FindProfileSettingByUserId(id uuid.UUID) *model.ProfileSettings {
	user := service.Repo.FindProfileSettingByUserId(id)
	return user
}

func (service * ProfileSettingsService) FindAllPublicUsers(allValidUsers []userModel.ClassicUser) []userModel.ClassicUser{
	publicUsers := service.Repo.FindAllPublicUsers(allValidUsers)
	if publicUsers != nil {
		return publicUsers
	}
	return nil
}
