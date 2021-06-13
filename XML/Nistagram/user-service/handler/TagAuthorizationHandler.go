package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"net/http"
	"time"
)

type TagAuthorizationHandler struct {
	UserService * service.UserService
	Rbac * gorbac.RBAC
	PermissionCreateCommentTagComments *gorbac.Permission
	PermissionFindAllTaggableUsersPost *gorbac.Permission
	PermissionCreateTag *gorbac.Permission
	PermissionCreatePostTagPosts *gorbac.Permission
	PermissionCreatePostAlbumTagPostAlbums *gorbac.Permission
	PermissionCreateStoryTagStories *gorbac.Permission
	PermissionCreateStoryAlbumTagStoryAlbums *gorbac.Permission
	PermissionFindAllTaggableUsersStory *gorbac.Permission
	PermissionFindAllCommentTagCommentsForComment *gorbac.Permission
	PermissionFindAllTaggableUsersComment *gorbac.Permission
	PermissionFindAllHashTags *gorbac.Permission
	LogInfo *logrus.Logger
	LogError *logrus.Logger
}

func (handler *TagAuthorizationHandler) CheckCreateCommentTagCommentsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateCommentTagCommentsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateCommentTagComments, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateCommentTagCommentsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckFindAllTaggableUsersPostPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllTaggableUsersPost, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersPostPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckCreateTagPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateTagPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateTag, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateTagPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckCreatePostTagPostsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreatePostTagPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostTagPosts, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreatePostTagPostsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckCreatePostAlbumTagPostAlbumsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreatePostAlbumTagPostAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreatePostAlbumTagPostAlbums, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreatePostAlbumTagPostAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckCreateStoryTagStoriesPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateStoryTagStoriesPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateStoryTagStories, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateStoryTagStoriesPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckCreateStoryAlbumTagStoryAlbumsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateStoryAlbumTagStoryAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionCreateStoryAlbumTagStoryAlbums, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckCreateStoryAlbumTagStoryAlbumsPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckFindAllTaggableUsersStoryPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersStoryPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllTaggableUsersStory, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersStoryPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckFindAllCommentTagCommentsForCommentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllCommentTagCommentsForCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllCommentTagCommentsForComment, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllCommentTagCommentsForCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckFindAllTaggableUsersCommentPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllTaggableUsersComment, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllTaggableUsersCommentPermission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TagAuthorizationHandler) CheckFindAllHashTagsPermission(w http.ResponseWriter, r *http.Request) {
	userName, err := getUserNameFromJWT(r)
	if err!=nil	{
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllHashTags Permission",
			"timestamp":   time.Now().String(),
		}).Error("Failed finding user from jwt token!")
		w.WriteHeader(http.StatusFailedDependency)
		return
	}
	var userSigned = handler.UserService.FindByUserName(userName)
	var userRole = getRoleByUser(userSigned)

	if !handler.Rbac.IsGranted(userRole, *handler.PermissionFindAllHashTags , nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status": "failure",
			"location":   "TagAuthorizationHandler",
			"action":   "CheckFindAllHashTags Permission",
			"timestamp":   time.Now().String(),
		}).Error("Forbidden method for logged in user!")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}