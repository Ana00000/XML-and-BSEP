package service

import (
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
)

type ProfileSettingsBlockedProfilesService struct {
	Repo *repository.ProfileSettingsBlockedProfilesRepository
}

func (service *ProfileSettingsBlockedProfilesService) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	err := service.Repo.CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProfileSettingsBlockedProfilesService) FindAllBlockedAndBlockingUsersForLoggedUser(id uuid.UUID, userId uuid.UUID) ([]uuid.UUID,[]uuid.UUID) {
	return service.Repo.FindAllBlockedAndBlockingUsersForLoggedUser(id,userId)
}
