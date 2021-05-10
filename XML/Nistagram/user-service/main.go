package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/rs/cors"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	db.AutoMigrate(&model.User{}, &model.RegisteredUser{}, &model.Admin{}, &model.ClassicUser{}, &model.Agent{},&model.RegisteredUserFollowers{},&model.RegisteredUserFollowings{}, &model.RegisteredUserCampaigns{}, &model.ConfirmationToken{})
	return db
}

func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}

func initAdminRepo(database *gorm.DB) *repository.AdminRepository{
	return &repository.AdminRepository { Database: database }
}

func initClassicUserRepo(database *gorm.DB) *repository.ClassicUserRepository{
	return &repository.ClassicUserRepository { Database: database }
}

func initAgentRepo(database *gorm.DB) *repository.AgentRepository{
	return &repository.AgentRepository { Database: database }
}

func initRegisteredUserCampaignsRepo(database *gorm.DB) *repository.RegisteredUserCampaignsRepository{
	return &repository.RegisteredUserCampaignsRepository { Database: database }
}

func initRegisteredUserFollowersRepo(database *gorm.DB) *repository.RegisteredUserFollowersRepository{
	return &repository.RegisteredUserFollowersRepository { Database: database }
}

func initRegisteredUserFollowingsRepo(database *gorm.DB) *repository.RegisteredUserFollowingsRepository{
	return &repository.RegisteredUserFollowingsRepository { Database: database }
}

func initConfirmationTokenRepo(database *gorm.DB) *repository.ConfirmationTokenRepository{
	return &repository.ConfirmationTokenRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}

func initAdminService(repo *repository.AdminRepository) *service.AdminService{
	return &service.AdminService { Repo: repo }
}

func initClassicUserService(repo *repository.ClassicUserRepository) *service.ClassicUserService{
	return &service.ClassicUserService { Repo: repo }
}

func initAgentService(repo *repository.AgentRepository) *service.AgentService{
	return &service.AgentService { Repo: repo }
}

func initRegisteredUserCampaignsService(repo *repository.RegisteredUserCampaignsRepository) *service.RegisteredUserCampaignsService{
	return &service.RegisteredUserCampaignsService { Repo: repo }
}

func initConfirmationTokenService(repo *repository.ConfirmationTokenRepository) *service.ConfirmationTokenService{
	return &service.ConfirmationTokenService { Repo: repo }
}

func initRegisteredUserFollowingsService(repo *repository.RegisteredUserFollowingsRepository) *service.RegisteredUserFollowingsService{
	return &service.RegisteredUserFollowingsService { Repo: repo }
}

func initRegisteredUserFollowersService(repo *repository.RegisteredUserFollowersRepository) *service.RegisteredUserFollowersService{
	return &service.RegisteredUserFollowersService { Repo: repo }
}

func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler { Service: service }
}

func initRegisteredUserHandler(service *service.RegisteredUserService) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler { Service: service }
}

func initAdminHandler(service *service.AdminService) *handler.AdminHandler{
	return &handler.AdminHandler { Service: service }
}

func initClassicUserHandler(classicUserService *service.ClassicUserService, userService *service.UserService, registeredUserService *service.RegisteredUserService,  confirmationTokenService *service.ConfirmationTokenService) *handler.ClassicUserHandler{
	return &handler.ClassicUserHandler { classicUserService, userService, registeredUserService , confirmationTokenService}
}

func initAgentHandler(service *service.AgentService) *handler.AgentHandler{
	return &handler.AgentHandler { Service: service }
}

func initRegisteredUserCampaignsHandler(service *service.RegisteredUserCampaignsService) *handler.RegisteredUserCampaignsHandler{
	return &handler.RegisteredUserCampaignsHandler { Service: service }
}

func initRegisteredUserFollowersHandler(service *service.RegisteredUserFollowersService) *handler.RegisteredUserFollowersHandler{
	return &handler.RegisteredUserFollowersHandler { Service: service }
}

func initRegisteredUserFollowingsHandler(service *service.RegisteredUserFollowingsService) *handler.RegisteredUserFollowingsHandler{
	return &handler.RegisteredUserFollowingsHandler { Service: service }
}

func initConfirmationTokenHandler(confirmationTokenService *service.ConfirmationTokenService, userService *service.UserService, registeredUserService *service.RegisteredUserService, classicUserService *service.ClassicUserService) *handler.ConfirmationTokenHandler{
	return &handler.ConfirmationTokenHandler{
		ConfirmationTokenService: confirmationTokenService,
		ClassicUserService:       classicUserService,
		RegisteredUserService:    registeredUserService,
		UserService:              userService,
	}
}

