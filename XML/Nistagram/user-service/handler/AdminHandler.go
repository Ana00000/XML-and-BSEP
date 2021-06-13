package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type AdminHandler struct {
	AdminService * service.AdminService
	UserService *service.UserService
	Validator *validator.Validate
	LogInfo *logrus.Logger
	LogError *logrus.Logger
	PasswordUtil *util.PasswordUtil
}

func (handler *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var adminDTO dto.AdminDTO
	if err := json.NewDecoder(r.Body).Decode(&adminDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to AdminDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&adminDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error("AdminDTO fields aren't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(adminDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error("Already exist user with entered username!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(adminDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error("Already exist user with entered email!")
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(adminDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(adminDTO.Password)
	}else {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error("Entered password isn't in the valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch adminDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	adminId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,adminDTO.DateOfBirth)
	admin := model.Admin{
		User : model.User{
			ID:          adminId,
			Username:    adminDTO.Username,
			Password:    password,
			Email:       adminDTO.Email,
			PhoneNumber: adminDTO.PhoneNumber,
			FirstName:   adminDTO.FirstName,
			LastName:    adminDTO.LastName,
			Gender:      gender,
			DateOfBirth: dateOfBirth,
			Website:     adminDTO.Website,
			Biography:   adminDTO.Biography,
			Salt: salt,
			IsConfirmed: true,
			UserType: model.ADMIN,
		},
	}

	err := handler.AdminService.CreateAdmin(&admin)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "AdminHandler",
			"action":   "CRADM524",
			"timestamp":   time.Now().String(),
		}).Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "AdminHandler",
		"action":   "CRADM524",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created admin!")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

