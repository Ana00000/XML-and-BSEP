package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type ConfirmationTokenHandler struct {
	ConfirmationTokenService * service.ConfirmationTokenService
	ClassicUserService * service.ClassicUserService
	RegisteredUserService * service.RegisteredUserService
	UserService * service.UserService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//VERFYCONFTOK322
func (handler *ConfirmationTokenHandler) VerifyConfirmationToken(w http.ResponseWriter, r *http.Request) {
	var confirmationAccountDTO dto.ConfirmationAccountDTO
	err := json.NewDecoder(r.Body).Decode(&confirmationAccountDTO)
	if err!=nil{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to ConfirmationAccountDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIdUUID := confirmationAccountDTO.UserId
	tokenUUID:= confirmationAccountDTO.ConfirmationToken

	var confirmationToken= handler.ConfirmationTokenService.FindByToken(tokenUUID)
	if !confirmationToken.IsValid{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Confirmation token isn't valid!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if confirmationToken.UserId!=userIdUUID || confirmationToken.ExpiredTime.Before(time.Now()){
		err := handler.ConfirmationTokenService.UpdateConfirmationTokenValidity(confirmationToken.ConfirmationToken, false)
		if err != nil {
			return 
		}
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Confirmation token isn't valid!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegisteredUserService.UpdateRegisteredUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating register user to confirmed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.ClassicUserService.UpdateClassicUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating classic user to confirmed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ConfirmationTokenHandler",
			"action":   "VERFYCONFTOK322",
			"timestamp":   time.Now().String(),
		}).Error("Failed updating basic user to confirmed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ConfirmationTokenHandler",
		"action":   "VERFYCONFTOK322",
		"timestamp":   time.Now().String(),
	}).Info("Successfully verifed account with token!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}