package handler

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
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

	locationIDJson, _ := json.Marshal(location.ID)
	w.Write(locationIDJson)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LocationHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var location = handler.Service.FindByID(uuid.MustParse(id))
	if  location == nil {
		fmt.Println("Location not found")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	locationJson, _ := json.Marshal(location)
	w.Write(locationJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}