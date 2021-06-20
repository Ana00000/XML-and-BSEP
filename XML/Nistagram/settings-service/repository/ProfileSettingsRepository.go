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
	ProfileSettingsPostNotificationsProfilesRepository *ProfileSettingsPostNotificationsProfilesRepository
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
	switch profileSettings.MessageApprovalType {
	case "PUBLIC":
		messageApprovalType = model.PUBLIC
	case "FRIENDS_ONLY":
		messageApprovalType = model.FRIENDS_ONLY
	}

	likesNotifications := model.ALL_NOTIFICATIONS
	switch profileSettings.LikesNotifications {
	case "ALL_NOTIFICATIONS":
		likesNotifications = model.ALL_NOTIFICATIONS
	case "FRIENDS_NOTIFICATIONS":
		likesNotifications = model.FRIENDS_NOTIFICATIONS
	case "NONE":
		likesNotifications = model.NONE
	}

	commentsNotifications := model.ALL_NOTIFICATIONS
	switch profileSettings.CommentsNotifications {
	case "ALL_NOTIFICATIONS":
		commentsNotifications = model.ALL_NOTIFICATIONS
	case "FRIENDS_NOTIFICATIONS":
		commentsNotifications = model.FRIENDS_NOTIFICATIONS
	case "NONE":
		commentsNotifications = model.NONE
	}

	messagesNotifications := model.ALL_NOTIFICATIONS
	switch profileSettings.MessagesNotifications {
	case "ALL_NOTIFICATIONS":
		messagesNotifications = model.ALL_NOTIFICATIONS
	case "FRIENDS_NOTIFICATIONS":
		messagesNotifications = model.FRIENDS_NOTIFICATIONS
	case "NONE":
		messagesNotifications = model.NONE
	}


	result := repo.Database.Model(&model.ProfileSettings{}).Where("user_id = ?", profileSettings.UserId)

	result.Update("user_visibility", userVisibility)
	result.Update("message_approval_type", messageApprovalType)
	result.Update("is_post_taggable", profileSettings.IsPostTaggable)
	result.Update("is_story_taggable", profileSettings.IsStoryTaggable)
	result.Update("is_comment_taggable", profileSettings.IsCommentTaggable)
	result.Update("likes_notifications", likesNotifications)
	result.Update("comments_notifications", commentsNotifications)
	result.Update("messages_notifications", messagesNotifications)
	fmt.Println(result.RowsAffected)

	fmt.Println("updating profile settings")
	return nil
}

func (repo *ProfileSettingsRepository) FindAllUserIdsFromProfileSettings(profileSettingsList []model.ProfileSettingsPostNotificationsProfiles) []uuid.UUID {

	var allProfileSetting = repo.FindAllProfileSettings()
	var userIdsList []uuid.UUID
	for i := 0; i < len(allProfileSetting); i++ {
		for j:=0; j<len(profileSettingsList); j++{
			if allProfileSetting[i].ID == profileSettingsList[j].ProfileSettingsId{
				userIdsList = append(userIdsList, allProfileSetting[i].UserId)
			}
		}
	}
	return userIdsList
}

func (repo *ProfileSettingsRepository) FindAllUsersForPostNotifications(id uuid.UUID) []uuid.UUID {
	var profileSettings []model.ProfileSettingsPostNotificationsProfiles
	profileSettings = repo.ProfileSettingsPostNotificationsProfilesRepository.FindAllProfileSettingsForPostNotifications(id)
	return repo.FindAllUserIdsFromProfileSettings(profileSettings)
}

func (repo *ProfileSettingsRepository) FindAllUsersForPostAlbumNotifications(id uuid.UUID) []uuid.UUID {
	var profileSettings []model.ProfileSettingsPostNotificationsProfiles
	profileSettings = repo.ProfileSettingsPostNotificationsProfilesRepository.FindAllProfileSettingsForPostNotifications(id)
	return repo.FindAllUserIdsFromProfileSettings(profileSettings)
}