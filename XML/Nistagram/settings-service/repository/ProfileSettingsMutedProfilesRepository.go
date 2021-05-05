package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type ProfileSettingsMutedProfilesRepository struct {
	Database * gorm.DB
}

func (repo * ProfileSettingsMutedProfilesRepository) CreateProfileSettingsMutedProfiles(profileSettingsMutedProfiles *model.ProfileSettingsMutedProfiles) error {
	result := repo.Database.Create(profileSettingsMutedProfiles)
	fmt.Print(result)
	return nil
}
