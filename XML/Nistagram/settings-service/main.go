package main

import (
	"fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB {
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.ProfileSettings{}, &model.ProfileSettingsRejectedMessageProfiles{}, &model.ProfileSettingsApprovedMessageProfiles{}, &model.ProfileSettingsMutedProfiles{}, &model.ProfileSettingsBlockedProfiles{},  &model.ProfileSettingsPostNotificationsProfiles{},  &model.ProfileSettingsStoryNotificationsProfiles{})
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

func initProfileSettingsRepo(database *gorm.DB,  profileSettingsPostNotificationsProfilesRepository *repository.ProfileSettingsPostNotificationsProfilesRepository) *repository.ProfileSettingsRepository {
	return &repository.ProfileSettingsRepository{Database: database, ProfileSettingsPostNotificationsProfilesRepository: profileSettingsPostNotificationsProfilesRepository}
}

func initProfileSettingsServices(repo *repository.ProfileSettingsRepository) *service.ProfileSettingsService {
	return &service.ProfileSettingsService{Repo: repo}
}

func initProfileSettingsHandler(service *service.ProfileSettingsService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsHandler {
	return &handler.ProfileSettingsHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initProfileSettingsRejectedMessageProfilesRepo(database *gorm.DB) *repository.ProfileSettingsRejectedMessageProfilesRepository {
	return &repository.ProfileSettingsRejectedMessageProfilesRepository{Database: database}
}

func initProfileSettingsRejectedMessageProfilesServices(repo *repository.ProfileSettingsRejectedMessageProfilesRepository) *service.ProfileSettingsRejectedMessageProfilesService {
	return &service.ProfileSettingsRejectedMessageProfilesService{Repo: repo}
}

func initProfileSettingsRejectedMessageProfilesHandler(service *service.ProfileSettingsRejectedMessageProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsRejectedMessageProfilesHandler {
	return &handler.ProfileSettingsRejectedMessageProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initProfileSettingsApprovedMessageProfilesRepo(database *gorm.DB) *repository.ProfileSettingsApprovedMessageProfilesRepository {
	return &repository.ProfileSettingsApprovedMessageProfilesRepository{Database: database }
}

func initProfileSettingsApprovedMessageProfilesServices(repo *repository.ProfileSettingsApprovedMessageProfilesRepository) *service.ProfileSettingsApprovedMessageProfilesService {
	return &service.ProfileSettingsApprovedMessageProfilesService{Repo: repo}
}

func initProfileSettingsApprovedMessageProfilesHandler(service *service.ProfileSettingsApprovedMessageProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsApprovedMessageProfilesHandler {
	return &handler.ProfileSettingsApprovedMessageProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initProfileSettingsMutedProfilesRepo(database *gorm.DB) *repository.ProfileSettingsMutedProfilesRepository {
	return &repository.ProfileSettingsMutedProfilesRepository{Database: database}
}

func initProfileSettingsMutedProfilesServices(repo *repository.ProfileSettingsMutedProfilesRepository) *service.ProfileSettingsMutedProfilesService {
	return &service.ProfileSettingsMutedProfilesService{Repo: repo}
}

func initProfileSettingsMutedProfilesHandler(service *service.ProfileSettingsMutedProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsMutedProfilesHandler {
	return &handler.ProfileSettingsMutedProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initProfileSettingsBlockedProfilesRepo(database *gorm.DB) *repository.ProfileSettingsBlockedProfilesRepository {
	return &repository.ProfileSettingsBlockedProfilesRepository{Database: database}
}

func initProfileSettingsBlockedProfilesServices(repo *repository.ProfileSettingsBlockedProfilesRepository) *service.ProfileSettingsBlockedProfilesService {
	return &service.ProfileSettingsBlockedProfilesService{Repo: repo}
}

func initProfileSettingsBlockedProfilesHandler(service *service.ProfileSettingsBlockedProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsBlockedProfilesHandler {
	return &handler.ProfileSettingsBlockedProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}


func initProfileSettingsPostNotificationsProfilesRepo(database *gorm.DB) *repository.ProfileSettingsPostNotificationsProfilesRepository {
	return &repository.ProfileSettingsPostNotificationsProfilesRepository{Database: database}
}

func initProfileSettingsPostNotificationsProfilesServices(repo *repository.ProfileSettingsPostNotificationsProfilesRepository) *service.ProfileSettingsPostNotificationsProfilesService {
	return &service.ProfileSettingsPostNotificationsProfilesService{Repo: repo}
}

func initProfileSettingsPostNotificationsProfilesHandler(service *service.ProfileSettingsPostNotificationsProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsPostNotificationsProfilesHandler {
	return &handler.ProfileSettingsPostNotificationsProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func initProfileSettingsStoryNotificationsProfilesRepo(database *gorm.DB) *repository.ProfileSettingsStoryNotificationsProfilesRepository {
	return &repository.ProfileSettingsStoryNotificationsProfilesRepository{Database: database}
}

func initProfileSettingsStoryNotificationsProfilesServices(repo *repository.ProfileSettingsStoryNotificationsProfilesRepository) *service.ProfileSettingsStoryNotificationsProfilesService {
	return &service.ProfileSettingsStoryNotificationsProfilesService{Repo: repo}
}

func initProfileSettingsStoryNotificationsProfilesHandler(service *service.ProfileSettingsStoryNotificationsProfilesService, LogInfo *logrus.Logger, LogError *logrus.Logger, validator *validator.Validate) *handler.ProfileSettingsStoryNotificationsProfilesHandler {
	return &handler.ProfileSettingsStoryNotificationsProfilesHandler{
		Service:   service,
		LogInfo:   LogInfo,
		LogError:  LogError,
		Validator: validator,
	}
}

func handleFunc(handlerProfileSettings *handler.ProfileSettingsHandler, handlerProfileSettingsApprovedMessageProfiles *handler.ProfileSettingsApprovedMessageProfilesHandler,
	handlerProfileSettingsBlockedProfiles *handler.ProfileSettingsBlockedProfilesHandler, handlerProfileSettingsMutedProfiles *handler.ProfileSettingsMutedProfilesHandler,
	handlerProfileSettingsRejectedMessageProfiles *handler.ProfileSettingsRejectedMessageProfilesHandler,
	handlerProfileSettingsPostNotificationsProfiles *handler.ProfileSettingsPostNotificationsProfilesHandler,
	handlerProfileSettingsStoryNotificationsProfiles *handler.ProfileSettingsStoryNotificationsProfilesHandler) {

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/profile_settings/{userID}", handlerProfileSettings.CreateProfileSettings).Methods("POST")
	router.HandleFunc("/profile_settings_approved_message_profiles/", handlerProfileSettingsApprovedMessageProfiles.CreateProfileSettingsApprovedMessageProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_blocked_profiles/", handlerProfileSettingsBlockedProfiles.CreateProfileSettingsBlockedProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_muted_profiles/", handlerProfileSettingsMutedProfiles.CreateProfileSettingsMutedProfiles).Methods("POST")
	router.HandleFunc("/profile_settings_rejected_message_profiles/", handlerProfileSettingsRejectedMessageProfiles.CreateProfileSettingsRejectedMessageProfiles).Methods("POST")
	router.HandleFunc("/find_profile_settings_by_user_id/{userID}", handlerProfileSettings.FindProfileSettingByUserId).Methods("GET")
	router.HandleFunc("/find_all_for_public_users/", handlerProfileSettings.FindProfileSettingsForPublicUsers).Methods("GET")
	router.HandleFunc("/find_all_public_users/", handlerProfileSettings.FindAllPublicUsers).Methods("POST")
	router.HandleFunc("/update_profile_settings/", handlerProfileSettings.UpdateProfileSettings).Methods("POST")
	router.HandleFunc("/find_profile_settings_with_id_by_user_id/{userID}", handlerProfileSettings.FindProfileSettingWithIDByUserId).Methods("GET")
	router.HandleFunc("/add_post_notifications_for_user/", handlerProfileSettingsPostNotificationsProfiles.CreateProfileSettingsPostNotificationsProfiles).Methods("POST")
	router.HandleFunc("/add_story_notifications_for_user/", handlerProfileSettingsStoryNotificationsProfiles.CreateProfileSettingsStoryNotificationsProfiles).Methods("POST")

	router.HandleFunc("/find_all_users_for_post_notifications/{userID}", handlerProfileSettings.FindAllUsersForPostNotifications).Methods("GET")
	router.HandleFunc("/find_all_users_for_post_album_notifications/{userID}", handlerProfileSettings.FindAllUsersForPostAlbumNotifications).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()
	validator := validator.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoSETTINGS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorSETTINGS.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()


	repoProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesRepo(database)
	serviceProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesServices(repoProfileSettingsBlockedProfiles)
	handlerProfileSettingsBlockedProfiles := initProfileSettingsBlockedProfilesHandler(serviceProfileSettingsBlockedProfiles, logInfo, logError, validator)

	repoProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesRepo(database)
	serviceProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesServices(repoProfileSettingsMutedProfiles)
	handlerProfileSettingsMutedProfiles := initProfileSettingsMutedProfilesHandler(serviceProfileSettingsMutedProfiles, logInfo, logError, validator)

	repoProfileSettingsPostNotificationsProfiles := initProfileSettingsPostNotificationsProfilesRepo(database)
	serviceProfileSettingsPostNotificationsProfiles := initProfileSettingsPostNotificationsProfilesServices(repoProfileSettingsPostNotificationsProfiles)
	handlerProfileSettingsPostNotificationsProfiles := initProfileSettingsPostNotificationsProfilesHandler(serviceProfileSettingsPostNotificationsProfiles, logInfo, logError, validator)


	repoProfileSettings := initProfileSettingsRepo(database, repoProfileSettingsPostNotificationsProfiles)
	serviceProfileSettings := initProfileSettingsServices(repoProfileSettings)
	handlerProfileSettings := initProfileSettingsHandler(serviceProfileSettings, logInfo, logError, validator)

	repoProfileSettingsStoryNotificationsProfiles := initProfileSettingsStoryNotificationsProfilesRepo(database)
	serviceProfileSettingsStoryNotificationsProfiles := initProfileSettingsStoryNotificationsProfilesServices(repoProfileSettingsStoryNotificationsProfiles)
	handlerProfileSettingsStoryNotificationsProfiles := initProfileSettingsStoryNotificationsProfilesHandler(serviceProfileSettingsStoryNotificationsProfiles, logInfo, logError, validator)


	repoProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesRepo(database)
	serviceProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesServices(repoProfileSettingsApprovedMessageProfiles)
	handlerProfileSettingsApprovedMessageProfiles := initProfileSettingsApprovedMessageProfilesHandler(serviceProfileSettingsApprovedMessageProfiles, logInfo, logError, validator)

	repoProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesRepo(database)
	serviceProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesServices(repoProfileSettingsRejectedMessageProfiles)
	handlerProfileSettingsRejectedMessageProfiles := initProfileSettingsRejectedMessageProfilesHandler(serviceProfileSettingsRejectedMessageProfiles, logInfo, logError, validator)
	handleFunc(handlerProfileSettings, handlerProfileSettingsApprovedMessageProfiles, handlerProfileSettingsBlockedProfiles, handlerProfileSettingsMutedProfiles, handlerProfileSettingsRejectedMessageProfiles, handlerProfileSettingsPostNotificationsProfiles, handlerProfileSettingsStoryNotificationsProfiles)
}
