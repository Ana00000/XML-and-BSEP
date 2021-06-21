package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsMutedProfilesService struct {
	Repo *repository.ProfileSettingsMutedProfilesRepository
}

func (service *ProfileSettingsMutedProfilesService) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	err := service.Repo.CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileSettingsMutedProfilesService) FindAllMutedUserForLoggedUser(id uuid.UUID) []uuid.UUID {
	return service.Repo.FindAllMutedUserForLoggedUser(id)
}

func (service *ProfileSettingsMutedProfilesService) CheckIfMuted(profileSettingsID uuid.UUID, mutedUserID uuid.UUID) bool {
	return service.Repo.CheckIfMuted(profileSettingsID,mutedUserID)
}

func (service *ProfileSettingsMutedProfilesService) FindProfileSettingsMutedProfiles(profileSettingsID uuid.UUID, mutedUserID uuid.UUID) *model.ProfileSettingsMutedProfiles {
	return service.Repo.FindProfileSettingsMutedProfiles(profileSettingsID,mutedUserID)
}

func (service *ProfileSettingsMutedProfilesService) UnmuteUser(id uuid.UUID) {
	service.Repo.UnmuteUser(id)
}
