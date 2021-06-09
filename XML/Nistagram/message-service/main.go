package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/message-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func initDB() *gorm.DB{
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.MessageSubstance{}, &model.Message{}, &model.StoryMessageSubstance{}, &model.PostMessageSubstance{})
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

func initMessageRepo(database *gorm.DB) *repository.MessageRepository{
	return &repository.MessageRepository{ Database: database }
}

func initMessageSubstanceRepo(database *gorm.DB) *repository.MessageSubstanceRepository{
	return &repository.MessageSubstanceRepository{ Database: database }
}

func initStoryMessageSubstanceRepo(database *gorm.DB) *repository.StoryMessageSubstanceRepository{
	return &repository.StoryMessageSubstanceRepository{ Database: database }
}

func initPostMessageSubstanceRepo(database *gorm.DB) *repository.PostMessageSubstanceRepository{
	return &repository.PostMessageSubstanceRepository{ Database: database }
}

func initMessageService(repo *repository.MessageRepository) *service.MessageService{
	return &service.MessageService{ Repo: repo }
}

func initMessageSubstanceService(repo *repository.MessageSubstanceRepository) *service.MessageSubstanceService{
	return &service.MessageSubstanceService{ Repo: repo }
}

func initStoryMessageSubstanceService(repo *repository.StoryMessageSubstanceRepository) *service.StoryMessageSubstanceService{
	return &service.StoryMessageSubstanceService{ Repo: repo }
}

func initPostMessageSubstanceService(repo *repository.PostMessageSubstanceRepository) *service.PostMessageSubstanceService{
	return &service.PostMessageSubstanceService{ Repo: repo }
}

func initMessageHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.MessageService) *handler.MessageHandler{
	return &handler.MessageHandler{ LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initMessageSubstanceHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.MessageSubstanceService) *handler.MessageSubstanceHandler{
	return &handler.MessageSubstanceHandler{ LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initStoryMessageSubstanceHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.StoryMessageSubstanceService) *handler.StoryMessageSubstanceHandler{
	return &handler.StoryMessageSubstanceHandler{ LogInfo: LogInfo, LogError: LogError, Service: service }
}

func initPostMessageSubstanceHandler(LogInfo *logrus.Logger,LogError *logrus.Logger,service *service.PostMessageSubstanceService) *handler.PostMessageSubstanceHandler{
	return &handler.PostMessageSubstanceHandler{ LogInfo: LogInfo, LogError: LogError, Service: service }
}

func handleFunc(handlerMessage *handler.MessageHandler, handlerMessageSubstance *handler.MessageSubstanceHandler,
				handlerStoryMessageSubstance *handler.StoryMessageSubstanceHandler, handlerPostMessageSubstance *handler.PostMessageSubstanceHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/message/", handlerMessage.CreateMessage).Methods("POST")
	router.HandleFunc("/message_content/", handlerMessageSubstance.CreateMessageSubstance).Methods("POST")
	router.HandleFunc("/story_message_content/", handlerStoryMessageSubstance.CreateStoryMessageSubstance).Methods("POST")
	router.HandleFunc("/post_message_content/", handlerPostMessageSubstance.CreatePostMessageSubstance).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()

	LogInfoFile, err := os.OpenFile("logInfo.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile("logError.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()
	repoMessage := initMessageRepo(database)
	repoMessageSubstance := initMessageSubstanceRepo(database)
	repoStoryMessageSubstance := initStoryMessageSubstanceRepo(database)
	repoPostMessageSubstance := initPostMessageSubstanceRepo(database)
	serviceMessage := initMessageService(repoMessage)
	serviceMessageSubstance := initMessageSubstanceService(repoMessageSubstance)
	serviceStoryMessageSubstance := initStoryMessageSubstanceService(repoStoryMessageSubstance)
	servicePostMessageSubstance := initPostMessageSubstanceService(repoPostMessageSubstance)
	handlerMessage := initMessageHandler(logInfo,logError,serviceMessage)
	handlerMessageSubstance := initMessageSubstanceHandler(logInfo,logError,serviceMessageSubstance)
	handlerStoryMessageSubstance := initStoryMessageSubstanceHandler(logInfo,logError,serviceStoryMessageSubstance)
	handlerPostMessageSubstance := initPostMessageSubstanceHandler(logInfo,logError,servicePostMessageSubstance)
	handleFunc(handlerMessage, handlerMessageSubstance, handlerStoryMessageSubstance, handlerPostMessageSubstance)
}