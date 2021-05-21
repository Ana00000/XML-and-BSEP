package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"os"
	_ "strconv"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
}

func (handler *ClassicUserHandler) FindSelectedUserById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var user = handler.ClassicUserService.FindSelectedUserById(uuid.MustParse(id))
	if user == nil {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
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
	//izmjenjeno da dobija requestove iz request microservica listu FollowerRequestForUserDTO i radi posle sa njom
	var  allFollowRequestsForUser []dto.FollowRequestForUserDTO
	reqUrlFollowRequests := fmt.Sprintf("http://%s:%s/find_all_requests_by_user_id/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"), logId)
	err = getJson(reqUrlFollowRequests, &allFollowRequestsForUser)
	if err!=nil{
		fmt.Println("Wrong cast response body to list FollowerRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	//var allFollowRequestsForUser = handler.FollowRequestService.FindAllFollowerRequestsForUser(uuid.MustParse(logId))
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
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	/*if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}*/

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ClassicUserHandler) FindAllPublicUsers(w http.ResponseWriter, r *http.Request) {

	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	var publicProfiles = handler.ProfileSettingsService.FindAllPublicUsers(allValidUsers)

	publicJson, _ := json.Marshal(publicProfiles)
	w.Write(publicJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

