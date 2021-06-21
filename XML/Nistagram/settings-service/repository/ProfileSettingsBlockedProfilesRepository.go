package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsBlockedProfilesRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsBlockedProfilesRepository) CreateProfileSettingsBlockedProfiles(profileSettingsBlockedProfiles *model.ProfileSettingsBlockedProfiles) error {
	result := repo.Database.Create(profileSettingsBlockedProfiles)
	fmt.Print(result)
	return nil
}

func (repo *ProfileSettingsBlockedProfilesRepository) FindAllBlockedAndBlockingUsersForLoggedUser(id uuid.UUID, userId uuid.UUID) ([]uuid.UUID, []uuid.UUID){
	var listRetValuesUserIds []uuid.UUID
	var listRetValuesProfileSettingsIds []uuid.UUID
	var profileSettingsBlockedProfiles []model.ProfileSettingsBlockedProfiles
	repo.Database.Select("*").Find(&profileSettingsBlockedProfiles)

	for i := 0; i < len(profileSettingsBlockedProfiles); i++ {
		if profileSettingsBlockedProfiles[i].ProfileSettingsId==id{
			listRetValuesUserIds = append(listRetValuesUserIds, profileSettingsBlockedProfiles[i].BlockedProfileId)
		}
		if profileSettingsBlockedProfiles[i].BlockedProfileId == userId {
			listRetValuesProfileSettingsIds = append(listRetValuesProfileSettingsIds,profileSettingsBlockedProfiles[i].ProfileSettingsId)
		}
	}
	return listRetValuesUserIds, listRetValuesProfileSettingsIds
}

func (repo *ProfileSettingsBlockedProfilesRepository) CheckIfBlocked(profileSettingsID uuid.UUID, blockedUserID uuid.UUID) bool {
	profileSettingsBlockedProfiles := &model.ProfileSettingsBlockedProfiles{}
	if repo.Database.First(&profileSettingsBlockedProfiles, "profile_settings_id = ? and blocked_profile_id = ?", profileSettingsID,blockedUserID).RowsAffected == 0 {
		return false
	}
	return true
}

func (repo *ProfileSettingsBlockedProfilesRepository) FindProfileSettingsBlockedProfiles(profileSettingsID uuid.UUID, blockedUserID uuid.UUID) *model.ProfileSettingsBlockedProfiles {
	profileSettingsBlockedProfiles := &model.ProfileSettingsBlockedProfiles{}
	if repo.Database.First(&profileSettingsBlockedProfiles, "profile_settings_id = ? and blocked_profile_id = ?",profileSettingsID,blockedUserID).RowsAffected == 0 {
		return nil
	}
	return profileSettingsBlockedProfiles
}

func (repo *ProfileSettingsBlockedProfilesRepository) UnblockUser(id uuid.UUID) {
	repo.Database.Delete(&model.ProfileSettingsBlockedProfiles{}, id)
}