package handler

import (
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
)

type StoryAuthorizationHandler struct {
	StoryService                                                 *service.StoryService
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
