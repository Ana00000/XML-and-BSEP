package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	settingsModel "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	gomail "gopkg.in/mail.v2"
	"net/http"
	_ "strconv"
	"time"
)

type RegisteredUserHandler struct {

	RegisteredUserService * service. RegisteredUserService
	UserService * service.UserService
	ClassicUserService * service.ClassicUserService
	ConfirmationTokenService * service.ConfirmationTokenService
	ProfileSettingsService * settingsService.ProfileSettingsService
	Validator                *validator.Validate
	PasswordUtil             *util.PasswordUtil
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
	text := "Dear " + user.FirstName + ",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttp://localhost:8081/confirmRegistration/" + token.String() + "/" + user.ID.String() + "\n\n\nBest regards,\nTim25"
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

func (handler *RegisteredUserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	var registeredUserDTO dto.RegisteredUserDTO
	if err := json.NewDecoder(r.Body).Decode(&registeredUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&registeredUserDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(registeredUserDTO.Username) != nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(registeredUserDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(registeredUserDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(registeredUserDTO.Password)
	}else {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch registeredUserDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	userId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, registeredUserDTO.DateOfBirth)
	registeredUser := model.RegisteredUser{
		ClassicUser: model.ClassicUser{
			User: model.User{
				ID:          userId,
				Username:    registeredUserDTO.Username,
				Password:    password,
				Email:       registeredUserDTO.Email,
				PhoneNumber: registeredUserDTO.PhoneNumber,
				FirstName:   registeredUserDTO.FirstName,
				LastName:    registeredUserDTO.LastName,
				Gender:      gender,
				DateOfBirth: dateOfBirth,
				Website:     registeredUserDTO.Website,
				Biography:   registeredUserDTO.Biography,
				Salt:        salt,
				IsConfirmed: false,
				UserType:    model.REGISTERED_USER,
			},
			IsDeleted: false,
		},
		RegisteredUserCategory: model.NONE,
		OfficialDocumentPath:   "",
	}

	if err := handler.RegisteredUserService.CreateRegisteredUser(&registeredUser); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	if err := handler.ClassicUserService.CreateClassicUser(&registeredUser.ClassicUser); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	if err := handler.UserService.CreateUser(&registeredUser.ClassicUser.User); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	fmt.Println(registeredUser.ClassicUser.User.ID)
	confirmationToken := model.ConfirmationToken{
		ID:                uuid.New(),
		ConfirmationToken: uuid.New(),
		UserId:            userId,
		CreatedTime:       time.Now(),
		ExpiredTime:       time.Now().Add(time.Hour * 120),
		IsValid:           true,
	}

	if err := handler.ConfirmationTokenService.CreateConfirmationToken(&confirmationToken); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}

	SendConfirmationMail(registeredUser.ClassicUser.User, confirmationToken.ConfirmationToken)

	profileSettings:= settingsModel.ProfileSettings{
		ID:                uuid.New(),
		UserId:            userId,
		UserVisibility:       settingsModel.PUBLIC_VISIBILITY,
		MessageApprovalType: settingsModel.PUBLIC,
		IsPostTaggable: true,
		IsStoryTaggable: true,
		IsCommentTaggable: true,

	}
	err := handler.ProfileSettingsService.CreateProfileSettings(&profileSettings)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
