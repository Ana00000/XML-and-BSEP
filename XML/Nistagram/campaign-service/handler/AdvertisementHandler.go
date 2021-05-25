package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/dto"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
	"net/http"
	_ "strconv"
)

type AdvertisementHandler struct {
	Service * service.AdvertisementService
}

func (handler *AdvertisementHandler) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	var advertisementDTO dto.AdvertisementDTO
	err := json.NewDecoder(r.Body).Decode(&advertisementDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	advertisement := model.Advertisement{
		ID:                     uuid.UUID{},
		CampaignId:          advertisementDTO.CampaignId,
	}

	err = handler.Service.CreateAdvertisement(&advertisement)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
