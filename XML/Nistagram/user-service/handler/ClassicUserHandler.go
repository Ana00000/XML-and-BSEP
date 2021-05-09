package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	_ "strconv"
	"strings"
	"time"
)

type ClassicUserHandler struct {
	ClassicUserService * service.ClassicUserService
	UserService * service.UserService
	RegisteredUserService * service.RegisteredUserService
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
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

	userId := uuid.UUID{}
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
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


