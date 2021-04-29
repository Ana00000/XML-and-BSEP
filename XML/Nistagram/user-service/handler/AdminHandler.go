package handler

import (
	"../dto"
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	_ "strconv"
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
			DateOfBirth: adminDTO.DateOfBirth,
			Website:     adminDTO.Website,
			Biography:   adminDTO.Biography,
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

