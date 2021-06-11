package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
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

	db.AutoMigrate(&model.Location{})
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

func initRepo(database *gorm.DB) *repository.LocationRepository{
	return &repository.LocationRepository { Database: database }
}

func initServices(repo *repository.LocationRepository) *service.LocationService{
	return &service.LocationService { Repo: repo }
}

func initHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.LocationService) *handler.LocationHandler{
	return &handler.LocationHandler { LogInfo: LogInfo, LogError: LogError, Service: service }
}

func handleFunc(handler *handler.LocationHandler){
	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/", handler.CreateLocation).Methods("POST")
	router.HandleFunc("/find_location_by_id", handler.FindByID).Methods("GET")
	router.HandleFunc("/find_locations_for_stories/", handler.FindAllLocationsForStories).Methods("POST")
	router.HandleFunc("/find_locations_for_story/", handler.FindAllLocationsForStory).Methods("POST")

	router.HandleFunc("/find_locations_for_posts/", handler.FindAllLocationsForPosts).Methods("POST")
	router.HandleFunc("/find_locations_for_post/", handler.FindAllLocationsForPost).Methods("POST")


	router.HandleFunc("/find_location_id_by_location_string/{locationString}", handler.FindLocationIdByLocationString).Methods("GET")

	router.HandleFunc("/find_locations_for_post_albums/", handler.FindAllLocationsForPostAlbums).Methods("POST")
	router.HandleFunc("/find_locations_for_post_album/", handler.FindAllLocationsForPostAlbum).Methods("POST")

	router.HandleFunc("/find_locations_for_story_albums/", handler.FindAllLocationsForStoryAlbums).Methods("POST")
	router.HandleFunc("/find_locations_for_story_album/", handler.FindAllLocationsForStoryAlbum).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoLOCATION.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorLOCATION.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(logInfo,logError,service)
	handleFunc(handler)
}