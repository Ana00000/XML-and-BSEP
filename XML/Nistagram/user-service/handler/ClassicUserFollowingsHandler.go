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
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type ClassicUserFollowingsHandler struct {
	ClassicUserFollowingsService * service.ClassicUserFollowingsService
	ClassicUserFollowersService * service.ClassicUserFollowersService
	ClassicUserService * service.ClassicUserService
	UserService *service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateClassicUserFollowing *gorbac.Permission
	PermissionAcceptFollowerRequest *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//CRCLASUSFOLLING712
// CreateClassicUserFollowing KAD NEKO KLIKNE FOLLOW NEKOGA = NJEMU SE KREIRA PRVO FOLLOWING PA ONDA FOLLOWER OVOM DRUGOM
func (handler *ClassicUserFollowingsHandler) CreateClassicUserFollowing(w http.ResponseWriter, r *http.Request) {

	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(user)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateClassicUserFollowing, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var classicUserFollowingDTO dto.ClassicUserFollowingsDTO
	err = json.NewDecoder(r.Body).Decode(&classicUserFollowingDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to list ClassicUserFollowingsDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	classicUserFollowings := model.ClassicUserFollowings{
		ID:              uuid.UUID{},
		ClassicUserId:   classicUserFollowingDTO.ClassicUserId,
		FollowingUserId: classicUserFollowingDTO.FollowingUserId,
	}

	err = handler.ClassicUserFollowingsService.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating claassic user followings!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	classicUserFollower := model.ClassicUserFollowers{
		ID:             uuid.UUID{},
		ClassicUserId:  classicUserFollowingDTO.FollowingUserId,
		FollowerUserId: classicUserFollowingDTO.ClassicUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollower)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating claassic user followers!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "CRCLASUSFOLLING712",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created followings and followers for classic user!")

	userReceiver := handler.UserService.FindByID(classicUserFollowingDTO.FollowingUserId)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	userSender := handler.UserService.FindByID(classicUserFollowingDTO.ClassicUserId)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "CRCLASUSFOLLING712",
			"timestamp":   time.Now().String(),
		}).Error("Failed to find user by id!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//SEND EMAIL NOTIFICATION
	handler.SendNotificationMail(userReceiver.Email, userSender.Username)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *ClassicUserFollowingsHandler) SendNotificationMail(email string, username string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	// Set E-Mail body. You can set plain text or html with text/html
	text := username + "\n\n\n is now following you!\n\n\nBest regards,\nTim25"

	m.SetBody("text/plain", text)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "xml.ftn.uns@gmail.com", "XMLFTNUNS1")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		//fmt.Println(err)
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "SEDCONFMAIL711",
			"timestamp":   time.Now().String(),
		}).Error("Failed sending email!")
		panic(err)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "SEDCONFMAIL711",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sent email!")
}

type ReturnValueBool struct {
	ReturnValue bool `json:"return_value"`
}

