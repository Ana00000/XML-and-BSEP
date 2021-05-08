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

type LocationHandler struct {
	Service * service.LocationService
}

func (handler *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var locationDTO dto.LocationDTO
	err := json.NewDecoder(r.Body).Decode(&locationDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	location := model.Location{
		ID:          uuid.UUID{},
		Longitude: locationDTO.Longitude,
		Latitude:      locationDTO.Latitude,
		Country:       locationDTO.Country,
		City:      locationDTO.City,
		StreetName:       locationDTO.StreetName,
		StreetNumber:       locationDTO.StreetNumber,
	}

	err = handler.Service.CreateLocation(&location)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

