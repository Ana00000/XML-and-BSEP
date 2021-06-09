package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"time"
)

type UserTagHandler struct {
	Service *service.UserTagService
	TagService *service.TagService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}
//CRUSRTG9821
func (handler *UserTagHandler) CreateUserTag(w http.ResponseWriter, r *http.Request) {
	var userTagDTO dto.UserTagDTO
	if err := json.NewDecoder(r.Body).Decode(&userTagDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTG9821",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to UserTagDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	if err := handler.Validator.Struct(&userTagDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTG9821",
			"timestamp":   time.Now().String(),
		}).Error("UserTagDTO fields doesn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var findTag = handler.TagService.FindTagByName(userTagDTO.Name)
	var userTag model.UserTag

	if findTag != nil && userTagDTO.TagType == "USER_TAG" {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTG9821",
			"timestamp":   time.Now().String(),
		}).Error("Tag type isn't valid or not founded tag bv name!")
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
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserTagHandler",
				"action":   "CRUSRTG9821",
				"timestamp":   time.Now().String(),
			}).Error("Failed creating user tag!")
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}

		if err := handler.TagService.CreateTag(&userTag.Tag); err != nil {
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserTagHandler",
				"action":   "CRUSRTG9821",
				"timestamp":   time.Now().String(),
			}).Error("Failed creating tag!")
			w.WriteHeader(http.StatusExpectationFailed) // 417
			return
		}
	}

	userTagIDJson, _ := json.Marshal(userTag.ID)
	w.Write(userTagIDJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "UserTagHandler",
		"action":   "CRUSRTG9821",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created user tag!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
//FIDALTAGBLUSRSPST2889
func (handler *UserTagHandler) FindAllTaggableUsersPost(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var userAllTags []model.UserTag

	allUserTags = handler.Service.FindAll()

	fmt.Print(len(allUserTags))

	for _,userTags := range allUserTags {
		//fmt.Println("UserTags: "+userTags.Name+" userId "+userTags.UserId.String())
		var userId = userTags.UserId
		//fmt.Println("The is user ID: ", userId)
		//var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		var userProfileSettings dto.ProfileSettingsDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), userId)
		err := getJson(reqUrl, &userProfileSettings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserTagHandler",
				"action":   "FIDALTAGBLUSRSPST2889",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast response body to ProfileSettingDTO!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}

		//fmt.Println("ID of userProfileSettings for user "+ userProfileSettings.UserId.String())
		//fmt.Println("User post taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsPostTaggable {
			//user := handler.ClassicUserService.FindById(userId)
			var user dto.ClassicUserFullDTO
			reqUrl := fmt.Sprintf("http://%s:%s/find_classic_user_by_id/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
			err := getJson(reqUrl, &user)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "UserTagHandler",
					"action":   "FIDALTAGBLUSRSPST2889",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast response body to ClassicUserFullDTO!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			/*fmt.Println("Username of user ",user.Username)
			fmt.Println("User is deleted: ", user.IsDeleted)
			fmt.Println("User is confirmed: ", user.IsConfirmed)*/
			if !user.IsDeleted && user.IsConfirmed {
				//fmt.Println("Tag name: "+userTags.Name)
				userAllTags = append(userAllTags, userTags)
			}
		}
	}

	userAllTagsJson, _ := json.Marshal(userAllTags)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "UserTagHandler",
		"action":   "FIDALTAGBLUSRSPST2889",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all taggable users in posts!")
	w.Write(userAllTagsJson)

}
//FIDALTAGBLUSRSSTRY8229
func (handler *UserTagHandler) FindAllTaggableUsersStory(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var userAllTags []model.UserTag

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		//fmt.Println("The is user ID: ", userId)
		//var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		var userProfileSettings dto.ProfileSettingsDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), userId)
		err := getJson(reqUrl, &userProfileSettings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserTagHandler",
				"action":   "FIDALTAGBLUSRSSTRY8229",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast response body to ProfileSettingDTO!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		//fmt.Println("User story taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsStoryTaggable {
			//user := handler.ClassicUserService.FindById(userId)
			var user dto.ClassicUserFullDTO
			reqUrl := fmt.Sprintf("http://%s:%s/find_classic_user_by_id/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
			err := getJson(reqUrl, &user)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "UserTagHandler",
					"action":   "FIDALTAGBLUSRSSTRY8229",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast response body to ClassicUserFullDTO!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			if !user.IsDeleted && user.IsConfirmed {
				userAllTags = append(userAllTags, userTags)
			}
		}
	}

	userAllTagsJson, _ := json.Marshal(userAllTags)
	w.Write(userAllTagsJson)
	w.WriteHeader(http.StatusOK)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "UserTagHandler",
		"action":   "FIDALTAGBLUSRSSTRY8229",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all taggable users in stories!")
	w.Header().Set("Content-Type", "application/json")
}
//FIDALTAGBLUSRSCOMM9882
func (handler *UserTagHandler) FindAllTaggableUsersComment(w http.ResponseWriter, r *http.Request) {
	var allUserTags []model.UserTag
	var userAllTags []model.UserTag

	allUserTags = handler.Service.FindAll()

	for _,userTags := range allUserTags {
		var userId = userTags.UserId
		//fmt.Println("The is user ID: ", userId)
		//var userProfileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(userId)
		var userProfileSettings dto.ProfileSettingsDTO
		reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), userId)
		err := getJson(reqUrl, &userProfileSettings)
		if err!=nil{
			handler.LogError.WithFields(logrus.Fields{
				"status": "failure",
				"location":   "UserTagHandler",
				"action":   "FIDALTAGBLUSRSCOMM9882",
				"timestamp":   time.Now().String(),
			}).Error("Wrong cast response body to ProfileSettingDTO!")
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		//fmt.Println("User comment taggable: ", userProfileSettings.IsPostTaggable)
		if userProfileSettings.IsCommentTaggable {
			//user := handler.ClassicUserService.FindById(userId)
			var user dto.ClassicUserFullDTO
			reqUrl := fmt.Sprintf("http://%s:%s/find_classic_user_by_id/%s", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"), userId)
			err := getJson(reqUrl, &user)
			if err!=nil{
				handler.LogError.WithFields(logrus.Fields{
					"status": "failure",
					"location":   "UserTagHandler",
					"action":   "FIDALTAGBLUSRSCOMM9882",
					"timestamp":   time.Now().String(),
				}).Error("Wrong cast response body to ClassicUserFullDTO!")
				//fmt.Println("Wrong cast response body to ProfileSettingDTO!")
				w.WriteHeader(http.StatusExpectationFailed)
				return
			}
			if !user.IsDeleted && user.IsConfirmed {
				userAllTags = append(userAllTags, userTags)
			}
		}
	}

	userAllTagsJson, _ := json.Marshal(userAllTags)
	w.Write(userAllTagsJson)
	w.WriteHeader(http.StatusOK)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "UserTagHandler",
		"action":   "FIDALTAGBLUSRSCOMM9882",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all taggable users in comments!")
	w.Header().Set("Content-Type", "application/json")
}
//CRUSRTGFORREGUSR7772
func (handler *UserTagHandler) CreateUserTagForRegisteredUser(w http.ResponseWriter, r *http.Request) {
	var userTagDTO dto.UserTagFullDTO
	if err := json.NewDecoder(r.Body).Decode(&userTagDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTGFORREGUSR7772",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to UserTagFullDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var userTag = model.UserTag{
		Tag:    model.Tag{
			ID:      userTagDTO.ID,
			Name:    userTagDTO.Name,
			TagType: model.USER_TAG,
		},
		UserId: uuid.UUID{},
	}

	if err := handler.Service.CreateUserTag(&userTag); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTGFORREGUSR7772",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating user tag for registered user")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.TagService.CreateTag(&userTag.Tag); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "UserTagHandler",
			"action":   "CRUSRTGFORREGUSR7772",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating tag for registered user")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "UserTagHandler",
		"action":   "CRUSRTGFORREGUSR7772",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created user tag for registered user!")
	userAllTagsJson, _ := json.Marshal(userTag)
	w.Write(userAllTagsJson)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
