package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func initMessageHandler(service *service.MessageService) *handler.MessageHandler{
	return &handler.MessageHandler{ Service: service }
}

func initMessageSubstanceHandler(service *service.MessageSubstanceService) *handler.MessageSubstanceHandler{
	return &handler.MessageSubstanceHandler{ Service: service }
}

func initStoryMessageSubstanceHandler(service *service.StoryMessageSubstanceService) *handler.StoryMessageSubstanceHandler{
	return &handler.StoryMessageSubstanceHandler{ Service: service }
}

func initPostMessageSubstanceHandler(service *service.PostMessageSubstanceService) *handler.PostMessageSubstanceHandler{
	return &handler.PostMessageSubstanceHandler{ Service: service }
}

func handleFunc(handlerMessage *handler.MessageHandler, handlerMessageSubstance *handler.MessageSubstanceHandler,
				handlerStoryMessageSubstance *handler.StoryMessageSubstanceHandler, handlerPostMessageSubstance *handler.PostMessageSubstanceHandler){

	mux := http.NewServeMux()

	mux.HandleFunc("/message/", handlerMessage.CreateMessage)
	mux.HandleFunc("/message_content/", handlerMessageSubstance.CreateMessageSubstance)
	mux.HandleFunc("/story_message_content/", handlerStoryMessageSubstance.CreateStoryMessageSubstance)
	mux.HandleFunc("/post_message_content/", handlerPostMessageSubstance.CreatePostMessageSubstance)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8089", handlerVar))
}

func main() {
	database := initDB()
	repoMessage := initMessageRepo(database)
	repoMessageSubstance := initMessageSubstanceRepo(database)
	repoStoryMessageSubstance := initStoryMessageSubstanceRepo(database)
	repoPostMessageSubstance := initPostMessageSubstanceRepo(database)
	serviceMessage := initMessageService(repoMessage)
	serviceMessageSubstance := initMessageSubstanceService(repoMessageSubstance)
	serviceStoryMessageSubstance := initStoryMessageSubstanceService(repoStoryMessageSubstance)
	servicePostMessageSubstance := initPostMessageSubstanceService(repoPostMessageSubstance)
	handlerMessage := initMessageHandler(serviceMessage)
	handlerMessageSubstance := initMessageSubstanceHandler(serviceMessageSubstance)
	handlerStoryMessageSubstance := initStoryMessageSubstanceHandler(serviceStoryMessageSubstance)
	handlerPostMessageSubstance := initPostMessageSubstanceHandler(servicePostMessageSubstance)
	handleFunc(handlerMessage, handlerMessageSubstance, handlerStoryMessageSubstance, handlerPostMessageSubstance)
}