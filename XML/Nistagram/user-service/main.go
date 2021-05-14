package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mikespook/gorbac"
	_ "github.com/mikespook/gorbac"
	"github.com/rs/cors"
	settingsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
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

//USER

func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initUserHandler(UserService *service.UserService,AdminService *service.AdminService, ClassicUserService *service.ClassicUserService, RegisteredUserService *service.RegisteredUserService, AgentService *service.AgentService, rbac *gorbac.RBAC, permissionFindAllUsers *gorbac.Permission, permissionUpdateUserInfo *gorbac.Permission ) *handler.UserHandler{
	return &handler.UserHandler{
		UserService:            UserService,
		AdminService:           AdminService,
		ClassicUserService:     ClassicUserService,
		RegisteredUserService:  RegisteredUserService,
		AgentService:           AgentService,
		Rbac:                   rbac,
		PermissionFindAllUsers: permissionFindAllUsers,
		PermissionUpdateUserInfo: permissionUpdateUserInfo,
	}
}
//SETTINGS
func initSettingsRepo(database *gorm.DB) *settingsRepository.ProfileSettingsRepository{
	return &settingsRepository.ProfileSettingsRepository { Database: database }
}

func initSettingsService(repo *settingsRepository.ProfileSettingsRepository) *settingsService.ProfileSettingsService{
	return &settingsService.ProfileSettingsService { Repo: repo }
}


//ADMIN
func initAdminRepo(database *gorm.DB) *repository.AdminRepository{
	return &repository.AdminRepository { Database: database }
}

func initAdminService(repo *repository.AdminRepository) *service.AdminService{
	return &service.AdminService { Repo: repo }
}

func initAdminHandler(service *service.AdminService) *handler.AdminHandler{
	return &handler.AdminHandler { Service: service }
}

//CLASSIC USER
func initClassicUserRepo(database *gorm.DB) *repository.ClassicUserRepository{
	return &repository.ClassicUserRepository { Database: database }
}

func initClassicUserService(repo *repository.ClassicUserRepository) *service.ClassicUserService{
	return &service.ClassicUserService { Repo: repo }
}

func initClassicUserHandler(service *service.ClassicUserService) *handler.ClassicUserHandler{
	return &handler.ClassicUserHandler { ClassicUserService: service }
}


//REGISTERED USER
func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}
func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}
func initRegisteredUserHandler(registeredUserService *service.RegisteredUserService, userService *service.UserService, classicUserService *service.ClassicUserService,  confirmationTokenService *service.ConfirmationTokenService, settingsService *settingsService.ProfileSettingsService) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler { RegisteredUserService: registeredUserService, UserService: userService, ClassicUserService: classicUserService , ConfirmationTokenService: confirmationTokenService, ProfileSettingsService: settingsService}
}

