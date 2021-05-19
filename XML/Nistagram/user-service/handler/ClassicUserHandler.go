package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	requestsService "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"os"
	_ "strconv"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
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
	var profileSettings = dto.ProfileSettingsDTO{}
	reqUrl := fmt.Sprintf("http://%s:%s/find_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err := getJson(reqUrl, &profileSettings)
	if err!=nil{
		fmt.Println("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	/*
	var profileSettings = handler.ProfileSettingsService.FindProfileSettingByUserId(uuid.MustParse(id))
	if profileSettings == nil {
		fmt.Println("Profile settings not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	*/
	if profileSettings.UserVisibility == "PRIVATE_VISIBILITY" {
		user.ProfileVisibility = "PRIVATE"
		fmt.Println("PRIVATE")
	} else {
		user.ProfileVisibility = "PUBLIC_VISIBILITY"
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

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (handler *ClassicUserHandler) FindAllUsersButLoggedIn(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var user = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

