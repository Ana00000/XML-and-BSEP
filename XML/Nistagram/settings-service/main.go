package main

import (
	"fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
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

	db.AutoMigrate(&model.ProfileSettings{},&model.ProfileSettingsRejectedMessageProfiles{},&model.ProfileSettingsApprovedMessageProfiles{},&model.ProfileSettingsMutedProfiles{},&model.ProfileSettingsBlockedProfiles{})
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

func initProfileSettingsRepo(database *gorm.DB) *repository.ProfileSettingsRepository{
	return &repository.ProfileSettingsRepository { Database: database }
}

func initProfileSettingsServices(repo *repository.ProfileSettingsRepository) *service.ProfileSettingsService{
	return &service.ProfileSettingsService { Repo: repo }
}

func initProfileSettingsHandler(service *service.ProfileSettingsService) *handler.ProfileSettingsHandler{
	return &handler.ProfileSettingsHandler { Service: service }
}

func initProfileSettingsRejectedMessageProfilesRepo(database *gorm.DB) *repository.ProfileSettingsRejectedMessageProfilesRepository{
	return &repository.ProfileSettingsRejectedMessageProfilesRepository { Database: database }
}

func initProfileSettingsRejectedMessageProfilesServices(repo *repository.ProfileSettingsRejectedMessageProfilesRepository) *service.ProfileSettingsRejectedMessageProfilesService{
	return &service.ProfileSettingsRejectedMessageProfilesService { Repo: repo }
}

func initProfileSettingsRejectedMessageProfilesHandler(service *service.ProfileSettingsRejectedMessageProfilesService) *handler.ProfileSettingsRejectedMessageProfilesHandler{
	return &handler.ProfileSettingsRejectedMessageProfilesHandler { Service: service }
}

func initProfileSettingsApprovedMessageProfilesRepo(database *gorm.DB) *repository.ProfileSettingsApprovedMessageProfilesRepository{
	return &repository.ProfileSettingsApprovedMessageProfilesRepository { Database: database }
}

func initProfileSettingsApprovedMessageProfilesServices(repo *repository.ProfileSettingsApprovedMessageProfilesRepository) *service.ProfileSettingsApprovedMessageProfilesService{
	return &service.ProfileSettingsApprovedMessageProfilesService { Repo: repo }
}

func initProfileSettingsApprovedMessageProfilesHandler(service *service.ProfileSettingsApprovedMessageProfilesService) *handler.ProfileSettingsApprovedMessageProfilesHandler{
	return &handler.ProfileSettingsApprovedMessageProfilesHandler { Service: service }
}

func initProfileSettingsMutedProfilesRepo(database *gorm.DB) *repository.ProfileSettingsMutedProfilesRepository{
	return &repository.ProfileSettingsMutedProfilesRepository { Database: database }
}

func initProfileSettingsMutedProfilesServices(repo *repository.ProfileSettingsMutedProfilesRepository) *service.ProfileSettingsMutedProfilesService{
	return &service.ProfileSettingsMutedProfilesService { Repo: repo }
}

func initProfileSettingsMutedProfilesHandler(service *service.ProfileSettingsMutedProfilesService) *handler.ProfileSettingsMutedProfilesHandler{
	return &handler.ProfileSettingsMutedProfilesHandler { Service: service }
}

func initProfileSettingsBlockedProfilesRepo(database *gorm.DB) *repository.ProfileSettingsBlockedProfilesRepository{
	return &repository.ProfileSettingsBlockedProfilesRepository { Database: database }
}

func initProfileSettingsBlockedProfilesServices(repo *repository.ProfileSettingsBlockedProfilesRepository) *service.ProfileSettingsBlockedProfilesService{
	return &service.ProfileSettingsBlockedProfilesService { Repo: repo }
}

func initProfileSettingsBlockedProfilesHandler(service *service.ProfileSettingsBlockedProfilesService) *handler.ProfileSettingsBlockedProfilesHandler{
	return &handler.ProfileSettingsBlockedProfilesHandler { Service: service }
}

func handleFunc(handlerProfileSettings *handler.ProfileSettingsHandler, handlerProfileSettingsApprovedMessageProfiles *handler.ProfileSettingsApprovedMessageProfilesHandler,
	handlerProfileSettingsBlockedProfiles *handler.ProfileSettingsBlockedProfilesHandler, handlerProfileSettingsMutedProfiles *handler.ProfileSettingsMutedProfilesHandler,
	handlerProfileSettingsRejectedMessageProfiles *handler.ProfileSettingsRejectedMessageProfilesHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/profile_settings/{userID}", handlerProfileSettings.CreateProfileSettings).Methods("POST")
	router.HandleFunc("/profile_settings_approved_message_profiles/", handlerProfileSettingsApprovedMessageProfiles.CreateProfileSettingsApprovedMessageProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_blocked_profiles/", handlerProfileSettingsBlockedProfiles.CreateProfileSettingsBlockedProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_muted_profiles/", handlerProfileSettingsMutedProfiles.CreateProfileSettingsMutedProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_rejected_message_profiles/", handlerProfileSettingsRejectedMessageProfiles.CreateProfileSettingsRejectedMessageProfiles).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	database := initDB()
	repoProfileSettings := initProfileSettingsRepo(database)
	serviceProfileSettings := initProfileSettingsServices(repoProfileSettings)
	handlerProfileSettings := initProfileSettingsHandler(serviceProfileSettings)

	repoProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesRepo(database)
	serviceProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesServices(repoProfileSettingsApprovedMessageProfiles)
	handlerProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesHandler(serviceProfileSettingsApprovedMessageProfiles)

	repoProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesRepo(database)
	serviceProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesServices(repoProfileSettingsBlockedProfiles)
	handlerProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesHandler(serviceProfileSettingsBlockedProfiles)

	repoProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesRepo(database)
	serviceProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesServices(repoProfileSettingsMutedProfiles)
	handlerProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesHandler(serviceProfileSettingsMutedProfiles)

	repoProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesRepo(database)
	serviceProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesServices(repoProfileSettingsRejectedMessageProfiles)
	handlerProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesHandler(serviceProfileSettingsRejectedMessageProfiles)
	handleFunc(handlerProfileSettings,handlerProfileSettingsApprovedMessageProfiles,handlerProfileSettingsBlockedProfiles,handlerProfileSettingsMutedProfiles,handlerProfileSettingsRejectedMessageProfiles)
}