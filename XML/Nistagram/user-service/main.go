package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mikespook/gorbac"
	requestsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
	requestsService "github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	settingsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
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
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=root dbname=nistagram-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.ClassicUser{}, &model.RegisteredUser{}, &model.Admin{}, &model.Agent{},&model.ClassicUserFollowers{},&model.ClassicUserFollowings{}, &model.ClassicUserCampaigns{}, &model.ConfirmationToken{}, &model.RecoveryPasswordToken{})
	return db
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

func initUserHandler(UserService *service.UserService,AdminService *service.AdminService, ClassicUserService *service.ClassicUserService, RegisteredUserService *service.RegisteredUserService, AgentService *service.AgentService, rbac *gorbac.RBAC, permissionFindAllUsers *gorbac.Permission, permissionUpdateUserInfo *gorbac.Permission, validator *validator.Validate, passwordUtil *util.PasswordUtil ) *handler.UserHandler{
	return &handler.UserHandler{
		UserService:            UserService,
		AdminService:           AdminService,
		ClassicUserService:     ClassicUserService,
		RegisteredUserService:  RegisteredUserService,
		AgentService:           AgentService,
		Rbac:                   rbac,
		PermissionFindAllUsers: permissionFindAllUsers,
		PermissionUpdateUserInfo: permissionUpdateUserInfo,
		Validator: validator,
		PasswordUtil: passwordUtil,
	}
}
//SETTINGS
func initSettingsRepo(database *gorm.DB) *settingsRepository.ProfileSettingsRepository{
	return &settingsRepository.ProfileSettingsRepository { Database: database }
}

func initSettingsService(repo *settingsRepository.ProfileSettingsRepository) *settingsService.ProfileSettingsService{
	return &settingsService.ProfileSettingsService { Repo: repo }
}

//REQUESTS
func initRequestsRepo(database *gorm.DB) *requestsRepository.FollowRequestRepository{
	return &requestsRepository.FollowRequestRepository { Database: database }
}

func initRequestsService(repo *requestsRepository.FollowRequestRepository) *requestsService.FollowRequestService{
	return &requestsService.FollowRequestService { Repo: repo }
}



//ADMIN
func initAdminRepo(database *gorm.DB) *repository.AdminRepository{
	return &repository.AdminRepository { Database: database }
}

func initAdminService(repo *repository.AdminRepository) *service.AdminService{
	return &service.AdminService { Repo: repo }
}

