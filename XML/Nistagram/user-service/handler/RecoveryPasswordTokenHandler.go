package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"gopkg.in/go-playground/validator.v9"
	gomail "gopkg.in/mail.v2"
	"net/http"
	_ "strconv"
	"time"
)

type RecoveryPasswordTokenHandler struct {
	RecoveryPasswordTokenService * service.RecoveryPasswordTokenService
	UserService * service.UserService
	Validator   *validator.Validate
}

func SendRecoveryPasswordMail(user *model.User, token uuid.UUID) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", user.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Recovery password email")

	// Set E-Mail body. You can set plain text or html with text/html
	text:= "Dear "+user.FirstName+",\n\nPlease, click on link in below to change your password on our social network!\n\nhttp://localhost:8082//changePasswordByToken/"+token.String()+"/"+user.ID.String()+"\n\nBest regards,\nTim25"
	m.SetBody("text/plain", text)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "xml.ftn.uns@gmail.com", "XMLFTNUNS1")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//Function when user clicks -> FORGOT PASSWORD -> enters email -> clicks RECOVER to get email
func (handler *RecoveryPasswordTokenHandler) GenerateRecoveryPasswordToken (w http.ResponseWriter, r *http.Request) {
	var emailDTO dto.EmailDTO
	if err := json.NewDecoder(r.Body).Decode(&emailDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.Validator.Struct(&emailDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	var user = handler.UserService.FindByEmail(emailDTO.Email)
	recoveryPasswordToken := model.RecoveryPasswordToken{
		ID:                    uuid.New(),
		RecoveryPasswordToken: uuid.New(),
		UserId:                user.ID,
		CreatedTime:           time.Now(),
		ExpirationTime:        time.Now().Add(time.Minute * 5),
		IsValid:               true,
	}

	err := handler.RecoveryPasswordTokenService.CreateRecoveryPasswordToken(&recoveryPasswordToken)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	SendRecoveryPasswordMail(user, recoveryPasswordToken.RecoveryPasswordToken)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

//Function that gets called when USER clicks on the link in the email
func (handler *RecoveryPasswordTokenHandler) VerifyRecoveryPasswordToken(w http.ResponseWriter, r *http.Request) {
	var recoveryPasswordDTO dto.RecoveryPasswordDTO

	err := json.NewDecoder(r.Body).Decode(&recoveryPasswordDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIdUUID := recoveryPasswordDTO.UserId
	tokenUUID:= recoveryPasswordDTO.RecoveryPasswordToken

	var recoveryPasswordToken= handler.RecoveryPasswordTokenService.FindByToken(tokenUUID)
	if !recoveryPasswordToken.IsValid{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if recoveryPasswordToken.UserId!=userIdUUID || recoveryPasswordToken.ExpirationTime.Before(time.Now()){
		err := handler.RecoveryPasswordTokenService.UpdateRecoveryPasswordTokenValidity(recoveryPasswordToken.RecoveryPasswordToken, false)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.UserService.FindByID(userIdUUID)
	emailJson, _ := json.Marshal(user.Email)
	w.Write(emailJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
