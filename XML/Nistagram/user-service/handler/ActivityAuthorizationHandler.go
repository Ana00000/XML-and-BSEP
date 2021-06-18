package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type ActivityAuthorizationHandler struct {
	UserService                            *service.UserService
	Rbac                                   *gorbac.RBAC
	PermissionFindAllLikedPostsByUserId    *gorbac.Permission
	PermissionFindAllDislikedPostsByUserId *gorbac.Permission
	LogInfo                                *logrus.Logger
	LogError                               *logrus.Logger
}

func (handler *ActivityAuthorizationHandler) CheckFindAllLikedPostsByUserIdPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityAuthorizationHandler",
			"action": "CheckFindAllLikedPostsByUserIdPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllLikedPostsByUserId, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityAuthorizationHandler",
			"action": "CheckFindAllLikedPostsByUserIdPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *ActivityAuthorizationHandler) CheckFindAllDislikedPostsByUserIdPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityAuthorizationHandler",
			"action": "CheckFindAllDislikedPostsByUserIdPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllLikedPostsByUserId, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location": "ActivityAuthorizationHandler",
			"action": "CheckFindAllDislikedPostsByUserIdPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}
