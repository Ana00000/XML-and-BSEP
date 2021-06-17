package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsMutedProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsMutedProfilesRepository) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	result := repo.Database.Create(profileSettingsMutedProfiles)
	fmt.Print(result)
	return nil
}

func (repo *ProfileSettingsMutedProfilesRepository) FindAllMutedUserForLoggedUser(id uuid.UUID) []uuid.UUID {
	var listRetValuesUserIds []uuid.UUID
	var profileSettingsMutedProfiles []model.ProfileSettingsMutedProfiles
	repo.Database.Select("*").Find(&profileSettingsMutedProfiles)

	for i := 0; i < len(profileSettingsMutedProfiles); i++ {
		if profileSettingsMutedProfiles[i].ProfileSettingsId==id{
			listRetValuesUserIds = append(listRetValuesUserIds, profileSettingsMutedProfiles[i].MutedProfileId)
		}
	}
	return listRetValuesUserIds
}
