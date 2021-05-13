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

type RegisteredUserHandler struct {
	RegisteredUserService * service. RegisteredUserService
	UserService * service.UserService
	ClassicUserService * service.ClassicUserService
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
	text:= "Dear "+user.FirstName+",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttp://localhost:8081/confirmRegistration/"+token.String()+"/"+user.ID.String()+"\n\n\nBest regards,\nTim25"
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

	err := json.NewDecoder(r.Body).Decode(&registeredUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}


	if handler.UserService.FindByUserName(registeredUserDTO.Username) !=  nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(registeredUserDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	var sb strings.Builder
	salt := uuid.New().String()
	sb.WriteString(registeredUserDTO.Password)
	sb.WriteString(salt)
	password := sb.String()
	hash,_ := HashPassword(password)

	gender := model.OTHER
	switch registeredUserDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}
	fmt.Printf(registeredUserDTO.DateOfBirth)
	userId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, registeredUserDTO.DateOfBirth)
	registeredUser := model.RegisteredUser{
		ClassicUser:       model.ClassicUser{
			User:                        model.User{
				ID:          userId,
				Username:    registeredUserDTO.Username,
				Password:    hash,
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
		RegisteredUserCategory:         model.NONE,
		OfficialDocumentPath: "",

	}

	err = handler.RegisteredUserService.CreateRegisteredUser(&registeredUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	err = handler.ClassicUserService.CreateClassicUser(&registeredUser.ClassicUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	err = handler.UserService.CreateUser(&registeredUser.ClassicUser.User)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	fmt.Println(registeredUser.ClassicUser.User.ID)
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
	SendConfirmationMail(registeredUser.ClassicUser.User, confirmationToken.ConfirmationToken)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


