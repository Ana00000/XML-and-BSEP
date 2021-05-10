package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
	"net/http"
	_ "strconv"
	"strings"
	"time"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
	UserService * service.UserService
	RegisteredUserService * service.RegisteredUserService
	ConfirmationTokenService * service.ConfirmationTokenService
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func SendConfirmationMail(user model.User, token uuid.UUID) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", user.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	// Set E-Mail body. You can set plain text or html with text/html
	text:= "Dear "+user.FirstName+",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttp://localhost:8082/confirm_registration/"+token.String()+"/"+user.ID.String()
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

func (handler *ClassicUserHandler) CreateClassicUser(w http.ResponseWriter, r *http.Request) {
	var classicUserDTO dto.ClassicUserDTO

	err := json.NewDecoder(r.Body).Decode(&classicUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var sb strings.Builder
	salt := uuid.New().String()
	sb.WriteString(classicUserDTO.Password)
	sb.WriteString(salt)
	password := sb.String()
	hash,_ := HashPassword(password)

	userId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,classicUserDTO.DateOfBirth)
	classicUser := model.ClassicUser{
		RegisteredUser:       model.RegisteredUser{
			User:                        model.User{
				ID:               userId,
				Username:         classicUserDTO.Username,
				Password:         hash,
				Email:            classicUserDTO.Email,
				PhoneNumber:      classicUserDTO.PhoneNumber,
				FirstName:        classicUserDTO.FirstName,
				LastName:         classicUserDTO.LastName,
				Gender:           classicUserDTO.Gender,
				DateOfBirth:      dateOfBirth,
				Website:          classicUserDTO.Website,
				Biography:        classicUserDTO.Biography,
				Salt: 			  salt,
				IsConfirmed: 	  false,
			},
		},
		IsBlocked:            false,
		UserCategory:         classicUserDTO.UserCategory,
		OfficialDocumentPath: classicUserDTO.OfficialDocumentPath,

	}

	err = handler.ClassicUserService.CreateClassicUser(&classicUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	err = handler.RegisteredUserService.CreateRegisteredUser(&classicUser.RegisteredUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	err = handler.UserService.CreateUser(&classicUser.RegisteredUser.User)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println(classicUser.RegisteredUser.User.ID)
	confirmationToken:= model.ConfirmationToken{
		ID:                uuid.New(),
		ConfirmationToken: uuid.New(),
		UserId:            userId,
		CreatedTime:       time.Now(),
		ExpiredTime:       time.Now().Add(time.Hour * 120),
		IsValid:           true,
	}
	err = handler.ConfirmationTokenService.CreateConfirmationToken(&confirmationToken)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	SendConfirmationMail(classicUser.RegisteredUser.User, confirmationToken.ConfirmationToken)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


