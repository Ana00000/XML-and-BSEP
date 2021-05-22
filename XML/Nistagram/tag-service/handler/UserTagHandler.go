package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	profileSettingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userModel "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UserTagHandler struct {
	Service *service.UserTagService
	TagService *service.TagService
	Validator *validator.Validate
	ProfileSettingsService *profileSettingsService.ProfileSettingsService
	ClassicUserService *userService.ClassicUserService
}

func (handler *UserTagHandler) CreateUserTag(w http.ResponseWriter, r *http.Request) {
	var userTagDTO dto.UserTagDTO
	if err := json.NewDecoder(r.Body).Decode(&userTagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&userTagDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
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
	var userAllTags []model.UserTag

	allUserTags = handler.Service.FindAll()

	fmt.Print(len(allUserTags))

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		fmt.Println("The is user ID: ", userId)
		var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		fmt.Println("User post taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsPostTaggable {
			user := handler.ClassicUserService.FindById(userId)
			if !user.IsDeleted && user.IsConfirmed {
				userAllTags = append(userAllTags, userTags)
			}
		}
	}

	userAllTagsJson, _ := json.Marshal(userAllTags)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userAllTagsJson)

}

func (handler *UserTagHandler) FindAllTaggableUsersStory(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var allTaggableUsersStory []userModel.ClassicUser

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		fmt.Println("The is user ID: ", userId)
		var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		fmt.Println("User story taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsStoryTaggable {
			user := handler.ClassicUserService.FindById(userId)
			if !user.IsDeleted && user.IsConfirmed {
				allTaggableUsersStory = append(allTaggableUsersStory, *user)
			}
		}
	}

	allTaggableUsersStoryJson, _ := json.Marshal(allTaggableUsersStory)
	w.Write(allTaggableUsersStoryJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserTagHandler) FindAllTaggableUsersComment(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var allTaggableUsersComment []userModel.ClassicUser

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		fmt.Println("The is user ID: ", userId)
		var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		fmt.Println("User comment taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsCommentTaggable {
			user := handler.ClassicUserService.FindById(userId)
			if !user.IsDeleted && user.IsConfirmed {
				allTaggableUsersComment = append(allTaggableUsersComment, *user)
			}
		}
	}

	allTaggableUsersCommentJson, _ := json.Marshal(allTaggableUsersComment)
	w.Write(allTaggableUsersCommentJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

