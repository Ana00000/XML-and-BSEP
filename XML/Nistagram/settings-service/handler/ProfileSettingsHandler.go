package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"net/http"
	"settings-service/dto"
	_ "strconv"
)

type ProfileSettingsHandler struct {
	Service * service.ProfileSettingsService
}

func (handler *ProfileSettingsHandler) CreateProfileSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]

	profileSettings := model.ProfileSettings{
		ID:          uuid.UUID{},
		UserId: uuid.MustParse(userId),
		UserVisibility:      model.PUBLIC_VISIBILITY,
		MessageApprovalType:       model.PUBLIC,
		IsPostTaggable: true,
		IsStoryTaggable: true,
		IsCommentTaggable: true,
	}

	err := handler.Service.CreateProfileSettings(&profileSettings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ProfileSettingsHandler) FindProfileSettingByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]

	var profileSettings = handler.Service.FindProfileSettingByUserId(uuid.MustParse(userId))
	if profileSettings == nil {
		fmt.Println("Ne postoji!")
		w.WriteHeader(http.StatusNotFound)
	}

	userVisibility :=""
	if profileSettings.UserVisibility == model.PRIVATE_VISIBILITY{
		userVisibility = "PRIVATE_VISIBILITY"
	} else if profileSettings.UserVisibility == model.PUBLIC_VISIBILITY{
		userVisibility = "PUBLIC_VISIBILITY"
	}

	messageApprovalType:=""
	if profileSettings.MessageApprovalType == model.PUBLIC{
		messageApprovalType = "PUBLIC"
	} else if profileSettings.MessageApprovalType == model.FRIENDS_ONLY{
		messageApprovalType = "FRIENDS_ONLY"
	}

	var profileSettingsDTO = dto.ProfileSettingsDTO{
		UserId:              profileSettings.UserId,
		UserVisibility:      userVisibility,
		MessageApprovalType: messageApprovalType,
		IsPostTaggable:      profileSettings.IsPostTaggable,
		IsStoryTaggable:     profileSettings.IsStoryTaggable,
		IsCommentTaggable:   profileSettings.IsCommentTaggable,
	}

	profileSettingsDTOJson, _ := json.Marshal(profileSettingsDTO)
	if profileSettingsDTOJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(profileSettingsDTOJson)
	}
	w.WriteHeader(http.StatusNotFound)
}

func (handler *ProfileSettingsHandler) FindProfileSettingsForPublicUsers(w http.ResponseWriter, r *http.Request) {

	var profileSettings = handler.Service.FindAllProfileSettingsForPublicUsers()
	if profileSettings == nil {
		fmt.Println("Ne postoji!")
		w.WriteHeader(http.StatusNotFound)
	}

	dataJson, _ := json.Marshal(convertListUUIDToListData(profileSettings))
	if dataJson != nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
	}
	w.WriteHeader(http.StatusNotFound)
}

type Data struct {
	Uuid uuid.UUID
}

func convertListUUIDToListData(uuids []uuid.UUID) []Data {
	var datas []Data
	for i := 0; i < len(uuids); i++ {
		datas=append(datas, Data{Uuid: uuids[i]})
	}
	return datas
}
