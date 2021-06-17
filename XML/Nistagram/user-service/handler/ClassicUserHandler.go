package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
	UserService *service.UserService
	Rbac * gorbac.RBAC
	PermissionFindAllUsersButLoggedIn *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//FIDSELUSBYID9993
func (handler *ClassicUserHandler) FindSelectedUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")
	logId := r.URL.Query().Get("logId")

	var user = handler.ClassicUserService.FindSelectedUserById(uuid.MustParse(id))
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDSELUSBYID9993",
			"timestamp":   time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	var profileSettings dto.ProfileSettingsDTO
	reqUrl := fmt.Sprintf("http://%s:%s/find_profile_settings_by_user_id/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), id)
	err := getJson(reqUrl, &profileSettings)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDSELUSBYID9993",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to ProfileSettingDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
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
		//fmt.Println("PRIVATE")
	} else {
		user.ProfileVisibility = "PUBLIC_VISIBILITY"
		//fmt.Println("PUBLIC")
	}
	//izmjenjeno da dobija requestove iz request microservica listu FollowerRequestForUserDTO i radi posle sa njom
	var allFollowRequestsForUser []dto.FollowRequestForUserDTO
	reqUrlFollowRequests := fmt.Sprintf("http://%s:%s/find_all_requests_by_user_id/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"), logId)
	err = getJson(reqUrlFollowRequests, &allFollowRequestsForUser)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDSELUSBYID9993",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to list FollowerRequestForUserDTO!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	//var allFollowRequestsForUser = handler.FollowRequestService.FindAllFollowerRequestsForUser(uuid.MustParse(logId))
	//fmt.Println("USPEO1")
	var checkFollowingStatus = handler.ClassicUserFollowingsService.CheckFollowingStatus(uuid.MustParse(logId),uuid.MustParse(id),allFollowRequestsForUser)
	if (checkFollowingStatus == "FOLLOWING") || (checkFollowingStatus == "NOT FOLLOWING") || (checkFollowingStatus == "PENDING"){
		user.FollowingStatus = checkFollowingStatus
		//fmt.Println("USPEO2")
	}else{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDSELUSBYID9993",
			"timestamp":   time.Now().String(),
		}).Error("Check if following failed!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	userJson, _ := json.Marshal(user)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDSELUSBYID9993",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded selected user by id! Result: "+string(userJson))

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

//FIDALUSRSBUTLOGEDIN3231
func (handler *ClassicUserHandler) FindAllUsersButLoggedIn(w http.ResponseWriter, r *http.Request) {

	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALUSRSBUTLOGEDIN3231",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALUSRSBUTLOGEDIN3231",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllUsersButLoggedIn, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALUSRSBUTLOGEDIN3231",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var user = handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	/*if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}*/
	//iz usera izbaciti sve usere koji su blokirani, i koji su usera blokirali
	userJson, _ := json.Marshal(user)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDALUSRSBUTLOGEDIN3231",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all users but logged in! Result: "+string(userJson))

	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALUSRSBUTLOGEDINDTO32
func (handler *ClassicUserHandler) FindAllUsersButLoggedInDTOs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var user = convertListClassicUserToListClassicUserDTO(handler.ClassicUserService.FindAllUsersButLoggedIn(uuid.MustParse(id)))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	/*if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}*/
	//iz usera izbaciti sve usere koji su blokirani, i koji su usera blokirali
	userJson, _ := json.Marshal(user)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDALUSRSBUTLOGEDINDTO32",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all users but logged in DTOs! Result: "+string(userJson))

	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALUSRSBYFOLLINGIDS9442
func (handler *ClassicUserHandler) FindAllUsersByFollowingIds(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var classicUserFollowingsFullDTO []dto.ClassicUserFollowingsFullDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserFollowingsFullDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALUSRSBYFOLLINGIDS9442",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast to list ClassicUserFollowingsFullDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var users = handler.ClassicUserService.FindAllUsersByFollowingIds(convertListClassicUserFollowingsDTOToListClassicUserFollowings(classicUserFollowingsFullDTO))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	/*if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}*/

	userJson, _ := json.Marshal(convertListClassicUserToListClassicUserDTO(users))

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDALUSRSBYFOLLINGIDS9442",
		"timestamp":   time.Now().String(),
	}).Info("Successfully checked user validity! Result: "+string(userJson))

	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//CHKIFUSVAL9929
