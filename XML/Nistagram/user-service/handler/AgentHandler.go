package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type AgentHandler struct {
	AgentService       *service.AgentService
	UserService        *service.UserService
	ClassicUserService *service.ClassicUserService
	Validator          *validator.Validate
	PasswordUtil       *util.PasswordUtil
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) {
	var agentDTO dto.AgentDTO
	if err := json.NewDecoder(r.Body).Decode(&agentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to AdminDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&agentDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("AgentDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(agentDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Already exist user with entered username!!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(agentDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Already exist user with entered email!!")
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(agentDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(agentDTO.Password)
	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Entered password isn't in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch agentDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	agentId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, agentDTO.DateOfBirth)
	agent := model.Agent{
		ClassicUser: model.ClassicUser{
			User: model.User{
				ID:          agentId,
				Username:    agentDTO.Username,
				Password:    password,
				Email:       agentDTO.Email,
				PhoneNumber: agentDTO.PhoneNumber,
				FirstName:   agentDTO.FirstName,
				LastName:    agentDTO.LastName,
				Gender:      gender,
				DateOfBirth: dateOfBirth,
				Website:     agentDTO.Website,
				Biography:   agentDTO.Biography,
				Salt:        salt,
				IsConfirmed: false,
				UserType:    model.AGENT,
			},
			IsDeleted: false,
		},
		AgentRegistrationRequestId: agentDTO.AgentRegistrationRequestId,
	}

	if err := handler.AgentService.CreateAgent(&agent); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating agent!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.ClassicUserService.CreateClassicUser(&agent.ClassicUser); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating classic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.UserService.CreateUser(&agent.ClassicUser.User); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	tagId:=uuid.New()
	var userTag = dto.UserTagFullDTO{
		ID:     tagId,
		Name:   agent.Username,
		UserId: agentId,
	}

	reqUrl := fmt.Sprintf("http://%s:%s/create_user_tag_for_registered_user/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonOrders, _ := json.Marshal(userTag)
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "AgentHandler",
		"action":   "CRAGT823",
		"timestamp":   time.Now().String(),
	}).Info("JSON sended to POST req to url "+reqUrl+" :\n"+string(jsonOrders))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 404 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AgentHandler",
			"action":   "CRAGT823",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating user tag for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "AgentHandler",
		"action":   "CRAGT823",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created admin!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
