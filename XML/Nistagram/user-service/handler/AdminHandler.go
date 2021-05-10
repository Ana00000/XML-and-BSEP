package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type AdminHandler struct {
	Service * service.AdminService
}

func (handler *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminDTO dto.AdminDTO
	err := json.NewDecoder(r.Body).Decode(&adminDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth,_ :=time.Parse(layout,adminDTO.DateOfBirth)
	admin := model.Admin{
		User : model.User{
			ID:          uuid.UUID{},
			Username:    adminDTO.Username,
			Password:    adminDTO.Password,
			Email:       adminDTO.Email,
			PhoneNumber: adminDTO.PhoneNumber,
			FirstName:   adminDTO.FirstName,
			LastName:    adminDTO.LastName,
			Gender:      adminDTO.Gender,
			DateOfBirth: dateOfBirth,
			Website:     adminDTO.Website,
			Biography:   adminDTO.Biography,
			IsConfirmed: true,
			UserType: model.ADMIN, //SET VALUE
		},
	}

	err = handler.Service.CreateAdmin(&admin)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