func (handler *ClassicUserHandler) CheckIfUserValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]
	var isValid = handler.ClassicUserService.CheckIfUserValid(uuid.MustParse(userId))
	//CHECK IF THIS SHOULD RETURN ERROR OR JUST EMPTY LIST
	/*if  user == nil {
		fmt.Println("No user found")
		w.WriteHeader(http.StatusExpectationFailed)
	}*/

	var userValid = UserValidDTO{IsValid: isValid}
	userJson, _ := json.Marshal(userValid)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "CHKIFUSVAL9929",
		"timestamp":   time.Now().String(),
	}).Info("Successfully checked user validity! Result: "+string(userJson))

	w.Write(userJson)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

type UserValidDTO struct {
	IsValid bool `json:"is_valid"`
}

//FIDALPUBUSRS431
func (handler *ClassicUserHandler) FindAllPublicUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	reqUrlUpdate := fmt.Sprintf("http://%s:%s/find_all_public_users/", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"))
	jsonClassicUsersDTO, _ := json.Marshal(convertListClassicUserToListClassicUserDTO(allValidUsers))
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrlUpdate)
	//fmt.Println(string(jsonClassicUsersDTO))
	resp, err := http.Post(reqUrlUpdate, "application/json", bytes.NewBuffer(jsonClassicUsersDTO))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALPUBUSRS431",
			"timestamp":   time.Now().String(),
		}).Error("Failed with founded all public users!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	//defer resp.Body.Close() mozda treba dodati
	var listClassicUsersDTO []dto.ClassicUserDTO
	if err := json.NewDecoder(resp.Body).Decode(&listClassicUsersDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserHandler",
			"action":   "FIDALPUBUSRS431",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast to list ClassicUserDTO!")
		w.WriteHeader(http.StatusConflict) //400
		return
	}
	//izmjeniti
	//var publicProfiles = handler.ProfileSettingsService.FindAllPublicUsers(allValidUsers)
	publicJson, _ := json.Marshal(convertListClassicUserDTOToListClassicUser(listClassicUsersDTO))
	w.Write(publicJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDCLASSUSBYID943",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded classic user by id! Result: "+string(publicJson))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FIDALVALUSRS999
func (handler *ClassicUserHandler) FindAllValidUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var allValidUsers = handler.ClassicUserService.FinAllValidUsers()
	validJson, _ := json.Marshal(convertListClassicUserToListClassicUserDTO(allValidUsers))
	w.Write(validJson)

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDALVALUSRS999",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all valid users! Result: "+string(validJson))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func convertListClassicUserToListClassicUserDTO(classicUsers []model.ClassicUser) []dto.ClassicUserDTO {
	var classicUsersDTO []dto.ClassicUserDTO
	for i := 0; i < len(classicUsers); i++ {
		classicUsersDTO = append(classicUsersDTO, convertClassicUserToClassicUserDTO(classicUsers[i]))
	}
	return classicUsersDTO
}

func convertClassicUserToClassicUserDTO(classicUser model.ClassicUser) dto.ClassicUserDTO {
	layout := "2006-01-02T15:04:05.000Z"
	gender := ""
	userType := ""
	if classicUser.Gender == model.MALE {
		gender = "MALE"
	} else if classicUser.Gender == model.FEMALE {
		gender = "FEMALE"
	} else {
		gender = "OTHER"
	}
	if classicUser.UserType == model.REGISTERED_USER {
		userType = "REGISTERED_USER"
	} else if classicUser.UserType == model.AGENT {
		userType = "AGENT"
	} else if classicUser.UserType == model.ADMIN {
		userType = "ADMIN"
	}
	var classicUserDTO = dto.ClassicUserDTO{
		ID:          classicUser.ID,
		Username:    classicUser.Username,
		Password:    classicUser.Password,
		Email:       classicUser.Email,
		PhoneNumber: classicUser.Email,
		FirstName:   classicUser.FirstName,
		LastName:    classicUser.LastName,
		Gender:      gender,
		DateOfBirth: classicUser.DateOfBirth.Format(layout),
		Website:     classicUser.Website,
		Biography:   classicUser.Biography,
		Salt:        classicUser.Salt,
		IsConfirmed: classicUser.IsConfirmed,
		UserType:    userType,
		IsDeleted:   classicUser.IsDeleted,
	}
	return classicUserDTO
}

