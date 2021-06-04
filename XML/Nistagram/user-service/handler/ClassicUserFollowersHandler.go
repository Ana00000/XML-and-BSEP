package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
)

type ClassicUserFollowersHandler struct {
	ClassicUserFollowersService *service.ClassicUserFollowersService
}

func (handler *ClassicUserFollowersHandler) FindAllMutualFollowerForUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var mutuals = handler.ClassicUserFollowersService.FindAllMutualFollowerForUser(uuid.MustParse(id))

	mutualsJson, _ := json.Marshal(mutuals)
	w.Write(mutualsJson)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
