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
	gomail "gopkg.in/mail.v2"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

type RegisteredUserHandler struct {
	RegisteredUserService    *service.RegisteredUserService
	UserService              *service.UserService
	ClassicUserService       *service.ClassicUserService
	ConfirmationTokenService *service.ConfirmationTokenService
	Validator                *validator.Validate
	PasswordUtil             *util.PasswordUtil
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//SEDCONFMAIL227
func (handler *RegisteredUserHandler) SendConfirmationMail(user model.User, token uuid.UUID) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "xml.ftn.uns@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", user.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Confirmation mail")

	// Set E-Mail body. You can set plain text or html with text/html
	text := "Dear " + user.FirstName + ",\n\nPlease, click on link in below to confirm your registration on our social network!\n\nhttps://localhost:8081/confirmRegistration/" + token.String() + "/" + user.ID.String() + "\n\n\nBest regards,\nTim25"
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
			"location":   "RegisteredUserHandler",
			"action":   "SEDCONFMAIL227",
			"timestamp":   time.Now().String(),
		}).Error("Failed sending email with confirmation token!")
		panic(err)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "SEDCONFMAIL227",
		"timestamp":   time.Now().String(),
	}).Info("Successfully sended email with confirmation token!")

}

//CRREGUS032
func (handler *RegisteredUserHandler) CreateRegisteredUser(w http.ResponseWriter, r *http.Request) {
	var registeredUserDTO dto.RegisteredUserDTO
	if err := json.NewDecoder(r.Body).Decode(&registeredUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Wrong cast json to RegisteredUserDTO!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&registeredUserDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("RegisteredUserDTO fields doesn't entered in valid format!")
		w.WriteHeader(http.StatusExpectationFailed) //400
		return
	}

	if handler.UserService.FindByUserName(registeredUserDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("User already exist with entered username!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(registeredUserDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("User already exist with entered email!")
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""

	answer:=""
	answerSalt:=""



	validPassword := handler.PasswordUtil.IsValidPassword(registeredUserDTO.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(registeredUserDTO.Password)

		//ANSWER SALT
		answerSalt, answer = handler.PasswordUtil.GeneratePasswordWithSalt(registeredUserDTO.Answer)

	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Password doesn't in valid format!")
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
				Question: registeredUserDTO.Question,
				Answer: answer,
				AnswerSalt: answerSalt,
			},
			IsDeleted: false,
		},
		RegisteredUserCategory: model.NONE,
		OfficialDocumentPath:   "",
	}

	if err := handler.RegisteredUserService.CreateRegisteredUser(&registeredUser); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating registered user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.ClassicUserService.CreateClassicUser(&registeredUser.ClassicUser); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating classic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := handler.UserService.CreateUser(&registeredUser.ClassicUser.User); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating basic user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	tagId := uuid.New()
	var userTag = dto.UserTagFullDTO{
		ID:     tagId,
		Name:   registeredUser.Username,
		UserId: userId,
	}

	reqUrl := fmt.Sprintf("http://%s:%s/create_user_tag_for_registered_user/", os.Getenv("TAG_SERVICE_DOMAIN"), os.Getenv("TAG_SERVICE_PORT"))
	jsonOrders, _ := json.Marshal(userTag)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonOrders))
	resp, err := http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 404 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating user tag for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
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
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating confirmation token for user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	handler.SendConfirmationMail(registeredUser.ClassicUser.User, confirmationToken.ConfirmationToken)

	reqUrl = fmt.Sprintf("http://%s:%s/profile_settings/%s", os.Getenv("SETTINGS_SERVICE_DOMAIN"), os.Getenv("SETTINGS_SERVICE_PORT"), userId)
	jsonOrders, _ = json.Marshal(nil)
	//fmt.Printf("Sending POST req to url %s\nJson being sent:\n", reqUrl)
	//fmt.Println(string(jsonOrders))
	resp, err = http.Post(reqUrl, "application/json", bytes.NewBuffer(jsonOrders))
	if err != nil || resp.StatusCode == 404 {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RegisteredUserHandler",
			"action":   "CRREGUS032",
			"timestamp":   time.Now().String(),
		}).Error("Failed creating profile settings for user!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "RegisteredUserHandler",
		"action":   "CRREGUS032",
		"timestamp":   time.Now().String(),
	}).Info("Successfully created registered user!")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