//FIDCLASSUSBYID943
func (handler *ClassicUserHandler) FindClassicUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	userId := vars["userID"]
	var classicUser = handler.ClassicUserService.FindById(uuid.MustParse(userId))
	validJson, _ := json.Marshal(convertClassicUserToClassicUserDTO(*classicUser))
	w.Write(validJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserHandler",
		"action":   "FIDCLASSUSBYID943",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded classic user by id! Result: "+string(validJson))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func convertListClassicUserDTOToListClassicUser(classicUsersDTO []dto.ClassicUserDTO) []model.ClassicUser {
	var classicUsers []model.ClassicUser
	for i := 0; i < len(classicUsersDTO); i++ {
		classicUsers = append(classicUsers, convertClassicUserDTOToClassicUser(classicUsersDTO[i]))
	}
	return classicUsers
}

func convertClassicUserDTOToClassicUser(classicUserDTO dto.ClassicUserDTO) model.ClassicUser {
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, classicUserDTO.DateOfBirth)
	var gender model.Gender
	var userType model.UserType
	if classicUserDTO.Gender == "MALE" {
		gender = model.MALE
	} else if classicUserDTO.Gender == "FEMALE" {
		gender = model.FEMALE
	} else {
		gender = model.OTHER
	}

	if classicUserDTO.UserType == "REGISTERED_USER" {
		userType = model.REGISTERED_USER
	} else if classicUserDTO.UserType == "AGENT" {
		userType = model.AGENT
	} else if classicUserDTO.UserType == "ADMIN" {
		userType = model.ADMIN
	}
	var classicUser = model.ClassicUser{
		User: model.User{
			ID:          classicUserDTO.ID,
			Username:    classicUserDTO.Username,
			Password:    classicUserDTO.Password,
			Email:       classicUserDTO.Email,
			PhoneNumber: classicUserDTO.Email,
			FirstName:   classicUserDTO.FirstName,
			LastName:    classicUserDTO.LastName,
			Gender:      gender,
			DateOfBirth: dateOfBirth,
			Website:     classicUserDTO.Website,
			Biography:   classicUserDTO.Biography,
			Salt:        classicUserDTO.Salt,
			IsConfirmed: classicUserDTO.IsConfirmed,
			UserType:    userType,
		},
		IsDeleted: false,
	}
	return classicUser
}
func convertClassicUserFollowingsDTOToClassicUserFollowings(classicUserFollowingsDTO dto.ClassicUserFollowingsFullDTO) model.ClassicUserFollowings {
	var classicUserFollowings = model.ClassicUserFollowings{
		ID:              classicUserFollowingsDTO.ID,
		ClassicUserId:   classicUserFollowingsDTO.ClassicUserId,
		FollowingUserId: classicUserFollowingsDTO.FollowingUserId,
	}
	return classicUserFollowings
}

func convertListClassicUserFollowingsDTOToListClassicUserFollowings(classicUserFollowingsDTOs []dto.ClassicUserFollowingsFullDTO) []model.ClassicUserFollowings {
	var classicUserFollowings []model.ClassicUserFollowings
	for i := 0; i < len(classicUserFollowingsDTOs); i++ {
		classicUserFollowings = append(classicUserFollowings, convertClassicUserFollowingsDTOToClassicUserFollowings(classicUserFollowingsDTOs[i]))
	}
	return classicUserFollowings
}
