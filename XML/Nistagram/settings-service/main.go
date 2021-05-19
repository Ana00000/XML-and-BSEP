package main

import (
	_ "github.com/antchfx/xpath"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
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

	db.AutoMigrate(&model.ProfileSettings{},&model.ProfileSettingsRejectedMessageProfiles{},&model.ProfileSettingsApprovedMessageProfiles{},&model.ProfileSettingsMutedProfiles{},&model.ProfileSettingsBlockedProfiles{})
	return db
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

	mux := http.NewServeMux()

	mux.HandleFunc("/profile_settings/", handlerProfileSettings.CreateProfileSettings)
	mux.HandleFunc("/profile_settings_approved_message_profiles/", handlerProfileSettingsApprovedMessageProfiles.CreateProfileSettingsApprovedMessageProfiles)
	mux.HandleFunc("/profile_settings_blocked_profiles/", handlerProfileSettingsBlockedProfiles.CreateProfileSettingsBlockedProfiles)
	mux.HandleFunc("/profile_settings_muted_profiles/", handlerProfileSettingsMutedProfiles.CreateProfileSettingsMutedProfiles)
	mux.HandleFunc("/profile_settings_rejected_message_profiles/", handlerProfileSettingsRejectedMessageProfiles.CreateProfileSettingsRejectedMessageProfiles)
	mux.HandleFunc("/find_profile_settings_by_user_id", handlerProfileSettings.FindProfileSettingByUserId)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8088", handlerVar))
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