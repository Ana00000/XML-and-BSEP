package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	_ "strconv"
	"time"
)

type ClassicUserFollowersHandler struct {
	ClassicUserFollowersService * service.ClassicUserFollowersService
	UserService *service.UserService
	Rbac * gorbac.RBAC
	PermissionFindAllMutualFollowerForUser *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

//FIDALMUTFOLLERFRUS921
func (handler *ClassicUserFollowersHandler) FindAllMutualFollowerForUser(w http.ResponseWriter, r *http.Request) {

	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowersHandler",
			"action":   "FIDALMUTFOLLERFRUS921",
			"timestamp":   time.Now().String(),
		}).Error("User doesn't logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowersHandler",
			"action":   "FIDALMUTFOLLERFRUS921",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(user)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllMutualFollowerForUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ClassicUserFollowersHandler",
			"action":   "FIDALMUTFOLLERFRUS921",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

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
