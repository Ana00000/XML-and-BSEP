package handler

import (
	"encoding/json"
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
	Service * service.UserService
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (handler *UserHandler)FindAllUsers(w http.ResponseWriter, r *http.Request){
	var users []model.User
	users = handler.Service.FindAllUsers()
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
	var logInUserDTO dto.LogInUserDTO
	err := json.NewDecoder(r.Body).Decode(&logInUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user = handler.Service.FindByUserName(logInUserDTO.Username)

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

	w.WriteHeader(http.StatusOK)
	tokenJson, _ := json.Marshal(token)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)

}