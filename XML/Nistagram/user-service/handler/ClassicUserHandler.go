package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	requestsService "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
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
	FollowRequestService * requestsService.FollowRequestService
}

func (handler *ClassicUserHandler) FindSelectedUserById(w http.ResponseWriter, r *http.Request) {

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

	var allFollowRequestsForUser = handler.FollowRequestService.FindAllFollowerRequestsForUser(uuid.MustParse(logId))
	fmt.Println("USPEO1")
	var checkFollowingStatus = handler.ClassicUserFollowingsService.CheckFollowingStatus(uuid.MustParse(logId),uuid.MustParse(id),allFollowRequestsForUser)
	if (checkFollowingStatus == "FOLLOWING") || (checkFollowingStatus == "NOT FOLLOWING") || (checkFollowingStatus == "PENDING"){
		user.FollowingStatus = checkFollowingStatus
		fmt.Println("USPEO2")
	}else{
		fmt.Println("Check if following failed")
		w.WriteHeader(http.StatusExpectationFailed)
	}


	userJson, _ := json.Marshal(user)
	w.Write(userJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ClassicUserHandler) FindAllUsersButLoggedIn(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var user = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

