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
	PermissionUpdateStatusFollowRequest *gorbac.Permission
	PermissionFindFollowRequestById *gorbac.Permission
	PermissionFindAllPendingFollowerRequestsForUser *gorbac.Permission
	PermissionCreateVerificationRequest *gorbac.Permission
	PermissionUpdateStatusVerificationRequest *gorbac.Permission
	PermissionFindVerificationRequestById *gorbac.Permission
	PermissionFindAllPendingVerificationRequests *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}


//FOLLOW REQUEST

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

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateStatusFollowRequest, nil) {
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

func (handler *RequestsAuthorizationHandler) CheckFindFollowerRequestByIdPermission(w http.ResponseWriter, r *http.Request) {
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

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindFollowRequestById, nil) {
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


//VERIFICATION REQUEST

func (handler *RequestsAuthorizationHandler) CheckCreateVerificationRequestPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckCreateVerificationRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateVerificationRequest, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckCreateVerificationRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckUpdateStatusVerificationRequestPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckUpdateStatusVerificationRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateStatusVerificationRequest, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckUpdateStatusVerificationRequestPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckFindVerificationRequestByIdPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindVerificationRequestByIdPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindVerificationRequestById, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindVerificationRequestByIdPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RequestsAuthorizationHandler) CheckFindAllPendingVerificationRequestsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindAllPendingVerificationRequestsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPendingVerificationRequests, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "RequestsAuthorizationHandler",
			"action":   "CheckFindAllPendingVerificationRequestsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