func initAdminHandler(adminService *service.AdminService, userService *service.UserService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.AdminHandler{
	return &handler.AdminHandler {
		AdminService: adminService,
		UserService: userService,
		Validator: validator,
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

func initClassicUserHandler(classicUserService *service.ClassicUserService, profileSettingsService *settingsService.ProfileSettingsService, classicUserFollowingsService *service.ClassicUserFollowingsService, requestsService *requestsService.FollowRequestService ) *handler.ClassicUserHandler{
	return &handler.ClassicUserHandler { ClassicUserService: classicUserService, ProfileSettingsService: profileSettingsService, ClassicUserFollowingsService: classicUserFollowingsService, FollowRequestService: requestsService}
}

//REGISTERED USER
func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}

func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}

func initRegisteredUserHandler(registeredUserService *service.RegisteredUserService, userService *service.UserService, classicUserService *service.ClassicUserService,  confirmationTokenService *service.ConfirmationTokenService,  settingsService *settingsService.ProfileSettingsService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler{
		registeredUserService,
		userService,
		classicUserService ,
		confirmationTokenService,
		settingsService,
		validator,
		passwordUtil,
	}

}

func initAgentRepo(database *gorm.DB) *repository.AgentRepository{
	return &repository.AgentRepository { Database: database }
}

func initAgentService(repo *repository.AgentRepository) *service.AgentService{
	return &service.AgentService { Repo: repo }
}

func initAgentHandler(agentService *service.AgentService, userService *service.UserService, classicUserService *service.ClassicUserService, validator *validator.Validate, passwordUtil *util.PasswordUtil) *handler.AgentHandler{
	return &handler.AgentHandler{
		AgentService: agentService,
		UserService: userService,
		ClassicUserService: classicUserService,
		Validator: validator,
		PasswordUtil: passwordUtil,
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

func initClassicUserCampaignsHandler(service *service.ClassicUserCampaignsService) *handler.ClassicUserCampaignsHandler{
	return &handler.ClassicUserCampaignsHandler { Service: service }
}

func initClassicUserFollowersHandler(service *service.ClassicUserFollowersService) *handler.ClassicUserFollowersHandler{
	return &handler.ClassicUserFollowersHandler { ClassicUserFollowersService: service}
}

func initClassicUserFollowingsHandler(classicUserFollowings *service.ClassicUserFollowingsService, classicUserFollowersService *service.ClassicUserFollowersService, followRequestService *requestsService.FollowRequestService) *handler.ClassicUserFollowingsHandler{
	return &handler.ClassicUserFollowingsHandler { ClassicUserFollowingsService: classicUserFollowings , ClassicUserFollowersService: classicUserFollowersService, FollowRequestService: followRequestService}
}

func initRecoveryPasswordTokenHandler(recoveryPasswordTokenService *service.RecoveryPasswordTokenService, classicUserService *service.ClassicUserService, registeredUserService *service.RegisteredUserService, userService *service.UserService, validator *validator.Validate) *handler.RecoveryPasswordTokenHandler{
	return &handler.RecoveryPasswordTokenHandler{
		RecoveryPasswordTokenService: recoveryPasswordTokenService,
		UserService:                  userService,
		Validator: validator,
	}
}

func initConfirmationTokenHandler(confirmationTokenService *service.ConfirmationTokenService, userService *service.UserService, registeredUserService *service.RegisteredUserService, classicUserService *service.ClassicUserService) *handler.ConfirmationTokenHandler{
	return &handler.ConfirmationTokenHandler{
		ConfirmationTokenService: confirmationTokenService,
		ClassicUserService:       classicUserService,
		RegisteredUserService:    registeredUserService,
		UserService:              userService,
	}
}

func handleFunc(userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler, adminHandler *handler.AdminHandler, classicUserHandler *handler.ClassicUserHandler, agentHandler *handler.AgentHandler, registeredUserHandler *handler.RegisteredUserHandler,classicUserCampaignsHandler *handler.ClassicUserCampaignsHandler,classicUserFollowingsHandler *handler.ClassicUserFollowingsHandler,classicUserFollowersHandler *handler.ClassicUserFollowersHandler, recoveryPasswordTokenHandler *handler.RecoveryPasswordTokenHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/login/", userHandler.LogIn).Methods("POST")
	router.HandleFunc("/update_user_profile_info/", userHandler.UpdateUserProfileInfo).Methods("POST")
	router.HandleFunc("/find_user_by_id", userHandler.FindByID).Methods("GET")
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


	log.Fatal(http.ListenAndServe(":8080", cors(router)))
}

func main() {
	rbac := gorbac.New()
	validator := validator.New()

	roleRegisteredUser := gorbac.NewStdRole("role-registered-user")
	roleAgent := gorbac.NewStdRole("role-agent")
	roleAdmin := gorbac.NewStdRole("role-admin")

	permissionFindAllUsers := gorbac.NewStdPermission("permission-find-all-users")
	permissionUpdateUserInfo := gorbac.NewStdPermission("permission-update-user-info")

	roleAdmin.Assign(permissionFindAllUsers)
	roleAdmin.Assign(permissionUpdateUserInfo)

	roleAgent.Assign(permissionUpdateUserInfo)

	roleRegisteredUser.Assign(permissionUpdateUserInfo)

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
	settingsRepo := initSettingsRepo(database)
	requestsRepo := initRequestsRepo(database)

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
	settingsService := initSettingsService(settingsRepo)
	requestsService := initRequestsService(requestsRepo)


	passwordUtil := initPasswordUtil()
	userHandler := initUserHandler(userService,adminService,classicUserService,registeredUserService,agentService, rbac, &permissionFindAllUsers, &permissionUpdateUserInfo, validator, passwordUtil)
	adminHandler := initAdminHandler(adminService, userService, validator, passwordUtil)
	registeredUserHandler := initRegisteredUserHandler(registeredUserService, userService, classicUserService,confirmationTokenService,settingsService,validator, passwordUtil)
	agentHandler := initAgentHandler(agentService, userService, classicUserService, validator, passwordUtil)
	confirmationTokenHandler := initConfirmationTokenHandler(confirmationTokenService,userService,registeredUserService,classicUserService)
	classicUserCampaignsHandler := initClassicUserCampaignsHandler(classicUserCampaignsService)
	classicUserFollowersHandler := initClassicUserFollowersHandler(classicUserFollowersService)
	classicUserFollowingsHandler := initClassicUserFollowingsHandler(classicUserFollowingsService, classicUserFollowersService, requestsService )
	recoveryPasswordTokenHandler := initRecoveryPasswordTokenHandler(recoveryPasswordTokenService,classicUserService,registeredUserService,userService, validator)
	classicUserHandler := initClassicUserHandler(classicUserService, settingsService, classicUserFollowingsService, requestsService)
	handleFunc(userHandler, confirmationTokenHandler, adminHandler,classicUserHandler, agentHandler,registeredUserHandler,classicUserCampaignsHandler,classicUserFollowingsHandler,classicUserFollowersHandler,recoveryPasswordTokenHandler)

}