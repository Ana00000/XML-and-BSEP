package handler

import (
	"encoding/json"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type ConfirmationTokenHandler struct {
	ConfirmationTokenService *service.ConfirmationTokenService
	ClassicUserService       *service.ClassicUserService
	RegisteredUserService    *service.RegisteredUserService
	UserService              *service.UserService
}

func (handler *ConfirmationTokenHandler) VerifyConfirmationToken(w http.ResponseWriter, r *http.Request) {
	var confirmationAccountDTO dto.ConfirmationAccountDTO
	err := json.NewDecoder(r.Body).Decode(&confirmationAccountDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIdUUID := confirmationAccountDTO.UserId
	tokenUUID := confirmationAccountDTO.ConfirmationToken

	var confirmationToken = handler.ConfirmationTokenService.FindByToken(tokenUUID)
	if !confirmationToken.IsValid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if confirmationToken.UserId != userIdUUID || confirmationToken.ExpiredTime.Before(time.Now()) {
		err := handler.ConfirmationTokenService.UpdateConfirmationTokenValidity(confirmationToken.ConfirmationToken, false)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegisteredUserService.UpdateRegisteredUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.ClassicUserService.UpdateClassicUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
