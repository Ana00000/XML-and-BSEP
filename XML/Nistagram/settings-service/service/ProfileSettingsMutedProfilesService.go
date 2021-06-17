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
