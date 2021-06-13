package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type ContentAuthorizationHandler struct {
	UserService * service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateSinglePostContent *gorbac.Permission
	PermissionCreatePostAlbumContent *gorbac.Permission
	PermissionCreateSingleStoryContent *gorbac.Permission
	PermissionCreateStoryAlbumContent *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ContentAuthorizationHandler) CheckCreateSinglePostContentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateSinglePostContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateSinglePostContent, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateSinglePostContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *ContentAuthorizationHandler) CheckCreatePostAlbumContentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreatePostAlbumContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostAlbumContent, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreatePostAlbumContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *ContentAuthorizationHandler) CheckCreateSingleStoryContentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateSingleStoryContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateSingleStoryContent, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateSingleStoryContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *ContentAuthorizationHandler) CheckCreateStoryAlbumContentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateStoryAlbumContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateStoryAlbumContent, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "ContentAuthorizationHandler",
			"action":   "CheckCreateStoryAlbumContentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}