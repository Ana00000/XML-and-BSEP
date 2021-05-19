package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
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

func (repo * ProfileSettingsRepository) FindAllProfileSettings() []model.ProfileSettings{
	var profileSettings []model.ProfileSettings
	repo.Database.Select("*").Find(&profileSettings)
	return profileSettings

}

func (repo * ProfileSettingsRepository) FindAllProfileSettingsForPublicUsers() []uuid.UUID{
	var profileSettings = repo.FindAllProfileSettings()
	var listPublicUsers []uuid.UUID

	for i := 0; i< len(profileSettings); i++{
		if profileSettings[i].UserVisibility == model.PRIVATE_VISIBILITY {
			listPublicUsers = append(listPublicUsers, profileSettings[i].UserId)
		}
	}

	return listPublicUsers

}

func (repo * ProfileSettingsRepository) FindProfileSettingByUserId(userId uuid.UUID) *model.ProfileSettings {

	profileSetting := &model.ProfileSettings{}

	if repo.Database.First(&profileSetting, "user_id = ?", userId).RowsAffected == 0{
		return nil
	}

	return profileSetting
}

func (repo * ProfileSettingsRepository) FindAllPublicUsers(allValidUsers []userModel.ClassicUser) []userModel.ClassicUser {

	var publicProfileSettings = repo.FindAllProfileSettingsForPublicUsers()
	var listOfPublicUsers []userModel.ClassicUser

	for i:=0; i<len(allValidUsers); i++{
		for j:=0; j<len(publicProfileSettings);j++{
			if publicProfileSettings[j] == allValidUsers[i].ID{
				listOfPublicUsers = append(listOfPublicUsers, allValidUsers[i])
			}
		}
	}

	return listOfPublicUsers
}



