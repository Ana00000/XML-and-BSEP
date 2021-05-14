package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
	ProfileSettingsService * settingsService.ProfileSettingsService
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
}

func (handler *ClassicUserHandler) FindSelectedUserById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("USAO  BACK 2 found")
	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var user = handler.ClassicUserService.FindSelectedUserById(uuid.MustParse(id))
	if user == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	var profileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings == nil {
		fmt.Println("Profile settings not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	if profileSettings.UserVisibility == settingsModel.PRIVATE_VISIBILITY {
		user.ProfileVisibility = "PRIVATE"
		fmt.Println("PRIVATE")
	} else {
		user.ProfileVisibility = "PUBLIC"
		fmt.Println("PUBLIC")
	}

	var checkIfFollowing = handler.ClassicUserFollowingsService.CheckIfFollowingUser(uuid.MustParse(logId),uuid.MustParse(id))
	if checkIfFollowing == false {
		user.FollowingCheck = false
		fmt.Println("FALSE")
	} else {
		user.FollowingCheck = true
		fmt.Println("TRUE")
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

