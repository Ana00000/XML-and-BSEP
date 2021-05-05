package repository

import (
	"../model"
	"fmt"
	"gorm.io/gorm"
)

type ProfileSettingsRejectedMessageProfilesRepository struct {
	Database * gorm.DB
}

func (repo * ProfileSettingsRejectedMessageProfilesRepository) CreateProfileSettingsRejectedMessageProfiles(profileSettingsRejectedMessageProfiles *model.ProfileSettingsRejectedMessageProfiles) error {
	result := repo.Database.Create(profileSettingsRejectedMessageProfiles)
	fmt.Print(result)
	return nil
}
