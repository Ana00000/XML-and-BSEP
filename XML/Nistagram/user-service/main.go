package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
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

	db.AutoMigrate(&model.User{}, &model.ClassicUser{}, &model.RegisteredUser{}, &model.Admin{}, &model.Agent{},&model.ClassicUserFollowers{},&model.ClassicUserFollowings{}, &model.ClassicUserCampaigns{}, &model.ConfirmationToken{})
	return db
}

//USER

func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler { Service: service }
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
	return &handler.ClassicUserHandler { Service: service }
}


//REGISTERED USER
func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}
func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}
func initRegisteredUserHandler(registeredUserService *service.RegisteredUserService, userService *service.UserService, classicUserService *service.ClassicUserService,  confirmationTokenService *service.ConfirmationTokenService) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler { registeredUserService, userService, classicUserService , confirmationTokenService}
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

func initAgentHandler(service *service.AgentService) *handler.AgentHandler{
	return &handler.AgentHandler { Service: service }
}

func initClassicUserCampaignsHandler(service *service.ClassicUserCampaignsService) *handler.ClassicUserCampaignsHandler{
	return &handler.ClassicUserCampaignsHandler { Service: service }
}

func initClassicUserFollowersHandler(service *service.ClassicUserFollowersService) *handler.ClassicUserFollowersHandler{
	return &handler.ClassicUserFollowersHandler { Service: service }
}

func initClassicUserFollowingsHandler(service *service.ClassicUserFollowingsService) *handler.ClassicUserFollowingsHandler{
	return &handler.ClassicUserFollowingsHandler { Service: service }
}

func initConfirmationTokenHandler(confirmationTokenService *service.ConfirmationTokenService, userService *service.UserService, registeredUserService *service.RegisteredUserService, classicUserService *service.ClassicUserService) *handler.ConfirmationTokenHandler{
	return &handler.ConfirmationTokenHandler{
		ConfirmationTokenService: confirmationTokenService,
		ClassicUserService:       classicUserService,
		RegisteredUserService:    registeredUserService,
		UserService:              userService,
	}
}

func handleFunc(userHandler *handler.UserHandler, confirmationTokenHandler *handler.ConfirmationTokenHandler, adminHandler *handler.AdminHandler, agentHandler *handler.AgentHandler, registeredUserHandler *handler.RegisteredUserHandler,classicUserCampaignsHandler *handler.ClassicUserCampaignsHandler,classicUserFollowingsHandler *handler.ClassicUserFollowingsHandler,classicUserFollowersHandler *handler.ClassicUserFollowersHandler){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login/", userHandler.LogIn).Methods("POST")
	router.HandleFunc("/users/", userHandler.FindAllUsers).Methods("GET")
	router.HandleFunc("/admin/", adminHandler.CreateAdmin).Methods("POST")
	router.HandleFunc("/agent/", agentHandler.CreateAgent).Methods("POST")
	router.HandleFunc("/confirm_registration/{confirmationToken}/{userId}", confirmationTokenHandler.VerifyConfirmationToken).Methods("POST")
	router.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/classic_user_campaigns/", classicUserCampaignsHandler.CreateClassicUserCampaigns).Methods("POST")
	router.HandleFunc("/classic_user_followings/", classicUserFollowingsHandler.CreateClassicUserFollowings).Methods("POST")
	router.HandleFunc("/classic_user_followers/", classicUserFollowersHandler.CreateClassicUserFollowers).Methods("POST")
	//missing for CLASSIC USER
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	userRepo := initUserRepo(database)
	registeredUserRepo := initRegisteredUserRepo(database)
	adminRepo := initAdminRepo(database)
	classicUserRepo := initClassicUserRepo(database)
	agentRepo := initAgentRepo(database)
	confirmationTokenRepo := initConfirmationTokenRepo(database)
	registeredUserCampaignsRepo := initClassicUserCampaignsRepo(database)
	registeredUserFollowersRepo := initClassicUserFollowersRepo(database)
	registeredUserFollowingsRepo := initClassicUserFollowingsRepo(database)

	userService := initUserService(userRepo)
	registeredUserService := initRegisteredUserService(registeredUserRepo)
	confirmationTokenService := initConfirmationTokenService(confirmationTokenRepo)
	adminService := initAdminService(adminRepo)
	classicUserService := initClassicUserService(classicUserRepo)
	agentService := initAgentService(agentRepo)
	registeredUserCampaignsService := initClassicUserCampaignsService(registeredUserCampaignsRepo)
	registeredUserFollowersService := initClassicUserFollowersService(registeredUserFollowersRepo)
	registeredUserFollowingsService := initClassicUserFollowingsService(registeredUserFollowingsRepo)


	userHandler := initUserHandler(userService)
	adminHandler := initAdminHandler(adminService)
	registeredUserHandler := initRegisteredUserHandler(registeredUserService, userService, classicUserService,confirmationTokenService)
	agentHandler := initAgentHandler(agentService)
	confirmationTokenHandler := initConfirmationTokenHandler(confirmationTokenService,userService,registeredUserService,classicUserService)
	registeredUserCampaignsHandler := initClassicUserCampaignsHandler(registeredUserCampaignsService)
	registeredUserFollowersHandler := initClassicUserFollowersHandler(registeredUserFollowersService)
	registeredUserFollowingsHandler := initClassicUserFollowingsHandler(registeredUserFollowingsService)
	handleFunc(userHandler, confirmationTokenHandler, adminHandler,agentHandler,registeredUserHandler,registeredUserCampaignsHandler,registeredUserFollowingsHandler,registeredUserFollowersHandler)
}