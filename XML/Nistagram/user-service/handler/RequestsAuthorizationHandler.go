package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type RequestsAuthorizationHandler struct {
	UserService * service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateFollowRequest *gorbac.Permission
	PermissionRejectFollowRequest *gorbac.Permission
	PermissionFindRequestById *gorbac.Permission
	PermissionFindAllPendingFollowerRequestsForUser *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *RequestsAuthorizationHandler) CheckCreateFollowRequestPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckCreateFollowRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateFollowRequest, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckCreateFollowRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckUpdateStatusFollowRequestPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckRejectFollowRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionRejectFollowRequest, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckRejectFollowRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckFindRequestByIdPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindRequestByIdPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindRequestById, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindRequestByIdPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckFindAllPendingFollowerRequestsForUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindAllPendingFollowerRequestsForUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPendingFollowerRequestsForUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindAllPendingFollowerRequestsForUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}