//FIDALVALFOLLINGSFRUS111
func (handler *ClassicUserFollowingsHandler) FindAllValidFollowingsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	id := vars["id"]

	var classicUserDTO []dto.ClassicUserDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "FIDALVALFOLLINGSFRUS111",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to list ClassicUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var validFollowingsForUser = handler.ClassicUserFollowingsService.FindAllValidFollowingsForUser(uuid.MustParse(id), convertListClassicUserDTOToListClassicUser(classicUserDTO))
	for i := 0; i < len(validFollowingsForUser); i++ {
		fmt.Println(validFollowingsForUser[i].FollowingUserId)
	}

	returnValueJson, _ := json.Marshal(convertListClassicUsersFollowingsToListClassicUsersFollowingsDTO(validFollowingsForUser))

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "FIDALVALFOLLINGSFRUS111",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all valid followings for user! Result : "+string(returnValueJson))

	w.Write(returnValueJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func  (handler *ClassicUserFollowingsHandler) FindAllValidUsersWhoFollowUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	id := vars["id"]

	var users = listWithoutLoggedInUser(uuid.MustParse(id),handler.ClassicUserService.FinAllValidUsers())
	var followings = handler.ClassicUserFollowingsService.FindAllUsersWhoFollowUserId(uuid.MustParse(id),users)

	returnValueJson,_ := json.Marshal(followings)

	w.Write(returnValueJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func listWithoutLoggedInUser(id uuid.UUID, users []model.ClassicUser) []model.ClassicUser{
	var retVal []model.ClassicUser
	for i := 0; i < len(users); i++ {
		if users[i].ID!=id{
			retVal = append(retVal, users[i])
		}
	}
	return retVal
}

//FIDALUSFOLWUSID2672
func (handler *ClassicUserFollowingsHandler) FindAllUserWhoFollowUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	id := vars["id"]

	var classicUserDTO []dto.ClassicUserDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "FIDALUSFOLWUSID2672",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to ClassicUserDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var validFollowingsForUser = handler.ClassicUserFollowingsService.FindAllUserWhoFollowUserId(uuid.MustParse(id), convertListClassicUserDTOToListClassicUser(classicUserDTO))

	returnValueJson, _ := json.Marshal(convertListClassicUsersFollowingsToListClassicUsersFollowingsDTO(validFollowingsForUser))

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "FIDALUSFOLWUSID2672",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all users who follow user id! Result : "+string(returnValueJson))

	w.Write(returnValueJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//CHEKFOLLINGPSTSTRY2111
func (handler *ClassicUserFollowingsHandler) CheckIfFollowingPostStory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	vars := mux.Vars(r)
	id := vars["id"]
	logId := vars["logId"]
	//izmjenio
	var check = handler.ClassicUserFollowingsService.CheckIfFollowingPostStory(uuid.MustParse(id), uuid.MustParse(logId))

	var returnValue = ReturnValueBool{ReturnValue: check}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "CHEKFOLLINGPSTSTRY2111",
		"timestamp":   time.Now().String(),
	}).Info("Successfully checked following post/story! Result : "+convertBoolToString(check))

	returnValueJson, _ := json.Marshal(returnValue)
	w.Write(returnValueJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//ACCFOLLERREQ832
func (handler *ClassicUserFollowingsHandler) AcceptFollowerRequest(w http.ResponseWriter, r *http.Request) {
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(user)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionAcceptFollowerRequest, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var followRequestDTO dto.FollowRequestDTO
	err = json.NewDecoder(r.Body).Decode(&followRequestDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to FollowRequestDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var followRequestForUser dto.FollowRequestForUserDTO
	reqUrlFollowRequests := fmt.Sprintf("http://%s:%s/find_request_by_classic_user_and_follower_user_ids/%s/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"), followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)
	err = getJson(reqUrlFollowRequests, &followRequestForUser)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast response body to FollowRequestForUserDTO or error with finding follow request!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// UPDATE REQUEST - ACCEPTED
	//var request = handler.FollowRequestService.FindFollowRequest(followRequestDTO.ClassicUserId, followRequestDTO.FollowerUserId)

	reqUrlUpdate := fmt.Sprintf("http://%s:%s/accept_follow_request/%s", os.Getenv("REQUESTS_SERVICE_DOMAIN"), os.Getenv("REQUESTS_SERVICE_PORT"), followRequestForUser.ID)
	jsonOrders, _ := json.Marshal(nil)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrlUpdate)
	//fmt.Println(string(jsonOrders))
	resp, err := http.Post(reqUrlUpdate, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 400 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Failed accepting follow request!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	// CREATE FOLLOWER
	classicUserFollowers := model.ClassicUserFollowers{
		ID:             uuid.UUID{},
		ClassicUserId:  followRequestDTO.FollowerUserId,
		FollowerUserId: followRequestDTO.ClassicUserId,
	}

	err = handler.ClassicUserFollowersService.CreateClassicUserFollowers(&classicUserFollowers)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating follower for classic user!")
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// CREATE FOLLOWING
	classicUserFollowings := model.ClassicUserFollowings{
		ID:              uuid.UUID{},
		ClassicUserId:   followRequestDTO.ClassicUserId,
		FollowingUserId: followRequestDTO.FollowerUserId,
	}

	err = handler.ClassicUserFollowingsService.CreateClassicUserFollowings(&classicUserFollowings)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowingsHandler",
			"action":   "ACCFOLLERREQ832",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating following for classic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowingsHandler",
		"action":   "ACCFOLLERREQ832",
		"timestamp":   time.Now().String(),
	}).Info("Successfully accepted follower request!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func convertListClassicUsersFollowingsToListClassicUsersFollowingsDTO(classicUserFollowings []model.ClassicUserFollowings) []dto.ClassicUserFollowingsFullDTO {
	var classicUserFollowingsDTO []dto.ClassicUserFollowingsFullDTO
	for i := 0; i < len(classicUserFollowings); i++ {
		classicUserFollowingsDTO = append(classicUserFollowingsDTO, convertClassicUserFollowingsToClassicUserFollowingsDTO(classicUserFollowings[i]))
	}
	return classicUserFollowingsDTO
}

func convertClassicUserFollowingsToClassicUserFollowingsDTO(classicUserFollowings model.ClassicUserFollowings) dto.ClassicUserFollowingsFullDTO {
	var classicUserFollowingsFullDTO = dto.ClassicUserFollowingsFullDTO{
		ID:              classicUserFollowings.ID,
		ClassicUserId:   classicUserFollowings.ClassicUserId,
		FollowingUserId: classicUserFollowings.FollowingUserId,
	}
	return classicUserFollowingsFullDTO
}
