package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/util"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.ClassicUser{}, &model.RegisteredUser{}, &model.Admin{}, &model.Agent{},&model.ClassicUserFollowers{},&model.ClassicUserFollowings{}, &model.ClassicUserCampaigns{}, &model.ConfirmationToken{}, &model.RecoveryPasswordToken{}, &model.ClassicUserCloseFriends{})
	return db
}

func initDSN() string {

	host := "localhost"
	user := "postgres"
	password := "root"
	dbname := "nistagram-db"
	dbport := "5432"
	if os.Getenv("DBHOST") != "" && os.Getenv("USER") != "" && os.Getenv("PASSWORD") != "" &&
		os.Getenv("DBNAME") != "" && os.Getenv("DBPORT") != "" {
		host = os.Getenv("DBHOST")
		user = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname = os.Getenv("DBNAME")
		dbport = os.Getenv("DBPORT")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)

	return dsn
}


func initPasswordUtil() *util.PasswordUtil{
	return &util.PasswordUtil { }
}

//USER
func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initUserHandler(permissionFindUserByID *gorbac.Permission,LogInfo *logrus.Logger,LogError *logrus.Logger,RecoveryPasswordTokenService *service.RecoveryPasswordTokenService,UserService *service.UserService,AdminService *service.AdminService, ClassicUserService *service.ClassicUserService, RegisteredUserService *service.RegisteredUserService, AgentService *service.AgentService, rbac *gorbac.RBAC, permissionFindAllUsers *gorbac.Permission, permissionUpdateUserInfo *gorbac.Permission, validator *validator.Validate, passwordUtil *util.PasswordUtil ) *handler.UserHandler{
	return &handler.UserHandler{
		UserService:                  UserService,
		AdminService:                 AdminService,
		ClassicUserService:           ClassicUserService,
		AgentService:                 AgentService,
		Rbac:                         rbac,
		PermissionFindAllUsers:       permissionFindAllUsers,
		RegisteredUserService:        RegisteredUserService,
		RecoveryPasswordTokenService: RecoveryPasswordTokenService,
		PermissionFindUserByID:       permissionFindUserByID,
		PermissionUpdateUserInfo:     permissionUpdateUserInfo,
		Validator:                    validator,
		PasswordUtil:                 passwordUtil,
		LogInfo:                      LogInfo,
		LogError:                     LogError,
	}
}
//SETTINGS



//ADMIN
func initAdminRepo(database *gorm.DB) *repository.AdminRepository{
	return &repository.AdminRepository { Database: database }
}

func initAdminService(repo *repository.AdminRepository) *service.AdminService{
	return &service.AdminService { Repo: repo }
}

func initAdminHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,adminService *service.AdminService, userService *service.UserService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.AdminHandler{
	return &handler.AdminHandler{
		AdminService: adminService,
		UserService:  userService,
		Validator:    validator,
		LogInfo:      LogInfo,
		LogError:     LogError,
		PasswordUtil: passwordUtil,
	}
}

//CLASSIC USER
func initClassicUserRepo(database *gorm.DB) *repository.ClassicUserRepository{
	return &repository.ClassicUserRepository { Database: database }
}

func initClassicUserService(repo *repository.ClassicUserRepository) *service.ClassicUserService{
	return &service.ClassicUserService { Repo: repo }
}

func initClassicUserHandler(userService *service.UserService,permissionFindAllUsersButLoggedIn *gorbac.Permission,rbac *gorbac.RBAC, LogInfo *logrus.Logger,LogError *logrus.Logger,classicUserService *service.ClassicUserService, classicUserFollowingsService *service.ClassicUserFollowingsService) *handler.ClassicUserHandler{
	return &handler.ClassicUserHandler{
		ClassicUserService:                classicUserService,
		ClassicUserFollowingsService:      classicUserFollowingsService,
		UserService:                       userService,
		Rbac:                              rbac,
		PermissionFindAllUsersButLoggedIn: permissionFindAllUsersButLoggedIn,
		LogInfo:                           LogInfo,
		LogError:                          LogError,
	}
}

//REGISTERED USER
func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}

func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}

func initRegisteredUserHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,registeredUserService *service.RegisteredUserService, userService *service.UserService, classicUserService *service.ClassicUserService,  confirmationTokenService *service.ConfirmationTokenService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler{
		registeredUserService,
		userService,
		classicUserService ,
		confirmationTokenService,
		validator,
		passwordUtil,
		LogInfo,
		LogError,
	}

}

func initAgentRepo(database *gorm.DB) *repository.AgentRepository{
	return &repository.AgentRepository { Database: database }
}

func initAgentService(repo *repository.AgentRepository) *service.AgentService{
	return &service.AgentService { Repo: repo }
}

func initAgentHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,agentService *service.AgentService, userService *service.UserService, classicUserService *service.ClassicUserService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.AgentHandler{
	return &handler.AgentHandler{
		AgentService: agentService,
		UserService: userService,
		ClassicUserService: classicUserService,
		Validator: validator,
		PasswordUtil: passwordUtil,
		LogInfo:      LogInfo,
		LogError:     LogError,
	}
}

func initClassicUserCampaignsRepo(database *gorm.DB) *repository.ClassicUserCampaignsRepository{
	return &repository.ClassicUserCampaignsRepository { Database: database }
}

func initClassicUserFollowersRepo(database *gorm.DB) *repository.ClassicUserFollowersRepository{
	return &repository.ClassicUserFollowersRepository { Database: database }
}

func initClassicUserFollowingsRepo(database *gorm.DB) *repository.ClassicUserFollowingsRepository{
	return &repository.ClassicUserFollowingsRepository { Database: database }
}

func initConfirmationTokenRepo(database *gorm.DB) *repository.ConfirmationTokenRepository{
	return &repository.ConfirmationTokenRepository { Database: database }
}

func initRecoveryPasswordTokenRepo(database *gorm.DB) *repository.RecoveryPasswordTokenRepository{
	return &repository.RecoveryPasswordTokenRepository { Database: database }
}



func initClassicUserCampaignsService(repo *repository.ClassicUserCampaignsRepository) *service.ClassicUserCampaignsService{
	return &service.ClassicUserCampaignsService { Repo: repo }
}

func initConfirmationTokenService(repo *repository.ConfirmationTokenRepository) *service.ConfirmationTokenService{
	return &service.ConfirmationTokenService { Repo: repo }
}

func initClassicUserFollowingsService(repo *repository.ClassicUserFollowingsRepository) *service.ClassicUserFollowingsService{
	return &service.ClassicUserFollowingsService { Repo: repo }
}

func initClassicUserFollowersService(repo *repository.ClassicUserFollowersRepository) *service.ClassicUserFollowersService{
	return &service.ClassicUserFollowersService { Repo: repo }
}

func initRecoveryPasswordTokenService(repo *repository.RecoveryPasswordTokenRepository) *service.RecoveryPasswordTokenService{
	return &service.RecoveryPasswordTokenService { Repo: repo }
}

func initClassicUserCampaignsHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.ClassicUserCampaignsService) *handler.ClassicUserCampaignsHandler{
	return &handler.ClassicUserCampaignsHandler {
		Service: service,
		LogInfo:      LogInfo,
		LogError:     LogError,
	}
}

func initClassicUserFollowersHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,rbac *gorbac.RBAC, permissionFindAllMutualFollowerForUser *gorbac.Permission,userService *service.UserService, service *service.ClassicUserFollowersService) *handler.ClassicUserFollowersHandler{
	return &handler.ClassicUserFollowersHandler{
		ClassicUserFollowersService:            service,
		UserService:                            userService,
		Rbac:                                   rbac,
		PermissionFindAllMutualFollowerForUser: permissionFindAllMutualFollowerForUser,
		LogInfo:                                LogInfo,
		LogError:                               LogError,
	}
}

func initClassicUserFollowingsHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,rbac *gorbac.RBAC,permissionAcceptFollowerRequest *gorbac.Permission, permissionCreateClassicUserFollowing *gorbac.Permission, userService *service.UserService,classicUserFollowings *service.ClassicUserFollowingsService, classicUserFollowersService *service.ClassicUserFollowersService) *handler.ClassicUserFollowingsHandler{
	return &handler.ClassicUserFollowingsHandler{
		ClassicUserFollowingsService:         classicUserFollowings,
		ClassicUserFollowersService:          classicUserFollowersService,
		UserService:                          userService,
		Rbac:                                 rbac,
		PermissionCreateClassicUserFollowing: permissionCreateClassicUserFollowing,
		PermissionAcceptFollowerRequest:      permissionAcceptFollowerRequest,
		LogInfo:                              LogInfo,
		LogError:                             LogError,
	}
}

func initRecoveryPasswordTokenHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,recoveryPasswordTokenService *service.RecoveryPasswordTokenService, userService *service.UserService, validator *validator.Validate) *handler.RecoveryPasswordTokenHandler {
	return &handler.RecoveryPasswordTokenHandler{
		RecoveryPasswordTokenService: recoveryPasswordTokenService,
		UserService:                  userService,
		Validator:                    validator,
		LogInfo:                      LogInfo,
		LogError:                     LogError,
	}
}

func initConfirmationTokenHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,confirmationTokenService *service.ConfirmationTokenService, userService *service.UserService, registeredUserService *service.RegisteredUserService, classicUserService *service.ClassicUserService) *handler.ConfirmationTokenHandler{
	return &handler.ConfirmationTokenHandler{
		ConfirmationTokenService: confirmationTokenService,
		ClassicUserService:       classicUserService,
		RegisteredUserService:    registeredUserService,
		UserService:              userService,
		LogInfo:                  LogInfo,
		LogError:                 LogError,
	}
}

//CLASSIC USER CLOSE FRIENDS
func initClassicUserCloseFriendsRepo(database *gorm.DB) *repository.ClassicUserCloseFriendsRepository{
	return &repository.ClassicUserCloseFriendsRepository { Database: database }
}

func initClassicUserCloseFriendsService(repo *repository.ClassicUserCloseFriendsRepository) *service.ClassicUserCloseFriendsService{
	return &service.ClassicUserCloseFriendsService { Repo: repo }
}

func initClassicUserCloseFriendsHandler(userService *service.UserService,rbac *gorbac.RBAC, permissionCreateClassicUserCloseFriend *gorbac.Permission, LogInfo *logrus.Logger,LogError *logrus.Logger,classicUserCloseFirendsService *service.ClassicUserCloseFriendsService, classicUserFollowersService *service.ClassicUserFollowersService ) *handler.ClassicUserCloseFriendsHandler{
	return &handler.ClassicUserCloseFriendsHandler{
		ClassicUserCloseFriendsService:         classicUserCloseFirendsService,
		ClassicUserFollowersService:            classicUserFollowersService,
		Rbac:                                   rbac,
		PermissionCreateClassicUserCloseFriend: permissionCreateClassicUserCloseFriend,
		UserService:                            userService,
		LogInfo:                                LogInfo,
		LogError:                               LogError,
	}
}

func initLocationAuthorizationHandler(rbac *gorbac.RBAC, permissionCreateLocation *gorbac.Permission,
	LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService) *handler.LocationAuthorizationHandler{
	return &handler.LocationAuthorizationHandler{
		UserService:                                   userService,
		Rbac:                                          rbac,
		PermissionCreateLocation:             		   permissionCreateLocation,
		LogInfo:                                       LogInfo,
		LogError:                                      LogError,
	}
}

func initContentAuthorizationHandler(rbac *gorbac.RBAC, permissionCreateSinglePostContent *gorbac.Permission,permissionCreatePostAlbumContent *gorbac.Permission,
	permissionCreateSingleStoryContent *gorbac.Permission,permissionCreateStoryAlbumContent *gorbac.Permission,
	LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService) *handler.ContentAuthorizationHandler{
	return &handler.ContentAuthorizationHandler{
		UserService:                                   userService,
		Rbac:                                          rbac,
		PermissionCreateSinglePostContent:             permissionCreateSinglePostContent,
		PermissionCreatePostAlbumContent:              permissionCreatePostAlbumContent,
		PermissionCreateSingleStoryContent:            permissionCreateSingleStoryContent,
		PermissionCreateStoryAlbumContent:             permissionCreateStoryAlbumContent,
		LogInfo:                                       LogInfo,
		LogError:                                      LogError,
	}
}

func initTagAuthorizationHandler(rbac *gorbac.RBAC, permissionCreateCommentTagComments *gorbac.Permission,permissionFindAllHashTags *gorbac.Permission,
	permissionCreateStoryAlbumTagStoryAlbums *gorbac.Permission,permissionFindAllTaggableUsersStory *gorbac.Permission,permissionFindAllTaggableUsersComment *gorbac.Permission,
	permissionCreatePostTagPosts *gorbac.Permission,permissionCreatePostAlbumTagPostAlbums *gorbac.Permission,permissionFindAllCommentTagCommentsForComment *gorbac.Permission,
	permissionCreateTag *gorbac.Permission,permissionFindAllTaggableUsersPost *gorbac.Permission, LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService) *handler.TagAuthorizationHandler{
	return &handler.TagAuthorizationHandler{
		UserService:                                   userService,
		Rbac:                                          rbac,
		PermissionCreateCommentTagComments:            permissionCreateCommentTagComments,
		PermissionFindAllTaggableUsersPost:            permissionFindAllTaggableUsersPost,
		PermissionCreateTag:                           permissionCreateTag,
		PermissionCreatePostTagPosts:                  permissionCreatePostTagPosts,
		PermissionCreatePostAlbumTagPostAlbums:        permissionCreatePostAlbumTagPostAlbums,
		PermissionCreateStoryTagStories:               permissionCreateCommentTagComments,
		PermissionCreateStoryAlbumTagStoryAlbums:      permissionCreateStoryAlbumTagStoryAlbums,
		PermissionFindAllTaggableUsersStory:           permissionFindAllTaggableUsersStory,
		PermissionFindAllCommentTagCommentsForComment: permissionFindAllCommentTagCommentsForComment,
		PermissionFindAllTaggableUsersComment:         permissionFindAllTaggableUsersComment,
		PermissionFindAllHashTags:                     permissionFindAllHashTags,
		LogInfo:                                       LogInfo,
		LogError:                                      LogError,
	}
}


