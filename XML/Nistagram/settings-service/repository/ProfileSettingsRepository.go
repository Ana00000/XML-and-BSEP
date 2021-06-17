package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"gorm.io/gorm"
)

type ProfileSettingsRepository struct {
	Database *gorm.DB
}

func (repo *ProfileSettingsRepository) CreateProfileSettings(profileSettings *model.ProfileSettings) error {
	result := repo.Database.Create(profileSettings)
	fmt.Print(result)
	return nil
}

func (repo *ProfileSettingsRepository) FindAllProfileSettings() []model.ProfileSettings {
	var profileSettings []model.ProfileSettings
	repo.Database.Select("*").Find(&profileSettings)
	return profileSettings

}

func (repo *ProfileSettingsRepository) FindAllProfileSettingsForPublicUsers() []uuid.UUID {
	var profileSettings = repo.FindAllProfileSettings()
	var listPublicUsers []uuid.UUID

	for i := 0; i < len(profileSettings); i++ {
		if profileSettings[i].UserVisibility == model.PUBLIC_VISIBILITY {
			listPublicUsers = append(listPublicUsers, profileSettings[i].UserId)
		}
	}

	return listPublicUsers

}

func (repo *ProfileSettingsRepository) FindProfileSettingByUserId(userId uuid.UUID) *model.ProfileSettings {

	profileSetting := &model.ProfileSettings{}

	if repo.Database.First(&profileSetting, "user_id = ?", userId).RowsAffected == 0 {
		return nil
	}

	return profileSetting
}

func (repo *ProfileSettingsRepository) FindAllPublicUsers(allValidUsers []dto.ClassicUserDTO) []dto.ClassicUserDTO {

	var publicProfileSettings = repo.FindAllProfileSettingsForPublicUsers()
	var listOfPublicUsers []dto.ClassicUserDTO

	for i := 0; i < len(allValidUsers); i++ {
		for j := 0; j < len(publicProfileSettings); j++ {
			if publicProfileSettings[j] == allValidUsers[i].ID {
				listOfPublicUsers = append(listOfPublicUsers, allValidUsers[i])
			}
		}
	}

	return listOfPublicUsers
}

func (repo *ProfileSettingsRepository) UpdateProfileSettings(profileSettings *dto.ProfileSettingsDTO) error {

	userVisibility := model.PUBLIC_VISIBILITY
	switch profileSettings.UserVisibility {
	case "PUBLIC_VISIBILITY":
		userVisibility = model.PUBLIC_VISIBILITY
	case "PRIVATE_VISIBILITY":
		userVisibility = model.PRIVATE_VISIBILITY
	}

	messageApprovalType := model.PUBLIC
	switch profileSettings.UserVisibility {
	case "PUBLIC":
		messageApprovalType = model.PUBLIC
	case "FRIENDS_ONLY":
		messageApprovalType = model.FRIENDS_ONLY
	}

	result := repo.Database.Model(&model.ProfileSettings{}).Where("user_id = ?", profileSettings.UserId)

	result.Update("user_visibility", userVisibility)
	fmt.Println(result.RowsAffected)
	result.Update("message_approval_type", messageApprovalType)
	fmt.Println(result.RowsAffected)
	result.Update("is_post_taggable", profileSettings.IsPostTaggable)
	fmt.Println(result.RowsAffected)
	result.Update("is_story_taggable", profileSettings.IsStoryTaggable)
	fmt.Println(result.RowsAffected)
	result.Update("is_comment_taggable", profileSettings.IsCommentTaggable)
	fmt.Println(result.RowsAffected)

	fmt.Println("updating profile settings")
	return nil
}