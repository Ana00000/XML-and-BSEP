package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "strconv"
	"time"
)

type AdminHandler struct {
	AdminService * service.AdminService
	UserService *service.UserService
	Validator *validator.Validate
	PasswordUtil *util.PasswordUtil
}

func (handler *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminDTO dto.AdminDTO
	if err := json.NewDecoder(r.Body).Decode(&adminDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.Validator.Struct(&adminDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if handler.UserService.FindByUserName(adminDTO.Username) != nil {
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if handler.UserService.FindByEmail(adminDTO.Email) != nil {
		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(adminDTO.Password)

	if validPassword {
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(adminDTO.Password)
	}else {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch adminDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	adminId := uuid.New()
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,adminDTO.DateOfBirth)
	admin := model.Admin{
		User : model.User{
			ID:          adminId,
			Username:    adminDTO.Username,
			Password:    password,
			Email:       adminDTO.Email,
			PhoneNumber: adminDTO.PhoneNumber,
			FirstName:   adminDTO.FirstName,
			LastName:    adminDTO.LastName,
			Gender:      gender,
			DateOfBirth: dateOfBirth,
			Website:     adminDTO.Website,
			Biography:   adminDTO.Biography,
			Salt: salt,
			IsConfirmed: true,
			UserType: model.ADMIN,
		},
	}

	err := handler.AdminService.CreateAdmin(&admin)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