func initPostAuthorizationHandler(rbac *gorbac.RBAC, LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService,permissionCreateSinglePost *gorbac.Permission,
	permissionCreatePostAlbum *gorbac.Permission,permissionFindAllFollowingPostAlbums *gorbac.Permission,permissionFindAllFollowingPosts *gorbac.Permission,permissionCreatePostCollection *gorbac.Permission,
	permissionFindAllPostCollectionsForUserRegisteredUser *gorbac.Permission,permissionFindAllPostsForLoggedUser *gorbac.Permission,
	permissionFindAllAlbumPostsForLoggedUser *gorbac.Permission,permissionCreateComment *gorbac.Permission,permissionFindSelectedPostByIdForLoggedUser *gorbac.Permission,
	permissionFindAllCommentsForPost *gorbac.Permission,permissionFindAllActivitiesForPost *gorbac.Permission,permissionUpdateActivity *gorbac.Permission,permissionCreateActivity *gorbac.Permission,
	permissionFindAllPostCollectionPostsForPost *gorbac.Permission,permissionCreatePostCollectionPosts *gorbac.Permission,permissionFindAllPostsForLocationRegUser *gorbac.Permission,
	permissionFindSelectedPostAlbumByIdForLoggedUser *gorbac.Permission,permissionFindAllPostsForTagRegUser *gorbac.Permission,permissionFindAllPublicPostsRegisteredUser *gorbac.Permission,
	permissionFindAllPostsForUserRegisteredUser *gorbac.Permission,permissionFindAllTagsForPublicAndFollowingPosts *gorbac.Permission,permissionFindAllLocationsForPublicAndFollowingPosts *gorbac.Permission) *handler.PostAuthorizationHandler{
	return &handler.PostAuthorizationHandler{
		UserService:                          userService,
		Rbac:                                 rbac,
		PermissionCreateSinglePost:           permissionCreateSinglePost,
		PermissionCreatePostAlbum:            permissionCreatePostAlbum,
		PermissionFindAllFollowingPostAlbums: permissionFindAllFollowingPostAlbums,
		PermissionFindAllFollowingPosts:      permissionFindAllFollowingPosts,
		PermissionCreatePostCollection:       permissionCreatePostCollection,
		PermissionFindAllPostCollectionsForUserRegisteredUser: permissionFindAllPostCollectionsForUserRegisteredUser,
		PermissionFindAllPostsForLoggedUser:                   permissionFindAllPostsForLoggedUser,
		PermissionFindAllAlbumPostsForLoggedUser:              permissionFindAllAlbumPostsForLoggedUser,
		PermissionCreateComment:                               permissionCreateComment,
		PermissionFindSelectedPostByIdForLoggedUser:           permissionFindSelectedPostByIdForLoggedUser,
		PermissionFindAllCommentsForPost:                      permissionFindAllCommentsForPost,
		PermissionFindAllActivitiesForPost:                    permissionFindAllActivitiesForPost,
		PermissionUpdateActivity:                              permissionUpdateActivity,
		PermissionCreateActivity:                              permissionCreateActivity,
		PermissionFindAllPostCollectionPostsForPost:           permissionFindAllPostCollectionPostsForPost,
		PermissionCreatePostCollectionPosts:                   permissionCreatePostCollectionPosts,
		PermissionFindAllPostsForLocationRegUser:              permissionFindAllPostsForLocationRegUser,
		PermissionFindSelectedPostAlbumByIdForLoggedUser:      permissionFindSelectedPostAlbumByIdForLoggedUser,
		PermissionFindAllPostsForTagRegUser:                   permissionFindAllPostsForTagRegUser,
		PermissionFindAllPublicPostsRegisteredUser:            permissionFindAllPublicPostsRegisteredUser,
		PermissionFindAllPostsForUserRegisteredUser:           permissionFindAllPostsForUserRegisteredUser,
		PermissionFindAllTagsForPublicAndFollowingPosts:       permissionFindAllTagsForPublicAndFollowingPosts,
		PermissionFindAllLocationsForPublicAndFollowingPosts:  permissionFindAllLocationsForPublicAndFollowingPosts,
		LogInfo:  LogInfo,
		LogError: LogError,
	}
}

func initSettingsAuthorizationHandler(rbac *gorbac.RBAC, LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService,permissionBlockUser *gorbac.Permission, permissionMuteUser *gorbac.Permission) *handler.SettingsAuthorizationHandler{
	return &handler.SettingsAuthorizationHandler{
		UserService:         userService,
		Rbac:                rbac,
		PermissionMuteUser:  permissionMuteUser,
		PermissionBlockUser: permissionBlockUser,
		LogInfo:             LogInfo,
		LogError:            LogError,
	}
}

func initRequestsAuthorizationHandler(rbac *gorbac.RBAC, permissionCreateFollowRequest *gorbac.Permission, permissionRejectFollowRequest *gorbac.Permission,  permissionFindRequestById *gorbac.Permission,  permissionFindAllPendingFollowerRequestsForUser *gorbac.Permission, LogInfo *logrus.Logger,LogError *logrus.Logger,userService *service.UserService) *handler.RequestsAuthorizationHandler{
	return &handler.RequestsAuthorizationHandler{
		UserService:                                   userService,
		Rbac:                                          rbac,
		PermissionCreateFollowRequest:				   permissionCreateFollowRequest,
		PermissionRejectFollowRequest:				   permissionRejectFollowRequest,
		PermissionFindRequestById:					   permissionFindRequestById,
		PermissionFindAllPendingFollowerRequestsForUser: permissionFindAllPendingFollowerRequestsForUser,
		LogInfo:                                       LogInfo,
		LogError:                                      LogError,
	}
}


func initStoryAuthorizationHandler(userService *service.UserService, rbac *gorbac.RBAC, permissionCreateSingleStory *gorbac.Permission, permissionFindAllPublicStoriesRegisteredUser *gorbac.Permission, permissionFindAllStoriesForUserRegisteredUser *gorbac.Permission, permissionFindAllFollowingStories *gorbac.Permission, permissionFindSelectedStoryByIdForRegisteredUsers *gorbac.Permission, permissionFindAllStoriesForLoggedUser *gorbac.Permission, permissionCreateSingleStoryStoryHighlights *gorbac.Permission, permissionFindAllSingleStoryStoryHighlightsForStory *gorbac.Permission, permissionFindAllSingleStoryStoryHighlightsForStoryHighlight *gorbac.Permission, permissionCreateStoryAlbum *gorbac.Permission, permissionFindAllAlbumStoriesForLoggedUser *gorbac.Permission, permissionFindSelectedStoryAlbumByIdForLoggedUser *gorbac.Permission, permissionFindAllPublicAlbumStoriesRegisteredUser *gorbac.Permission, permissionFindAllFollowingStoryAlbums *gorbac.Permission, permissionCreateStoryHighlight *gorbac.Permission, permissionFindAllStoryHighlightsForUser *gorbac.Permission, LogInfo *logrus.Logger,LogError *logrus.Logger) *handler.StoryAuthorizationHandler{
	return &handler.StoryAuthorizationHandler{
		UserService:                                   userService,
		Rbac:                                          rbac,
		PermissionCreateSingleStory: permissionCreateSingleStory,
		PermissionFindAllPublicStoriesRegisteredUser: permissionFindAllPublicStoriesRegisteredUser,
		PermissionFindAllStoriesForUserRegisteredUser: permissionFindAllStoriesForUserRegisteredUser,
		PermissionFindAllFollowingStories: permissionFindAllFollowingStories,
		PermissionFindSelectedStoryByIdForRegisteredUsers: permissionFindSelectedStoryByIdForRegisteredUsers,
		PermissionFindAllStoriesForLoggedUser: permissionFindAllStoriesForLoggedUser,
		PermissionCreateSingleStoryStoryHighlights: permissionCreateSingleStoryStoryHighlights,
		PermissionFindAllSingleStoryStoryHighlightsForStory: permissionFindAllSingleStoryStoryHighlightsForStory,
		PermissionFindAllSingleStoryStoryHighlightsForStoryHighlight: permissionFindAllSingleStoryStoryHighlightsForStoryHighlight,
		PermissionCreateStoryAlbum: permissionCreateStoryAlbum,
		PermissionFindAllAlbumStoriesForLoggedUser: permissionFindAllAlbumStoriesForLoggedUser,
		PermissionFindSelectedStoryAlbumByIdForLoggedUser: permissionFindSelectedStoryAlbumByIdForLoggedUser,
		PermissionFindAllPublicAlbumStoriesRegisteredUser: permissionFindAllPublicAlbumStoriesRegisteredUser,
		PermissionFindAllFollowingStoryAlbums: permissionFindAllFollowingStoryAlbums,
		PermissionCreateStoryHighlight: permissionCreateStoryHighlight,
		PermissionFindAllStoryHighlightsForUser: permissionFindAllStoryHighlightsForUser,
		LogInfo:                                       LogInfo,
		LogError:                                      LogError,
	}
}