func handleFunc(userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler,adminHandler *handler.AdminHandler, agentHandler *handler.AgentHandler, classicUserHandler *handler.ClassicUserHandler,registeredUserCampaignsHandler *handler.RegisteredUserCampaignsHandler,registeredUserFollowingsHandler *handler.RegisteredUserFollowingsHandler,registeredUserFollowersHandler *handler.RegisteredUserFollowersHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login/", userHandler.LogIn).Methods("POST")
	router.HandleFunc("/confirm_registration/{confirmationToken}/{userId}", confirmationTokenHandler.VerifyConfirmationToken).Methods("POST")
	router.HandleFunc("/classic_user/", classicUserHandler.CreateClassicUser).Methods("POST")

	router.HandleFunc("/users/", userHandler.FindAllUsers).Methods("GET")
	router.HandleFunc("/admin/", adminHandler.CreateAdmin).Methods("POST")
	router.HandleFunc("/agent/", agentHandler.CreateAgent).Methods("POST")
	router.HandleFunc("/registered_user_campaigns/", registeredUserCampaignsHandler.CreateRegisteredUserCampaigns).Methods("POST")
	router.HandleFunc("/registered_user_followings/", registeredUserFollowingsHandler.CreateRegisteredUserFollowings).Methods("POST")
	router.HandleFunc("/registered_user_followers/", registeredUserFollowersHandler.CreateRegisteredUserFollowers).Methods("POST")


	mux := http.NewServeMux()
	mux.HandleFunc("/classic_user/", classicUserHandler.CreateClassicUser)
	mux.HandleFunc("/login/", userHandler.LogIn)
	mux.HandleFunc("/confirm_registration/{confirmationToken}/{userId}", confirmationTokenHandler.VerifyConfirmationToken)
	handler := cors.Default().Handler(mux)
	//go http.ListenAndServe(":8081", handler(router))
	log.Fatal(http.ListenAndServe(":8082", handler))
}


func main() {
	database := initDB()
	userRepo := initUserRepo(database)
	registeredUserRepo := initRegisteredUserRepo(database)
	adminRepo := initAdminRepo(database)
	classicUserRepo := initClassicUserRepo(database)
	agentRepo := initAgentRepo(database)
	confirmationTokenRepo := initConfirmationTokenRepo(database)
	registeredUserCampaignsRepo := initRegisteredUserCampaignsRepo(database)
	registeredUserFollowersRepo := initRegisteredUserFollowersRepo(database)
	registeredUserFollowingsRepo := initRegisteredUserFollowingsRepo(database)
	userService := initUserService(userRepo)
	registeredUserService := initRegisteredUserService(registeredUserRepo)
	confirmationTokenService := initConfirmationTokenService(confirmationTokenRepo)
	adminService := initAdminService(adminRepo)
	classicUserService := initClassicUserService(classicUserRepo)
	agentService := initAgentService(agentRepo)
	registeredUserCampaignsService := initRegisteredUserCampaignsService(registeredUserCampaignsRepo)
	registeredUserFollowersService := initRegisteredUserFollowersService(registeredUserFollowersRepo)
	registeredUserFollowingsService := initRegisteredUserFollowingsService(registeredUserFollowingsRepo)
	userHandler := initUserHandler(userService)
	//registeredUserHandler := initRegisteredUserHandler(registeredUserService)
	adminHandler := initAdminHandler(adminService)
	classicUserHandler := initClassicUserHandler(classicUserService, userService, registeredUserService,confirmationTokenService)
	agentHandler := initAgentHandler(agentService)
	confirmationTokenHandler := initConfirmationTokenHandler(confirmationTokenService,userService,registeredUserService,classicUserService)
	registeredUserCampaignsHandler := initRegisteredUserCampaignsHandler(registeredUserCampaignsService)
	registeredUserFollowersHandler := initRegisteredUserFollowersHandler(registeredUserFollowersService)
	registeredUserFollowingsHandler := initRegisteredUserFollowingsHandler(registeredUserFollowingsService)
	handleFunc(userHandler, confirmationTokenHandler, adminHandler,agentHandler,classicUserHandler,registeredUserCampaignsHandler,registeredUserFollowingsHandler,registeredUserFollowersHandler)
}