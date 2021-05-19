package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"net/http"
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

