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
	Service * service.AdminService
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

	salt,password := handler.PasswordUtil.GeneratePasswordWithSalt(adminDTO.Password)

	gender := model.OTHER
	switch adminDTO.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	fmt.Printf(adminDTO.DateOfBirth)
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,adminDTO.DateOfBirth)
	admin := model.Admin{
		User : model.User{
			ID:          uuid.UUID{},
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

	err := handler.Service.CreateAdmin(&admin)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

