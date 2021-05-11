package repository

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"fmt"
	"gorm.io/gorm"
)

type ProfileSettingsRepository struct {
	Database * gorm.DB
}

func (repo * ProfileSettingsRepository) CreateProfileSettings(profileSettings *model.ProfileSettings) error {
	result := repo.Database.Create(profileSettings)
	fmt.Print(result)
	return nil
}

