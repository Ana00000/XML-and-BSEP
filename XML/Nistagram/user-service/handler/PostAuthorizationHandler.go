package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type PostAuthorizationHandler struct {
	UserService * service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateSinglePost *gorbac.Permission
	PermissionCreatePostAlbum *gorbac.Permission
	PermissionFindAllFollowingPostAlbums *gorbac.Permission
	PermissionFindAllFollowingPosts *gorbac.Permission
	PermissionCreatePostCollection *gorbac.Permission
	PermissionFindAllPostCollectionsForUserRegisteredUser *gorbac.Permission
	PermissionFindAllPostsForLoggedUser *gorbac.Permission
	PermissionFindAllAlbumPostsForLoggedUser *gorbac.Permission
	PermissionCreateComment *gorbac.Permission
	PermissionFindSelectedPostByIdForLoggedUser *gorbac.Permission
	PermissionFindAllCommentsForPost *gorbac.Permission
	PermissionFindAllActivitiesForPost *gorbac.Permission
	PermissionUpdateActivity *gorbac.Permission
	PermissionCreateActivity *gorbac.Permission
	PermissionFindAllPostCollectionPostsForPost *gorbac.Permission
	PermissionCreatePostCollectionPosts *gorbac.Permission
	PermissionFindAllPostsForLocationRegUser *gorbac.Permission
	PermissionFindSelectedPostAlbumByIdForLoggedUser *gorbac.Permission
	PermissionFindAllPostsForTagRegUser *gorbac.Permission
	PermissionFindAllPublicPostsRegisteredUser *gorbac.Permission
	PermissionFindAllPostsForUserRegisteredUser *gorbac.Permission
	PermissionFindAllTagsForPublicAndFollowingPosts *gorbac.Permission
	PermissionFindAllLocationsForPublicAndFollowingPosts *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *PostAuthorizationHandler) CheckCreateSinglePostPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateSinglePostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateSinglePost, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateSinglePostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckCreatePostAlbumPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostAlbumPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostAlbum, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostAlbumPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllFollowingPostAlbumsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllFollowingPostAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllFollowingPostAlbums, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllFollowingPostAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllFollowingPostsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllFollowingPosts, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckCreatePostCollectionPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostCollectionPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostCollection, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostCollectionPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostCollectionsForUserRegisteredUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostCollectionsForUserRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostCollectionsForUserRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostCollectionsForUserRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostsForLoggedUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostsForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllAlbumPostsForLoggedUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllAlbumPostsForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllAlbumPostsForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllAlbumPostsForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckCreateCommentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateComment, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindSelectedPostByIdForLoggedUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindSelectedPostByIdForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindSelectedPostByIdForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindSelectedPostByIdForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllCommentsForPostPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllCommentsForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllCommentsForPost, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllCommentsForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllActivitiesForPostPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllActivitiesForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllActivitiesForPost, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllActivitiesForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckUpdateActivityPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckUpdateActivityPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateActivity, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckUpdateActivityPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckCreateActivityPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateActivityPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateActivity, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreateActivityPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostCollectionPostsForPostPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostCollectionPostsForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostCollectionPostsForPost, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostCollectionPostsForPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostsForLocationRegUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForLocationRegUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostsForLocationRegUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForLocationRegUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckCreatePostCollectionPostsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostCollectionPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostCollectionPosts, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckCreatePostCollectionPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindSelectedPostAlbumByIdForLoggedUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindSelectedPostAlbumByIdForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindSelectedPostAlbumByIdForLoggedUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindSelectedPostAlbumByIdForLoggedUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostsForTagRegUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForTagRegUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostsForTagRegUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForTagRegUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPublicPostsRegisteredUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPublicPostsRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPublicPostsRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPublicPostsRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllPostsForUserRegisteredUserPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForUserRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllPostsForUserRegisteredUser, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllPostsForUserRegisteredUserPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllTagsForPublicAndFollowingPostsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllTagsForPublicAndFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllTagsForPublicAndFollowingPosts, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllTagsForPublicAndFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostAuthorizationHandler) CheckFindAllLocationsForPublicAndFollowingPostsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllLocationsForPublicAndFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllLocationsForPublicAndFollowingPosts, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "PostAuthorizationHandler",
			"action":   "CheckFindAllLocationsForPublicAndFollowingPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}
