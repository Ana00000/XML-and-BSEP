package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type LocationAuthorizationHandler struct {
	UserService * service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateLocation *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *LocationAuthorizationHandler) CheckCreateLocationPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationAuthorizationHandler",
			"action":   "CheckCreateLocationPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateLocation, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "LocationAuthorizationHandler",
			"action":   "CheckCreateLocationPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}