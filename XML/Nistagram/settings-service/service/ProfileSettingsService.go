package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsService struct {
	Repo *repository.ProfileSettingsRepository
}

func (service *ProfileSettingsService) CreateProfileSettings(profileSettings *model.ProfileSettings) error {
	err := service.Repo.CreateProfileSettings(profileSettings)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileSettingsService) FindAllProfileSettings() []model.ProfileSettings {
	profileSettings := service.Repo.FindAllProfileSettings()
	if profileSettings != nil {
		return profileSettings
	}
	return nil
}

func (service *ProfileSettingsService) FindAllProfileSettingsForPublicUsers() []uuid.UUID {
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

func (service *ProfileSettingsService) FindAllPublicUsers(allValidUsers []dto.ClassicUserDTO) []dto.ClassicUserDTO {
	publicUsers := service.Repo.FindAllPublicUsers(allValidUsers)
	if publicUsers != nil {
		return publicUsers
	}
	return nil
}

func (service *ProfileSettingsService) UpdateProfileSettings(profileSettings *dto.ProfileSettingsDTO) error {
	err := service.Repo.UpdateProfileSettings(profileSettings)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileSettingsService) FindAllUsersForPostNotifications(id uuid.UUID)[]uuid.UUID {
	users := service.Repo.FindAllUsersForPostNotifications(id)
	if users != nil {
		return users
	}
	return nil
}

func (service *ProfileSettingsService) FindAllUsersForPostAlbumNotifications(id uuid.UUID)[]uuid.UUID {
	users := service.Repo.FindAllUsersForPostAlbumNotifications(id)
	if users != nil {
		return users
	}
	return nil
}


func (service *ProfileSettingsService) FindAllUsersForStoryNotifications(id uuid.UUID)[]uuid.UUID {
	users := service.Repo.FindAllUsersForStoryNotifications(id)
	if users != nil {
		return users
	}
	return nil
}