func initAgentRepo(database *gorm.DB) *repository.AgentRepository{
	return &repository.AgentRepository { Database: database }
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

func initAgentService(repo *repository.AgentRepository) *service.AgentService{
	return &service.AgentService { Repo: repo }
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

func initAgentHandler(service *service.AgentService) *handler.AgentHandler{
	return &handler.AgentHandler { Service: service }
}

func initClassicUserCampaignsHandler(service *service.ClassicUserCampaignsService) *handler.ClassicUserCampaignsHandler{
	return &handler.ClassicUserCampaignsHandler { Service: service }
}

func initClassicUserFollowersHandler(service *service.ClassicUserFollowersService, userService *service.UserService) *handler.ClassicUserFollowersHandler{
	return &handler.ClassicUserFollowersHandler { ClassicUserFollowersService: service, UserService: userService}
}

func initClassicUserFollowingsHandler(service *service.ClassicUserFollowingsService) *handler.ClassicUserFollowingsHandler{
	return &handler.ClassicUserFollowingsHandler { Service: service }
}

func initRecoveryPasswordTokenHandler(recoveryPasswordTokenService *service.RecoveryPasswordTokenService, classicUserService *service.ClassicUserService, registeredUserService *service.RegisteredUserService, userService *service.UserService) *handler.RecoveryPasswordTokenHandler{
	return &handler.RecoveryPasswordTokenHandler{
		RecoveryPasswordTokenService: recoveryPasswordTokenService,
		UserService:                  userService,
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

	router.HandleFunc("/login/", userHandler.LogIn).Methods("POST")
	router.HandleFunc("/confirm_registration/", confirmationTokenHandler.VerifyConfirmationToken).Methods("POST")
	//router.HandleFunc("/classic_user/", classicUserHandler.CreateClassicUser).Methods("POST")

	router.HandleFunc("/users/", userHandler.FindAllUsers).Methods("GET")
	router.HandleFunc("/admin/", adminHandler.CreateAdmin).Methods("POST")
	router.HandleFunc("/agent/", agentHandler.CreateAgent).Methods("POST")
	router.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/classic_user_campaigns/", classicUserCampaignsHandler.CreateClassicUserCampaigns).Methods("POST")
	router.HandleFunc("/classic_user_followings/", classicUserFollowingsHandler.CreateClassicUserFollowings).Methods("POST")
	router.HandleFunc("/classic_user_followers/", classicUserFollowersHandler.CreateClassicUserFollowers).Methods("POST")
	
	mux := http.NewServeMux()
	mux.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser)
	mux.HandleFunc("/login/", userHandler.LogIn)
	mux.HandleFunc("/recovery_password/", recoveryPasswordTokenHandler.GenerateRecoveryPasswordToken)
	mux.HandleFunc("/verify_recovery_password_token/", recoveryPasswordTokenHandler.VerifyRecoveryPasswordToken)
	mux.HandleFunc("/confirm_registration/", confirmationTokenHandler.VerifyConfirmationToken)
	mux.HandleFunc("/change_user_password/", userHandler.ChangeUserPassword)
	mux.HandleFunc("/users/all",userHandler.FindAllUsers)
	mux.HandleFunc("/find_all_followers_for_user",classicUserFollowersHandler.FindAllFollowersInfoForUser)
	mux.HandleFunc("/create_follower/",classicUserFollowersHandler.CreateClassicUserFollowers)
	mux.HandleFunc("/update_user_profile_info/", userHandler.UpdateUserProfileInfo)
	mux.HandleFunc("/find_user_by_id", userHandler.FindByID)
	mux.HandleFunc("/find_user_by_username", userHandler.FindByUserName)
	mux.HandleFunc("/find_all_users_but_logged_in", userHandler.FindAllUsersButLoggedIn)
	mux.HandleFunc("/find_selected_user_by_id", classicUserHandler.FindSelectedUserById)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handlerVar))
}

func main() {
	rbac := gorbac.New()

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

	userHandler := initUserHandler(userService,adminService,classicUserService,registeredUserService,agentService, rbac, &permissionFindAllUsers, &permissionUpdateUserInfo)
	adminHandler := initAdminHandler(adminService)
	registeredUserHandler := initRegisteredUserHandler(registeredUserService, userService, classicUserService,confirmationTokenService, settingsService)
	agentHandler := initAgentHandler(agentService)
	confirmationTokenHandler := initConfirmationTokenHandler(confirmationTokenService,userService,registeredUserService,classicUserService)
	classicUserCampaignsHandler := initClassicUserCampaignsHandler(classicUserCampaignsService)
	classicUserFollowersHandler := initClassicUserFollowersHandler(classicUserFollowersService, userService)
	classicUserFollowingsHandler := initClassicUserFollowingsHandler(classicUserFollowingsService)
	recoveryPasswordTokenHandler := initRecoveryPasswordTokenHandler(recoveryPasswordTokenService,classicUserService,registeredUserService,userService)
	classicUserHandler := initClassicUserHandler(classicUserService)
	handleFunc(userHandler, confirmationTokenHandler, adminHandler,classicUserHandler, agentHandler,registeredUserHandler,classicUserCampaignsHandler,classicUserFollowingsHandler,classicUserFollowersHandler,recoveryPasswordTokenHandler)
}