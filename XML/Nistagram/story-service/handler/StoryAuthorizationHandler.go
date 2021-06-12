package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
)

type StoryAuthorizationHandler struct {
	UserService                                                  *service.UserService
	Rbac                                                         *gorbac.RBAC
	PermissionCreateSingleStory                                  *gorbac.Permission
	PermissionFindAllPublicStoriesRegisteredUser                 *gorbac.Permission
	PermissionFindAllStoriesForUserRegisteredUser                *gorbac.Permission
	PermissionFindAllFollowingStories                            *gorbac.Permission
	PermissionFindSelectedStoryByIdForRegisteredUsers            *gorbac.Permission
	PermissionFindAllStoriesForLoggedUser                        *gorbac.Permission
	PermissionCreateSingleStoryStoryHighlights                   *gorbac.Permission
	PermissionFindAllSingleStoryStoryHighlightsForStory          *gorbac.Permission
	PermissionFindAllSingleStoryStoryHighlightsForStoryHighlight *gorbac.Permission
	PermissionCreateStoryAlbum                                   *gorbac.Permission
	PermissionFindAllAlbumStoriesForLoggedUser                   *gorbac.Permission
	PermissionFindSelectedStoryAlbumByIdForLoggedUser            *gorbac.Permission
	PermissionFindAllPublicAlbumStoriesRegisteredUser            *gorbac.Permission
	PermissionFindAllFollowingStoryAlbums                        *gorbac.Permission
	PermissionCreateStoryHighlight                               *gorbac.Permission
	PermissionFindAllStoryHighlightsForUser                      *gorbac.Permission
	LogInfo                                                      *logrus.Logger
	LogError                                                     *logrus.Logger
}

func (handler *StoryAuthorizationHandler) CheckCreateSingleStoryPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateSingleStoryPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateSingleStory, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateSingleStoryPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllPublicStoriesRegisteredUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllPublicStoriesRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPublicStoriesRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllPublicStoriesRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllStoriesForUserRegisteredUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoriesForUserRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllStoriesForUserRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoriesForUserRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllFollowingStoriesPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllFollowingStoriesPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllFollowingStories , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllFollowingStoriesPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindSelectedStoryByIdForRegisteredUsersPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindSelectedStoryByIdForRegisteredUsersPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindSelectedStoryByIdForRegisteredUsers , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindSelectedStoryByIdForRegisteredUsersPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllStoriesForLoggedUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoriesForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllStoriesForLoggedUser , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoriesForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckCreateSingleStoryStoryHighlightsPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateSingleStoryStoryHighlightsPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateSingleStoryStoryHighlights , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateSingleStoryStoryHighlightsPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllSingleStoryStoryHighlightsForStoryPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllSingleStoryStoryHighlightsForStoryPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllSingleStoryStoryHighlightsForStory , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllSingleStoryStoryHighlightsForStoryPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllSingleStoryStoryHighlightsForStoryHighlightPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllSingleStoryStoryHighlightsForStoryHighlightPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllSingleStoryStoryHighlightsForStoryHighlight , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllSingleStoryStoryHighlightsForStoryHighlightPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckCreateStoryAlbumPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateStoryAlbumPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateStoryAlbum, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateStoryAlbumPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllAlbumStoriesForLoggedUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllAlbumStoriesForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllAlbumStoriesForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllAlbumStoriesForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindSelectedStoryAlbumByIdForLoggedUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindSelectedStoryAlbumByIdForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindSelectedStoryAlbumByIdForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindSelectedStoryAlbumByIdForLoggedUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllPublicAlbumStoriesRegisteredUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllPublicAlbumStoriesRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPublicAlbumStoriesRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllPublicAlbumStoriesRegisteredUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllFollowingStoryAlbumsPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllFollowingStoryAlbumsPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllFollowingStoryAlbums, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllFollowingStoryAlbumsPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckCreateStoryHighlightPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateStoryHighlightPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateStoryHighlight, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckCreateStoryHighlightPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *StoryAuthorizationHandler) CheckFindAllStoryHighlightsForUserPermission (w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoryHighlightsForUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllStoryHighlightsForUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "StoryAuthorizationHandler",
			"action":    "CheckFindAllStoryHighlightsForUserPermission",
			"timestamp": time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}