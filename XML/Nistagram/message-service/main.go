package main

import (
	"./handler"
	"./model"
	"./repository"
	"./service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func initDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=root dbname=nistagram-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.MessageSubstance{}, &model.Message{}, &model.StoryMessageSubstance{}, &model.PostMessageSubstance{})
	return db
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
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/message/", handlerMessage.CreateMessage).Methods("POST")
	router.HandleFunc("/message_content/", handlerMessageSubstance.CreateMessageSubstance).Methods("POST")
	router.HandleFunc("/story_message_content/", handlerStoryMessageSubstance.CreateStoryMessageSubstance).Methods("POST")
	router.HandleFunc("/post_message_content/", handlerPostMessageSubstance.CreatePostMessageSubstance).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
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