func handleFunc(settingsAuthorizationHandler *handler.SettingsAuthorizationHandler, storyAuthorizationHandler *handler.StoryAuthorizationHandler, postAuthorizationHandler *handler.PostAuthorizationHandler,locationAuthorizationHandler *handler.LocationAuthorizationHandler, requestAuthorizationHandler *handler.RequestsAuthorizationHandler, contentAuthorizationHandler *handler.ContentAuthorizationHandler, tagAuthorizationHandler *handler.TagAuthorizationHandler,userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler, adminHandler *handler.AdminHandler, classicUserHandler *handler.ClassicUserHandler, agentHandler *handler.AgentHandler, registeredUserHandler *handler.RegisteredUserHandler,classicUserCampaignsHandler *handler.ClassicUserCampaignsHandler,classicUserFollowingsHandler *handler.ClassicUserFollowingsHandler,classicUserFollowersHandler *handler.ClassicUserFollowersHandler, recoveryPasswordTokenHandler *handler.RecoveryPasswordTokenHandler, classicUserCloseFriendsHandler *handler.ClassicUserCloseFriendsHandler){



	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/login/", userHandler.LogIn).Methods("POST")
	router.HandleFunc("/update_user_profile_info/", userHandler.UpdateUserProfileInfo).Methods("POST")
	router.HandleFunc("/find_user_by_id", userHandler.FindByID).Methods("GET")
	router.HandleFunc("/find_classic_user_by_id/{userID}", classicUserHandler.FindClassicUserById).Methods("GET")
	router.HandleFunc("/registered_admin/", adminHandler.CreateAdmin).Methods("POST")
	router.HandleFunc("/agent/", agentHandler.CreateAgent).Methods("POST")
	router.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/recovery_password/", recoveryPasswordTokenHandler.GenerateRecoveryPasswordToken).Methods("POST")
	router.HandleFunc("/verify_recovery_password_token/", recoveryPasswordTokenHandler.VerifyRecoveryPasswordToken).Methods("POST")
	router.HandleFunc("/confirm_registration/", confirmationTokenHandler.VerifyConfirmationToken).Methods("POST")
	router.HandleFunc("/change_user_password/", userHandler.ChangeUserPassword).Methods("POST")
	router.HandleFunc("/users/all",userHandler.FindAllUsers).Methods("GET")
	router.HandleFunc("/create_following/",classicUserFollowingsHandler.CreateClassicUserFollowing).Methods("POST")
	router.HandleFunc("/find_user_by_username", userHandler.FindByUserName).Methods("GET")
	router.HandleFunc("/find_all_classic_users_but_logged_in", classicUserHandler.FindAllUsersButLoggedIn).Methods("GET")
	router.HandleFunc("/find_selected_user_by_id", classicUserHandler.FindSelectedUserById).Methods("GET")
	router.HandleFunc("/accept_follow_request/", classicUserFollowingsHandler.AcceptFollowerRequest).Methods("POST")
	router.HandleFunc("/find_all_public_users/", classicUserHandler.FindAllPublicUsers).Methods("GET")
	router.HandleFunc("/find_all_valid_users/", classicUserHandler.FindAllValidUsers).Methods("GET")
	router.HandleFunc("/create_close_friend/", classicUserCloseFriendsHandler.CreateClassicUserCloseFriend).Methods("POST")
	router.HandleFunc("/check_if_user_valid/{userID}", classicUserHandler.CheckIfUserValid).Methods("GET")
	router.HandleFunc("/find_all_mutual_followers_for_user", classicUserFollowersHandler.FindAllMutualFollowerForUser).Methods("GET")
	router.HandleFunc("/dto/find_all_classic_users_but_logged_in", classicUserHandler.FindAllUsersButLoggedInDTOs).Methods("GET")
	router.HandleFunc("/check_if_following_post_story/{id}/{logId}", classicUserFollowingsHandler.CheckIfFollowingPostStory).Methods("GET")
	router.HandleFunc("/check_if_close_friend/{id}/{logId}", classicUserCloseFriendsHandler.CheckIfCloseFriend).Methods("GET")
	router.HandleFunc("/find_all_valid_followings_for_user/{id}", classicUserFollowingsHandler.FindAllValidFollowingsForUser).Methods("POST")
	router.HandleFunc("/find_all_user_who_follow_user_id/{id}", classicUserFollowingsHandler.FindAllUserWhoFollowUserId).Methods("POST")
	router.HandleFunc("/find_all_users_by_following_ids/", classicUserHandler.FindAllUsersByFollowingIds).Methods("POST")
	router.HandleFunc("/check_if_authentificated/", userHandler.CheckIfAuthentificated).Methods("GET")
	router.HandleFunc("/get_user_id_by_jwt/", userHandler.GetUserIDFromJWTToken).Methods("GET")
	router.HandleFunc("/remove_followings_between_users/", classicUserFollowingsHandler.RemoveFollowingsBetweenUsers).Methods("POST")

	//LOCATION MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-location-permission/", locationAuthorizationHandler.CheckCreateLocationPermission).Methods("POST")

	//CONTENT MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-single-post-content-permission/", contentAuthorizationHandler.CheckCreateSinglePostContentPermission).Methods("POST")
	router.HandleFunc("/auth/check-create-post-album-content-permission/", contentAuthorizationHandler.CheckCreatePostAlbumContentPermission).Methods("POST")
	router.HandleFunc("/auth/check-create-single-story-content-permission/", contentAuthorizationHandler.CheckCreateSingleStoryContentPermission).Methods("POST")
	router.HandleFunc("/auth/check-create-story-album-content-permission/", contentAuthorizationHandler.CheckCreateStoryAlbumContentPermission).Methods("POST")

	//STORY MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-single-story-permission/", storyAuthorizationHandler.CheckCreateSingleStoryPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-public-stories-registered-user-permission/", storyAuthorizationHandler.CheckFindAllPublicStoriesRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-stories-for-user-registered-user-permission/", storyAuthorizationHandler.CheckFindAllStoriesForUserRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-following-stories-permission/", storyAuthorizationHandler.CheckFindAllFollowingStoriesPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-selected-story-by-id-for-registered-users-permission/", storyAuthorizationHandler.CheckFindSelectedStoryByIdForRegisteredUsersPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-stories-for-logged-user-permission/", storyAuthorizationHandler.CheckFindAllStoriesForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-single-story-story-highlights-permission/", storyAuthorizationHandler.CheckCreateSingleStoryStoryHighlightsPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-single-story-story-highlights-for-story-highlight-permission/", storyAuthorizationHandler.CheckFindAllSingleStoryStoryHighlightsForStoryHighlightPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-story-album-permission/", storyAuthorizationHandler.CheckCreateStoryAlbumPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-album-stories-for-logged-user-permission/", storyAuthorizationHandler.CheckFindAllAlbumStoriesForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-selected-story-album-by-id-for-logged-user-permission/", storyAuthorizationHandler.CheckFindSelectedStoryAlbumByIdForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-public-album-stories-registered-user-permission/", storyAuthorizationHandler.CheckFindAllPublicAlbumStoriesRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-following-story-albums-permission/", storyAuthorizationHandler.CheckFindAllFollowingStoryAlbumsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-story-highlight-permission/", storyAuthorizationHandler.CheckCreateStoryHighlightPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-story-highlights-for-user-permission/", storyAuthorizationHandler.CheckFindAllStoryHighlightsForUserPermission).Methods("GET")

	//TAG MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-comment-tag-comments-permission/", tagAuthorizationHandler.CheckCreateCommentTagCommentsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-post-tag-posts-permission/", tagAuthorizationHandler.CheckCreatePostTagPostsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-post-album-tag-post-albums-permission/", tagAuthorizationHandler.CheckCreatePostAlbumTagPostAlbumsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-story-album-tag-story-albums-permission/", tagAuthorizationHandler.CheckCreateStoryAlbumTagStoryAlbumsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-story-tag-stories-permission/", tagAuthorizationHandler.CheckCreateStoryTagStoriesPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-taggable-users-story-permission/", tagAuthorizationHandler.CheckFindAllTaggableUsersStoryPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-taggable-users-post-permission/", tagAuthorizationHandler.CheckFindAllTaggableUsersPostPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-taggable-users-comment-permission/", tagAuthorizationHandler.CheckFindAllTaggableUsersCommentPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-comment-tag-comments-for-comment-permission/", tagAuthorizationHandler.CheckFindAllCommentTagCommentsForCommentPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-tag-permission/", tagAuthorizationHandler.CheckCreateTagPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-hashtags-permission/", tagAuthorizationHandler.CheckFindAllHashTagsPermission).Methods("GET")

	//POST MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-activity-permission/", postAuthorizationHandler.CheckCreateActivityPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-comment-permission/", postAuthorizationHandler.CheckCreateCommentPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-post-album-permission/", postAuthorizationHandler.CheckCreatePostAlbumPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-post-collection-permission/", postAuthorizationHandler.CheckCreatePostCollectionPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-post-collection-posts-permission/", postAuthorizationHandler.CheckCreatePostCollectionPostsPermission).Methods("GET")
	router.HandleFunc("/auth/check-create-single-post-permission/", postAuthorizationHandler.CheckCreateSinglePostPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-activities-for-post-permission/", postAuthorizationHandler.CheckFindAllActivitiesForPostPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-comments-for-post-permission/", postAuthorizationHandler.CheckFindAllCommentsForPostPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-following-post-albums-permission/", postAuthorizationHandler.CheckFindAllFollowingPostAlbumsPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-following-posts-permission/", postAuthorizationHandler.CheckFindAllFollowingPostsPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-album-posts-for-logged-user-permission/", postAuthorizationHandler.CheckFindAllAlbumPostsForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-locations-for-public-and-following-posts-permission/", postAuthorizationHandler.CheckFindAllLocationsForPublicAndFollowingPostsPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-post-collection-posts-for-post-permission/", postAuthorizationHandler.CheckFindAllPostCollectionPostsForPostPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-post-collections-for-user-registered-user-permission/", postAuthorizationHandler.CheckFindAllPostCollectionsForUserRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-tags-for-public-and-following-posts-permission/", postAuthorizationHandler.CheckFindAllTagsForPublicAndFollowingPostsPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-public-posts-registered-user-permission/", postAuthorizationHandler.CheckFindAllPublicPostsRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-update-activity-permission/", postAuthorizationHandler.CheckUpdateActivityPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-selected-post-by-id-for-logged-user-permission/", postAuthorizationHandler.CheckFindSelectedPostByIdForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-selected-post-album-by-id-for-logged-user-permission/", postAuthorizationHandler.CheckFindSelectedPostAlbumByIdForLoggedUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-posts-for-user-registered-user-permission/", postAuthorizationHandler.CheckFindAllPostsForUserRegisteredUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-posts-for-tag-reg-user-permission/", postAuthorizationHandler.CheckFindAllPostsForTagRegUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-posts-for-location-reg-user-permission/", postAuthorizationHandler.CheckFindAllPostsForLocationRegUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-posts-for-logged-user-permission/", postAuthorizationHandler.CheckFindAllPostsForLoggedUserPermission).Methods("GET")
	
	//REQUESTS MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-follow-request-permission/", requestAuthorizationHandler.CheckCreateFollowRequestPermission).Methods("GET")
	router.HandleFunc("/auth/check-reject-follow-request-permission/", requestAuthorizationHandler.CheckRejectFollowRequestPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-request-by-id-permission/", requestAuthorizationHandler.CheckFindRequestByIdPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-pending-follower-requests-for-user-permission/", requestAuthorizationHandler.CheckFindAllPendingFollowerRequestsForUserPermission).Methods("GET")

	router.HandleFunc("/auth/check-mute-user-permission/", settingsAuthorizationHandler.CheckMuteUserPermission).Methods("GET")
	router.HandleFunc("/auth/check-block-user-permission/", settingsAuthorizationHandler.CheckBlockUserPermission).Methods("GET")

	//FindAllValidFollowingsForUser
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()
	rbac := gorbac.New()
	validator := validator.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoUSER.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorUSER.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	roleRegisteredUser := gorbac.NewStdRole("role-registered-user")
	roleAgent := gorbac.NewStdRole("role-agent")
	roleAdmin := gorbac.NewStdRole("role-admin")

	permissionFindAllUsers := gorbac.NewStdPermission("permission-find-all-users")
	permissionUpdateUserInfo := gorbac.NewStdPermission("permission-update-user-info")
	permissionCreateClassicUserCloseFriend := gorbac.NewStdPermission("permission-create-classic-user-close-friend")
	permissionFindAllMutualFollowerForUser := gorbac.NewStdPermission("permission-find-all-mutual-follower-for-user")
	permissionCreateClassicUserFollowing := gorbac.NewStdPermission("permission-create-classic-user-following")
	permissionAcceptFollowerRequest := gorbac.NewStdPermission("permission-accept-follower-request")
	permissionFindAllUsersButLoggedIn := gorbac.NewStdPermission("permission-find-all-users-but-logged-in")
	permissionFindUserByID := gorbac.NewStdPermission("permission-find-user-by-id")
	
	permissionCreateCommentTagComments := gorbac.NewStdPermission("permission-create-comment-tag-comments")
	permissionFindAllTaggableUsersPost := gorbac.NewStdPermission("permission-find-all-taggable-users-post")
	permissionCreateTag := gorbac.NewStdPermission("permission-create-tag")
	permissionCreatePostTagPosts := gorbac.NewStdPermission("permission-create-post-tag-posts")
	permissionCreatePostAlbumTagPostAlbums := gorbac.NewStdPermission("permission-create-post-album-tag-post-albums")
	permissionCreateStoryTagStories := gorbac.NewStdPermission("permission-create-story-tag-stories")
	permissionCreateStoryAlbumTagStoryAlbums := gorbac.NewStdPermission("permission-create-story-album-tag-story-albums")
	permissionFindAllTaggableUsersStory := gorbac.NewStdPermission("permission-find-all-taggable-users-story")
	permissionFindAllCommentTagCommentsForComment := gorbac.NewStdPermission("permission-find-all-comment-tag-comments-for-comment")
	permissionFindAllTaggableUsersComment := gorbac.NewStdPermission("permission-find-all-taggable-users-comment")
	permissionFindAllHashTags := gorbac.NewStdPermission("permission-find-all-hash-tags")

	permissionCreateSinglePost := gorbac.NewStdPermission("permission-create-single-post")
	permissionCreatePostAlbum := gorbac.NewStdPermission("permission-create-post-album")
	permissionFindAllFollowingPostAlbums := gorbac.NewStdPermission("permission-find-all-following-post-albums")
	permissionFindAllFollowingPosts := gorbac.NewStdPermission("permission-find-all-following-posts")
	permissionCreatePostCollection := gorbac.NewStdPermission("permission-create-post-collection")
	permissionFindAllPostCollectionsForUserRegisteredUser := gorbac.NewStdPermission("permission-find-all-post-collections-for-user-registered-user")
	permissionFindAllPostsForLoggedUser := gorbac.NewStdPermission("permission-find-all-posts-for-logged-user")
	permissionFindAllAlbumPostsForLoggedUser := gorbac.NewStdPermission("permission-find-all-album-posts-for-logged-user")
	permissionCreateComment := gorbac.NewStdPermission("permission-create-comment")
	permissionFindSelectedPostByIdForLoggedUser := gorbac.NewStdPermission("permission-find-selected-post-by-id-for-logged-user")
	permissionFindAllCommentsForPost := gorbac.NewStdPermission("permission-find-all-comments-for-post")
	permissionFindAllActivitiesForPost := gorbac.NewStdPermission("permission-find-all-activities-for-post")
	permissionUpdateActivity := gorbac.NewStdPermission("permission-update-activity")
	permissionCreateActivity := gorbac.NewStdPermission("permission-create-activity")
	permissionFindAllPostCollectionPostsForPost := gorbac.NewStdPermission("permission-find-all-post-collection-posts-for-post")
	permissionCreatePostCollectionPosts := gorbac.NewStdPermission("permission-create-post-collection-posts")
	permissionFindAllPostsForLocationRegUser := gorbac.NewStdPermission("permission-find-all-posts-for-location-reg-user")
	permissionFindSelectedPostAlbumByIdForLoggedUser := gorbac.NewStdPermission("permission-find-selected-post-album-by-id-for-logged-user")
	permissionFindAllPostsForTagRegUser := gorbac.NewStdPermission("permission-find-all-posts-for-tag-reg-user")
	permissionFindAllPublicPostsRegisteredUser := gorbac.NewStdPermission("permission-find-all-public-posts-registered-user")
	permissionFindAllPostsForUserRegisteredUser := gorbac.NewStdPermission("permission-find-all-posts-for-user-registered-user")
	permissionFindAllTagsForPublicAndFollowingPosts := gorbac.NewStdPermission("permission-find-all-tags-for-public-and-following-posts")
	permissionFindAllLocationsForPublicAndFollowingPosts := gorbac.NewStdPermission("permission-find-all-locations-for-public-and-following-posts")

	permissionCreateSinglePostContent := gorbac.NewStdPermission("permission-create-single-post-content")
	permissionCreatePostAlbumContent := gorbac.NewStdPermission("permission-create-post-album-content")
	permissionCreateSingleStoryContent := gorbac.NewStdPermission("permission-create-single-story-content")
	permissionCreateStoryAlbumContent := gorbac.NewStdPermission("permission-create-story-album-content")
	
	permissionCreateFollowRequest := gorbac.NewStdPermission("permission-create-follow-request")
	permissionRejectFollowRequest := gorbac.NewStdPermission("permission-reject-follow-request")
	permissionFindRequestById := gorbac.NewStdPermission("permission-find-request-by-id")
	permissionFindAllPendingFollowerRequestsForUser := gorbac.NewStdPermission("permission-find-all-pending-follower-requests-for-user")
	
	permissionCreateLocation := gorbac.NewStdPermission("permission-create-location")

	permissionCreateSingleStory := gorbac.NewStdPermission("permission-create-single-story")
	permissionFindAllPublicStoriesRegisteredUser := gorbac.NewStdPermission("permission-find-all-public-stories-registered-user")
	permissionFindAllStoriesForUserRegisteredUser := gorbac.NewStdPermission("permission-finc-all-stories-for-user-registered-user")
	permissionFindAllFollowingStories := gorbac.NewStdPermission("permission-find-all-following-stories")
	permissionFindSelectedStoryByIdForRegisteredUsers := gorbac.NewStdPermission("permission-find-selected-story-by-id-for-registered-users")
	permissionFindAllStoriesForLoggedUser := gorbac.NewStdPermission("permission-find-all-stories-for-logged-user")
	permissionCreateSingleStoryStoryHighlights  := gorbac.NewStdPermission("permission-create-single-story-story-highlights")
	permissionFindAllSingleStoryStoryHighlightsForStory := gorbac.NewStdPermission("permission-create-single-story-story-highlights")
	permissionFindAllSingleStoryStoryHighlightsForStoryHighlight := gorbac.NewStdPermission("permission-find-all-single-story-story-highlights-for-story-highlight")
	permissionCreateStoryAlbum := gorbac.NewStdPermission("permission-create-story-album")
	permissionFindAllAlbumStoriesForLoggedUser := gorbac.NewStdPermission("permission-find-all-album-stories-for-logged-user")
	permissionFindSelectedStoryAlbumByIdForLoggedUser := gorbac.NewStdPermission("permission-create-single-story-story-highlights")
	permissionFindAllPublicAlbumStoriesRegisteredUser := gorbac.NewStdPermission("permission-find-selected-story-album-by-id-for-logged-user")
	permissionFindAllFollowingStoryAlbums := gorbac.NewStdPermission("permission-find-all-following-story-albums")
	permissionCreateStoryHighlight := gorbac.NewStdPermission("permission-create-story-highlight")
	permissionFindAllStoryHighlightsForUser := gorbac.NewStdPermission("permission-find-all-story-highlights-for-user")

	permissionMuteUser := gorbac.NewStdPermission("permission-mute-user")
	permissionBlockUser := gorbac.NewStdPermission("permission-block-user")

	roleAdmin.Assign(permissionFindAllUsers)
	roleAdmin.Assign(permissionUpdateUserInfo)
	roleAdmin.Assign(permissionFindUserByID)

	roleAgent.Assign(permissionUpdateUserInfo)
	roleAgent.Assign(permissionCreateClassicUserCloseFriend)
	roleAgent.Assign(permissionFindAllMutualFollowerForUser)
	roleAgent.Assign(permissionCreateClassicUserFollowing)
	roleAgent.Assign(permissionAcceptFollowerRequest)
	roleAgent.Assign(permissionFindAllUsersButLoggedIn)
	roleAgent.Assign(permissionFindUserByID)
	roleAgent.Assign(permissionCreateCommentTagComments)
	roleAgent.Assign(permissionFindAllTaggableUsersPost)
	roleAgent.Assign(permissionCreateTag)
	roleAgent.Assign(permissionCreatePostTagPosts)
	roleAgent.Assign(permissionCreatePostAlbumTagPostAlbums)
	roleAgent.Assign(permissionCreateStoryTagStories)
	roleAgent.Assign(permissionCreateStoryAlbumTagStoryAlbums)
	roleAgent.Assign(permissionFindAllTaggableUsersStory)
	roleAgent.Assign(permissionFindAllCommentTagCommentsForComment)
	roleAgent.Assign(permissionFindAllTaggableUsersComment)
	roleAgent.Assign(permissionFindAllHashTags)
	roleAgent.Assign(permissionCreateSinglePost)
	roleAgent.Assign(permissionCreatePostAlbum)
	roleAgent.Assign(permissionFindAllFollowingPostAlbums)
	roleAgent.Assign(permissionFindAllFollowingPosts)
	roleAgent.Assign(permissionCreatePostCollection)
	roleAgent.Assign(permissionFindAllPostCollectionsForUserRegisteredUser)
	roleAgent.Assign(permissionFindAllPostsForLoggedUser)
	roleAgent.Assign(permissionFindAllAlbumPostsForLoggedUser)
	roleAgent.Assign(permissionCreateComment)
	roleAgent.Assign(permissionFindSelectedPostByIdForLoggedUser)
	roleAgent.Assign(permissionFindAllCommentsForPost)
	roleAgent.Assign(permissionFindAllActivitiesForPost)
	roleAgent.Assign(permissionUpdateActivity)
	roleAgent.Assign(permissionCreateActivity)
	roleAgent.Assign(permissionFindAllPostCollectionPostsForPost)
	roleAgent.Assign(permissionCreatePostCollectionPosts)
	roleAgent.Assign(permissionFindAllPostsForLocationRegUser)
	roleAgent.Assign(permissionFindSelectedPostAlbumByIdForLoggedUser)
	roleAgent.Assign(permissionFindAllPostsForTagRegUser)
	roleAgent.Assign(permissionFindAllPublicPostsRegisteredUser)
	roleAgent.Assign(permissionFindAllPostsForUserRegisteredUser)
	roleAgent.Assign(permissionFindAllTagsForPublicAndFollowingPosts)
	roleAgent.Assign(permissionFindAllLocationsForPublicAndFollowingPosts)
	roleAgent.Assign(permissionCreateSinglePostContent)
	roleAgent.Assign(permissionCreatePostAlbumContent)
	roleAgent.Assign(permissionCreateSingleStoryContent)
	roleAgent.Assign(permissionCreateStoryAlbumContent)
	roleAgent.Assign(permissionCreateFollowRequest)
	roleAgent.Assign(permissionRejectFollowRequest)
	roleAgent.Assign(permissionFindRequestById)
	roleAgent.Assign(permissionFindAllPendingFollowerRequestsForUser)
	roleAgent.Assign(permissionCreateLocation)
	roleAgent.Assign(permissionBlockUser)
	roleAgent.Assign(permissionMuteUser)

	roleAgent.Assign(permissionCreateSingleStory)
	roleAgent.Assign(permissionFindAllPublicStoriesRegisteredUser)
	roleAgent.Assign(permissionFindAllStoriesForUserRegisteredUser)
	roleAgent.Assign(permissionFindAllFollowingStories)
	roleAgent.Assign(permissionFindSelectedStoryByIdForRegisteredUsers)
	roleAgent.Assign(permissionFindAllStoriesForLoggedUser)
	roleAgent.Assign(permissionCreateSingleStoryStoryHighlights)
	roleAgent.Assign(permissionFindAllSingleStoryStoryHighlightsForStory)
	roleAgent.Assign(permissionCreateStoryAlbum)
	roleAgent.Assign(permissionFindAllAlbumStoriesForLoggedUser)
	roleAgent.Assign(permissionFindSelectedStoryAlbumByIdForLoggedUser)
	roleAgent.Assign(permissionFindAllPublicAlbumStoriesRegisteredUser)
	roleAgent.Assign(permissionFindAllFollowingStoryAlbums)
	roleAgent.Assign(permissionCreateStoryHighlight)
	roleAgent.Assign(permissionFindAllStoryHighlightsForUser)

	roleRegisteredUser.Assign(permissionUpdateUserInfo)
	roleRegisteredUser.Assign(permissionCreateClassicUserCloseFriend)
	roleRegisteredUser.Assign(permissionFindAllMutualFollowerForUser)
	roleRegisteredUser.Assign(permissionCreateClassicUserFollowing)
	roleRegisteredUser.Assign(permissionAcceptFollowerRequest)
	roleRegisteredUser.Assign(permissionFindAllUsersButLoggedIn)
	roleRegisteredUser.Assign(permissionFindUserByID)
	roleRegisteredUser.Assign(permissionCreateCommentTagComments)
	roleRegisteredUser.Assign(permissionFindAllTaggableUsersPost)
	roleRegisteredUser.Assign(permissionCreateTag)
	roleRegisteredUser.Assign(permissionCreatePostTagPosts)
	roleRegisteredUser.Assign(permissionCreatePostAlbumTagPostAlbums)
	roleRegisteredUser.Assign(permissionCreateStoryTagStories)
	roleRegisteredUser.Assign(permissionCreateStoryAlbumTagStoryAlbums)
	roleRegisteredUser.Assign(permissionFindAllTaggableUsersStory)
	roleRegisteredUser.Assign(permissionFindAllCommentTagCommentsForComment)
	roleRegisteredUser.Assign(permissionFindAllTaggableUsersComment)
	roleRegisteredUser.Assign(permissionFindAllHashTags)
	roleRegisteredUser.Assign(permissionCreateSinglePost)
	roleRegisteredUser.Assign(permissionCreatePostAlbum)
	roleRegisteredUser.Assign(permissionFindAllFollowingPostAlbums)
	roleRegisteredUser.Assign(permissionFindAllFollowingPosts)
	roleRegisteredUser.Assign(permissionCreatePostCollection)
	roleRegisteredUser.Assign(permissionFindAllPostCollectionsForUserRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllPostsForLoggedUser)
	roleRegisteredUser.Assign(permissionFindAllAlbumPostsForLoggedUser)
	roleRegisteredUser.Assign(permissionCreateComment)
	roleRegisteredUser.Assign(permissionFindSelectedPostByIdForLoggedUser)
	roleRegisteredUser.Assign(permissionFindAllCommentsForPost)
	roleRegisteredUser.Assign(permissionFindAllActivitiesForPost)
	roleRegisteredUser.Assign(permissionUpdateActivity)
	roleRegisteredUser.Assign(permissionCreateActivity)
	roleRegisteredUser.Assign(permissionFindAllPostCollectionPostsForPost)
	roleRegisteredUser.Assign(permissionCreatePostCollectionPosts)
	roleRegisteredUser.Assign(permissionFindAllPostsForLocationRegUser)
	roleRegisteredUser.Assign(permissionFindSelectedPostAlbumByIdForLoggedUser)
	roleRegisteredUser.Assign(permissionFindAllPostsForTagRegUser)
	roleRegisteredUser.Assign(permissionFindAllPublicPostsRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllPostsForUserRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllTagsForPublicAndFollowingPosts)
	roleRegisteredUser.Assign(permissionFindAllLocationsForPublicAndFollowingPosts)
	roleRegisteredUser.Assign(permissionCreateSinglePostContent)
	roleRegisteredUser.Assign(permissionCreatePostAlbumContent)
	roleRegisteredUser.Assign(permissionCreateSingleStoryContent)
	roleRegisteredUser.Assign(permissionCreateStoryAlbumContent)
	roleRegisteredUser.Assign(permissionCreateFollowRequest)
	roleRegisteredUser.Assign(permissionRejectFollowRequest)
	roleRegisteredUser.Assign(permissionFindRequestById)
	roleRegisteredUser.Assign(permissionFindAllPendingFollowerRequestsForUser)
	roleRegisteredUser.Assign(permissionCreateLocation)

	roleRegisteredUser.Assign(permissionCreateSingleStory)
	roleRegisteredUser.Assign(permissionFindAllPublicStoriesRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllStoriesForUserRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllFollowingStories)
	roleRegisteredUser.Assign(permissionFindSelectedStoryByIdForRegisteredUsers)
	roleRegisteredUser.Assign(permissionFindAllStoriesForLoggedUser)
	roleRegisteredUser.Assign(permissionCreateSingleStoryStoryHighlights)
	roleRegisteredUser.Assign(permissionFindAllSingleStoryStoryHighlightsForStory)
	roleRegisteredUser.Assign(permissionCreateStoryAlbum)
	roleRegisteredUser.Assign(permissionFindAllAlbumStoriesForLoggedUser)
	roleRegisteredUser.Assign(permissionFindSelectedStoryAlbumByIdForLoggedUser)
	roleRegisteredUser.Assign(permissionFindAllPublicAlbumStoriesRegisteredUser)
	roleRegisteredUser.Assign(permissionFindAllFollowingStoryAlbums)
	roleRegisteredUser.Assign(permissionCreateStoryHighlight)
	roleRegisteredUser.Assign(permissionFindAllStoryHighlightsForUser)
	roleRegisteredUser.Assign(permissionBlockUser)
	roleRegisteredUser.Assign(permissionMuteUser)

	rbac.Add(roleAdmin)
	rbac.Add(roleAgent)
	rbac.Add(roleRegisteredUser)


	database := initDB()
	userRepo := initUserRepo(database)
	registeredUserRepo := initRegisteredUserRepo(database)
	adminRepo := initAdminRepo(database)
	classicUserRepo := initClassicUserRepo(database)
	agentRepo := initAgentRepo(database)
	confirmationTokenRepo := initConfirmationTokenRepo(database)
	classicUserCampaignsRepo := initClassicUserCampaignsRepo(database)
	classicUserFollowersRepo := initClassicUserFollowersRepo(database)
	classicUserFollowingsRepo := initClassicUserFollowingsRepo(database)
	recoveryPasswordTokenRepo := initRecoveryPasswordTokenRepo(database)
	classicUserCloseFriendsRepo := initClassicUserCloseFriendsRepo(database)

	userService := initUserService(userRepo)
	registeredUserService := initRegisteredUserService(registeredUserRepo)
	confirmationTokenService := initConfirmationTokenService(confirmationTokenRepo)
	adminService := initAdminService(adminRepo)
	classicUserService := initClassicUserService(classicUserRepo)
	agentService := initAgentService(agentRepo)
	classicUserCampaignsService := initClassicUserCampaignsService(classicUserCampaignsRepo)
	classicUserFollowersService := initClassicUserFollowersService(classicUserFollowersRepo)
	classicUserFollowingsService := initClassicUserFollowingsService(classicUserFollowingsRepo)
	recoveryPasswordTokenService := initRecoveryPasswordTokenService(recoveryPasswordTokenRepo)
	classicUserCloseFriendsService := initClassicUserCloseFriendsService(classicUserCloseFriendsRepo)

	tagAuthorizationHandler := initTagAuthorizationHandler(rbac, &permissionCreateCommentTagComments,&permissionFindAllHashTags,
		&permissionCreateStoryAlbumTagStoryAlbums,&permissionFindAllTaggableUsersStory,&permissionFindAllTaggableUsersComment,
		&permissionCreatePostTagPosts,&permissionCreatePostAlbumTagPostAlbums,&permissionFindAllCommentTagCommentsForComment,
		&permissionCreateTag,&permissionFindAllTaggableUsersPost, logInfo,logError,userService)

	storyAuthorizationHandler := initStoryAuthorizationHandler(userService, rbac, &permissionCreateSingleStory, &permissionFindAllPublicStoriesRegisteredUser, 
		&permissionFindAllStoriesForUserRegisteredUser, &permissionFindAllFollowingStories, &permissionFindSelectedStoryByIdForRegisteredUsers,
		&permissionFindAllStoriesForLoggedUser, &permissionCreateSingleStoryStoryHighlights, &permissionFindAllSingleStoryStoryHighlightsForStory,
		&permissionFindAllSingleStoryStoryHighlightsForStoryHighlight, &permissionCreateStoryAlbum, &permissionFindAllAlbumStoriesForLoggedUser,
		&permissionFindSelectedStoryAlbumByIdForLoggedUser, &permissionFindAllPublicAlbumStoriesRegisteredUser, &permissionFindAllFollowingStoryAlbums,
		&permissionCreateStoryHighlight, &permissionFindAllStoryHighlightsForUser, logInfo, logError)

	settingsAuthorizationHandler := initSettingsAuthorizationHandler(rbac, logInfo,logError,userService, &permissionBlockUser, &permissionMuteUser)

	postAuthorizationHandler := initPostAuthorizationHandler(rbac, logInfo,logError,userService,&permissionCreateSinglePost,
		&permissionCreatePostAlbum,&permissionFindAllFollowingPostAlbums,&permissionFindAllFollowingPosts,&permissionCreatePostCollection,
		&permissionFindAllPostCollectionsForUserRegisteredUser,&permissionFindAllPostsForLoggedUser,
		&permissionFindAllAlbumPostsForLoggedUser,&permissionCreateComment,&permissionFindSelectedPostByIdForLoggedUser,
		&permissionFindAllCommentsForPost,&permissionFindAllActivitiesForPost,&permissionUpdateActivity,&permissionCreateActivity,
		&permissionFindAllPostCollectionPostsForPost,&permissionCreatePostCollectionPosts,&permissionFindAllPostsForLocationRegUser,
		&permissionFindSelectedPostAlbumByIdForLoggedUser,&permissionFindAllPostsForTagRegUser,&permissionFindAllPublicPostsRegisteredUser,
		&permissionFindAllPostsForUserRegisteredUser,&permissionFindAllTagsForPublicAndFollowingPosts,&permissionFindAllLocationsForPublicAndFollowingPosts)
	
	requestsAuthorizationHandler := initRequestsAuthorizationHandler(rbac, &permissionCreateFollowRequest, &permissionRejectFollowRequest, 
		&permissionFindRequestById, &permissionFindAllPendingFollowerRequestsForUser, logInfo,logError,userService)
	
	contentAuthorizationHandler := initContentAuthorizationHandler(rbac, &permissionCreateSinglePostContent,&permissionCreatePostAlbumContent,
		&permissionCreateSingleStoryContent,&permissionCreateStoryAlbumContent,logInfo,logError,userService)

	locationAuthorizationHandler := initLocationAuthorizationHandler(rbac, &permissionCreateLocation,logInfo,logError,userService)


	passwordUtil := initPasswordUtil()
	userHandler := initUserHandler(&permissionFindUserByID,logInfo,logError,recoveryPasswordTokenService,userService,adminService,classicUserService,registeredUserService,agentService, rbac, &permissionFindAllUsers, &permissionUpdateUserInfo, validator, passwordUtil)
	adminHandler := initAdminHandler(logInfo,logError,adminService, userService, validator, passwordUtil)
	registeredUserHandler := initRegisteredUserHandler(logInfo,logError,registeredUserService, userService, classicUserService,confirmationTokenService,validator, passwordUtil)
	agentHandler := initAgentHandler(logInfo,logError,agentService, userService, classicUserService, validator, passwordUtil)
	confirmationTokenHandler := initConfirmationTokenHandler(logInfo,logError,confirmationTokenService,userService,registeredUserService,classicUserService)
	classicUserCampaignsHandler := initClassicUserCampaignsHandler(logInfo,logError,classicUserCampaignsService)
	classicUserFollowersHandler := initClassicUserFollowersHandler(logInfo,logError,rbac,&permissionFindAllMutualFollowerForUser,userService,classicUserFollowersService)
	classicUserFollowingsHandler := initClassicUserFollowingsHandler(logInfo,logError,rbac,&permissionAcceptFollowerRequest,&permissionCreateClassicUserFollowing,userService,classicUserFollowingsService,classicUserFollowersService)
	recoveryPasswordTokenHandler := initRecoveryPasswordTokenHandler(logInfo,logError,recoveryPasswordTokenService,userService, validator)
	classicUserHandler := initClassicUserHandler(userService,&permissionFindAllUsersButLoggedIn,rbac,logInfo,logError,classicUserService, classicUserFollowingsService)
	classicUserCloseFriendsHandler := initClassicUserCloseFriendsHandler( userService, rbac, &permissionCreateClassicUserCloseFriend, logInfo,logError,classicUserCloseFriendsService, classicUserFollowersService)

	handleFunc(settingsAuthorizationHandler,storyAuthorizationHandler, postAuthorizationHandler,locationAuthorizationHandler, requestsAuthorizationHandler, contentAuthorizationHandler,tagAuthorizationHandler,userHandler, confirmationTokenHandler, adminHandler,classicUserHandler, agentHandler,registeredUserHandler,classicUserCampaignsHandler,classicUserFollowingsHandler,classicUserFollowersHandler,recoveryPasswordTokenHandler, classicUserCloseFriendsHandler)

}