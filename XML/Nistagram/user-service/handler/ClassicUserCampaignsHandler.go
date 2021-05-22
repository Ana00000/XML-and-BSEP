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
)

type ClassicUserCampaignsHandler struct {
	Service * service.ClassicUserCampaignsService
}

func (handler *ClassicUserCampaignsHandler) CreateClassicUserCampaigns(w http.ResponseWriter, r *http.Request) {
	var classicUserCampaignsDTO dto.ClassicUserCampaignsDTO
	err := json.NewDecoder(r.Body).Decode(&classicUserCampaignsDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	classicUserCampaigns := model.ClassicUserCampaigns{
		ID:               uuid.UUID{},
		ClassicUserId: classicUserCampaignsDTO.ClassicUserId,
		CampaignId:       classicUserCampaignsDTO.CampaignId,
	}

	err = handler.Service.CreateClassicUserCampaigns(&classicUserCampaigns)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}