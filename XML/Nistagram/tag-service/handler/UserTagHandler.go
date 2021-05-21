package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	profileSettingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"net/http"
)

type UserTagHandler struct {
	Service *service.UserTagService
	TagService *service.TagService
	ProfileSettingsService *profileSettingsService.ProfileSettingsService
}

func (handler *UserTagHandler) CreateUserTag(w http.ResponseWriter, r *http.Request) {
	var userTagDTO dto.UserTagDTO
	if err := json.NewDecoder(r.Body).Decode(&userTagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var findTag = handler.TagService.FindTagByName(userTagDTO.Name)
	var userTag model.UserTag

	if findTag != nil && userTagDTO.TagType == "USER_TAG" {
		w.WriteHeader(http.StatusExpectationFailed) // 417
		return
	} else {
		id := uuid.New()
		userTag = model.UserTag{
			Tag: model.Tag{
				ID: id,
				Name: userTagDTO.Name,
				TagType: model.USER_TAG,
			},
			UserId: userTagDTO.UserId,
		}

		if err := handler.Service.CreateUserTag(&userTag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}

		if err := handler.TagService.CreateTag(&userTag.Tag); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	userTagIDJson, _ := json.Marshal(userTag.ID)
	w.Write(userTagIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserTagHandler) FindAllTaggableUsersPost(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var allTaggableUsersPost []model.UserTag

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		fmt.Println("The is userId", userId)

		var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)

		fmt.Println("The is userId", userProfileSettings.UserVisibility)

		fmt.Println("The is userId", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsPostTaggable {
			allTaggableUsersPost = append(allTaggableUsersPost, userTags)
		}
	}

	allTaggableUsersPostJson, _ := json.Marshal(allTaggableUsersPost)
	w.Write(allTaggableUsersPostJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *UserTagHandler) FindAllTaggableUsersStory(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var allTaggableUsersStory []model.UserTag

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		userId := userTags.UserId
		fmt.Println("The is userId", userId)
		userProfileSettings :=  handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		if userProfileSettings.IsStoryTaggable == true {
			allTaggableUsersStory = append(allTaggableUsersStory, userTags)
		}
	}

	allTaggableUsersStoryJson, _ := json.Marshal(allTaggableUsersStory)
	w.Write(allTaggableUsersStoryJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserTagHandler) FindAllTaggableUsersComment(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var allTaggableUsersComment []model.UserTag

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		userId := userTags.UserId
		fmt.Println("The is userId", userId)
		userProfileSettings :=  handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		if userProfileSettings.IsCommentTaggable == true {
			allTaggableUsersComment = append(allTaggableUsersComment, userTags)
		}
	}

	allTaggableUsersCommentJson, _ := json.Marshal(allTaggableUsersComment)
	w.Write(allTaggableUsersCommentJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

