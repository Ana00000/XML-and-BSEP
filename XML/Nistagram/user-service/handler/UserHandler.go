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

const(
	SALT_BYTE_SIZE = 24
)

type UserHandler struct {
	Service * service.UserService
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

//for login
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)

	var sb strings.Builder
	salt := uuid.New().String()
	sb.WriteString(userDTO.Password)
	sb.WriteString(salt)
	password := sb.String()
	hash,_ := HashPassword(password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,userDTO.DateOfBirth)
	user := model.User{
		ID:               uuid.UUID{},
		Username:         userDTO.Username,
		Password:         hash,
		Email:            userDTO.Email,
		PhoneNumber:      userDTO.PhoneNumber,
		FirstName:        userDTO.FirstName,
		LastName:         userDTO.LastName,
		Gender:           userDTO.Gender,
		DateOfBirth:      dateOfBirth,
		Website:          userDTO.Website,
		Biography:        userDTO.Biography,
		Salt: 			  salt,
	}

	err = handler.Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

