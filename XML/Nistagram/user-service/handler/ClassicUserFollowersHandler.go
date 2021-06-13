package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type ClassicUserFollowersHandler struct {
	ClassicUserFollowersService * service.ClassicUserFollowersService
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//FIDALMUTFOLLERFRUS921
func (handler *ClassicUserFollowersHandler) FindAllMutualFollowerForUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	id := r.URL.Query().Get("id")

	var mutuals = handler.ClassicUserFollowersService.FindAllMutualFollowerForUser(uuid.MustParse(id))

	handler.LogInfo.WithFields(logrus.Fields{
		"status": "success",
		"location":   "ClassicUserFollowersHandler",
		"action":   "FIDALMUTFOLLERFRUS921",
		"timestamp":   time.Now().String(),
	}).Info("Successfully founded all mutual followers for user!")

	mutualsJson, _ := json.Marshal(mutuals)
	w.Write(mutualsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
