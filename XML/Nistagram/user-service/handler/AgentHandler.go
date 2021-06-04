package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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
}

func (handler *AgentHandler) CreateAgent(w http.ResponseWriter, r *http.Request) {
	var agentDTO dto.AgentDTO
	if err := json.NewDecoder(r.Body).Decode(&agentDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&agentDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(agentDTO.Username) != nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(agentDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(agentDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(agentDTO.Password)
	} else {
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
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.ClassicUserService.CreateClassicUser(&agent.ClassicUser); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.UserService.CreateUser(&agent.ClassicUser.User); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	tagId := uuid.New()
	var userTag = dto.UserTagFullDTO{
		ID:     tagId,
		Name:   agent.Username,
		UserId: agentId,
	}

	reqUrl := fmt.Sprintf("http://%s:%s/create_user_tag_for_registered_user/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonOrders, _ := json.Marshal(userTag)
	fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	fmt.Println(string(jsonOrders))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 404 {
		print("Failed creating profile settings for user")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
