package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	_ "strconv"
	"strings"
	"time"
)


type UserHandler struct {
	UserService * service.UserService
	AdminService * service.AdminService
	ClassicUserService * service.ClassicUserService
	AgentService * service.AgentService
	RegisteredUserService * service.RegisteredUserService
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (handler *UserHandler)FindAllUsers(w http.ResponseWriter, r *http.Request){
	var users []model.User
	users = handler.UserService.FindAllUsers()
	usersJson, _ := json.Marshal(users)
	if usersJson != nil {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJson)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func CreateToken(userName string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userName
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (handler *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("USAO")
	var logInUserDTO dto.LogInUserDTO
	err := json.NewDecoder(r.Body).Decode(&logInUserDTO)
	if err != nil {
		fmt.Println("USAOLOSE")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user = handler.UserService.FindByUserName(logInUserDTO.Username)

	if !user.IsConfirmed {
		fmt.Println("GRESKA")
	}
	fmt.Println("PROSAO")
	if user == nil || !user.IsConfirmed {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var sb strings.Builder
	salt := user.Salt
	sb.WriteString(logInUserDTO.Password)
	sb.WriteString(salt)
	password := sb.String()

	if !CheckPasswordHash(password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	tokenJson, _ := json.Marshal(token)
	w.Write(tokenJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

func (handler *UserHandler) UpdateUserProfileInfo(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserUpdateProfileInfoDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.UpdateUserProfileInfo(&userDTO)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}


	if userDTO.UserType == "ADMIN" {
		err = handler.AdminService.UpdateAdminProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}
	} else{
		err = handler.ClassicUserService.UpdateClassicUserProfileInfo(&userDTO)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusExpectationFailed)
		}

		if userDTO.UserType == "AGENT" {
			err = handler.AgentService.UpdateAgentProfileInfo(&userDTO)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusExpectationFailed)
			}
		}else{
			err = handler.RegisteredUserService.UpdateRegisteredUserProfileInfo(&userDTO)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusExpectationFailed)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}