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

func handleFunc(requestAuthorizationHandler *handler.RequestsAuthorizationHandler, tagAuthorizationHandler *handler.TagAuthorizationHandler,userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler, adminHandler *handler.AdminHandler, classicUserHandler *handler.ClassicUserHandler, agentHandler *handler.AgentHandler, registeredUserHandler *handler.RegisteredUserHandler,classicUserCampaignsHandler *handler.ClassicUserCampaignsHandler,classicUserFollowingsHandler *handler.ClassicUserFollowingsHandler,classicUserFollowersHandler *handler.ClassicUserFollowersHandler, recoveryPasswordTokenHandler *handler.RecoveryPasswordTokenHandler, classicUserCloseFriendsHandler *handler.ClassicUserCloseFriendsHandler){

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


	//autorizacija tag
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

	//REQUEST MICROSERVICE AUTHORIZATION
	router.HandleFunc("/auth/check-create-follow-request-permission/", requestAuthorizationHandler.CheckCreateFollowRequestPermission).Methods("GET")
	router.HandleFunc("/auth/check-reject-follow-request-permission/", requestAuthorizationHandler.CheckRejectFollowRequestPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-request-by-id-permission/", requestAuthorizationHandler.CheckFindRequestByIdPermission).Methods("GET")
	router.HandleFunc("/auth/check-find-all-pending-follower-requests-for-user-permission/", requestAuthorizationHandler.CheckFindAllPendingFollowerRequestsForUserPermission).Methods("GET")

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
	permissionFindAllHashTags := gorbac.NewStdPermission("permission-FindAllHashTags")

	//REQUESTS MICROSERVICE
	permissionCreateFollowRequest := gorbac.NewStdPermission("permission-create-follow-request")
	permissionRejectFollowRequest := gorbac.NewStdPermission("permission-reject-follow-request")
	permissionFindRequestById := gorbac.NewStdPermission("permission-find-request-by-id")
	permissionFindAllPendingFollowerRequestsForUser := gorbac.NewStdPermission("permission-find-all-pending-follower-requests-for-user")



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

	//REQUESTS MICROSERVICE
	roleAgent.Assign(permissionCreateFollowRequest)
	roleAgent.Assign(permissionRejectFollowRequest)
	roleAgent.Assign(permissionFindRequestById)
	roleAgent.Assign(permissionFindAllPendingFollowerRequestsForUser)

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

	//REQUESTS MICROSERVICE
	roleRegisteredUser.Assign(permissionCreateFollowRequest)
	roleRegisteredUser.Assign(permissionRejectFollowRequest)
	roleRegisteredUser.Assign(permissionFindRequestById)
	roleRegisteredUser.Assign(permissionFindAllPendingFollowerRequestsForUser)

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

	requestsAuthorizationHandler := initRequestsAuthorizationHandler(rbac, &permissionCreateFollowRequest, &permissionRejectFollowRequest, &permissionFindRequestById, &permissionFindAllPendingFollowerRequestsForUser, logInfo,logError,userService)

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
	handleFunc(requestsAuthorizationHandler,tagAuthorizationHandler,userHandler, confirmationTokenHandler, adminHandler,classicUserHandler, agentHandler,registeredUserHandler,classicUserCampaignsHandler,classicUserFollowingsHandler,classicUserFollowersHandler,recoveryPasswordTokenHandler, classicUserCloseFriendsHandler)

}