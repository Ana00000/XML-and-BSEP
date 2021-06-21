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

func (repo *ProfileSettingsMutedProfilesRepository) CheckIfMuted(profileSettingsID uuid.UUID, mutedUserID uuid.UUID) bool {
	profileSettingsMutedProfiles := &model.ProfileSettingsMutedProfiles{}
	if repo.Database.First(&profileSettingsMutedProfiles, "profile_settings_id = ? and muted_profile_id = ?", profileSettingsID,mutedUserID).RowsAffected == 0 {
		return false
	}
	return true
}

func (repo *ProfileSettingsMutedProfilesRepository) FindProfileSettingsMutedProfiles(profileSettingsID uuid.UUID, mutedUserID uuid.UUID) *model.ProfileSettingsMutedProfiles {
	profileSettingsMutedProfiles := &model.ProfileSettingsMutedProfiles{}
	if repo.Database.First(&profileSettingsMutedProfiles, "profile_settings_id = ? and muted_profile_id = ?",profileSettingsID,mutedUserID).RowsAffected == 0 {
		return nil
	}
	return profileSettingsMutedProfiles
}

func (repo *ProfileSettingsMutedProfilesRepository) UnmuteUser(id uuid.UUID) {
	repo.Database.Delete(&model.ProfileSettingsMutedProfiles{}, id)
}
