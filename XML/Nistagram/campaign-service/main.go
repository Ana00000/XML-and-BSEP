package main

import (
	"fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/campaign-service/service"
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

	db.AutoMigrate(&model.Campaign{}, &model.Advertisement{}, &model.MultiUseCampaign{}, &model.DisposableCampaign{}, &model.CampaignChosenGroup{})
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

func initCampaignChosenGroupRepo(database *gorm.DB) *repository.CampaignChosenGroupRepository{
	return &repository.CampaignChosenGroupRepository { Database: database }
}

func initCampaignChosenGroupServices(repo *repository.CampaignChosenGroupRepository) *service.CampaignChosenGroupService{
	return &service.CampaignChosenGroupService { Repo: repo }
}

func initCampaignChosenGroupHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.CampaignChosenGroupService) *handler.CampaignChosenGroupHandler{
	return &handler.CampaignChosenGroupHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initCampaignRepo(database *gorm.DB) *repository.CampaignRepository{
	return &repository.CampaignRepository { Database: database }
}

func initDisposableCampaignRepo(database *gorm.DB) *repository.DisposableCampaignRepository{
	return &repository.DisposableCampaignRepository { Database: database }
}

func initMultiUseCampaignRepo(database *gorm.DB) *repository.MultiUseCampaignRepository{
	return &repository.MultiUseCampaignRepository { Database: database }
}

func initAdvertisementRepo(database *gorm.DB) *repository.AdvertisementRepository{
	return &repository.AdvertisementRepository { Database: database }
}

func initAdvertisementServices(repo *repository.AdvertisementRepository) *service.AdvertisementService{
	return &service.AdvertisementService { Repo: repo }
}

func initMultiUseCampaignServices(repo *repository.MultiUseCampaignRepository) *service.MultiUseCampaignService{
	return &service.MultiUseCampaignService { Repo: repo }
}

func initDisposableCampaignServices(repo *repository.DisposableCampaignRepository) *service.DisposableCampaignService{
	return &service.DisposableCampaignService { Repo: repo }
}

func initCampaignServices(repo *repository.CampaignRepository) *service.CampaignService{
	return &service.CampaignService { Repo: repo }
}

func initAdvertisementHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.AdvertisementService) *handler.AdvertisementHandler{
	return &handler.AdvertisementHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initCampaignHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.CampaignService) *handler.CampaignHandler{
	return &handler.CampaignHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initDisposableCampaignHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.DisposableCampaignService) *handler.DisposableCampaignHandler{
	return &handler.DisposableCampaignHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initMultiUseCampaignHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.MultiUseCampaignService) *handler.MultiUseCampaignHandler{
	return &handler.MultiUseCampaignHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}



func handleFunc(handlerMultiUseCampaign *handler.MultiUseCampaignHandler,handlerDisposableCampaign *handler.DisposableCampaignHandler,handlerCampaign *handler.CampaignHandler,handlerAdvertisement *handler.AdvertisementHandler,
	handlerCampaignChosenGroup *handler.CampaignChosenGroupHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/multi_use_campaign/", handlerMultiUseCampaign.CreateMultiUseCampaign).Methods("POST")
	router.HandleFunc("/disposable_campaign/", handlerDisposableCampaign.CreateDisposableCampaign).Methods("POST")
	router.HandleFunc("/campaign/", handlerCampaign.CreateCampaign).Methods("POST")
	router.HandleFunc("/advertisement/", handlerAdvertisement.CreateAdvertisement).Methods("POST")
	router.HandleFunc("/campaign_chosen_group/", handlerCampaignChosenGroup.CreateCampaignChosenGroup).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")),cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoCAMPAIGN.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorCAMPAIGN.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()

	repoMultiUseCampaign := initMultiUseCampaignRepo(database)
	repoDisposableCampaign := initDisposableCampaignRepo(database)
	repoCampaign := initCampaignRepo(database)
	repoAdvertisement := initAdvertisementRepo(database)

	serviceMultiUseCampaign := initMultiUseCampaignServices(repoMultiUseCampaign)
	serviceDisposableCampaign := initDisposableCampaignServices(repoDisposableCampaign)
	serviceCampaign := initCampaignServices(repoCampaign)
	serviceAdvertisement := initAdvertisementServices(repoAdvertisement)

	handlerMultiUseCampaign := initMultiUseCampaignHandler(logInfo,logError,serviceMultiUseCampaign)
	handlerDisposableCampaign := initDisposableCampaignHandler(logInfo,logError,serviceDisposableCampaign)
	handlerCampaign := initCampaignHandler(logInfo,logError,serviceCampaign)
	handlerAdvertisement := initAdvertisementHandler(logInfo,logError,serviceAdvertisement)

	repoCampaignChosenGroup := initCampaignChosenGroupRepo(database)
	serviceCampaignChosenGroup := initCampaignChosenGroupServices(repoCampaignChosenGroup)
	handlerCampaignChosenGroup := initCampaignChosenGroupHandler(logInfo,logError,serviceCampaignChosenGroup)
	handleFunc(handlerMultiUseCampaign,handlerDisposableCampaign,handlerCampaign,handlerAdvertisement,handlerCampaignChosenGroup)
}