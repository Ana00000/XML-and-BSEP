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

	db.AutoMigrate(&model.MessageContent{}, &model.Message{}, &model.StoryMessageContent{}, &model.PostMessageContent{})
	return db
}

func initMessageRepo(database *gorm.DB) *repository.MessageRepository{
	return &repository.MessageRepository{ Database: database }
}

func initMessageContentRepo(database *gorm.DB) *repository.MessageContentRepository{
	return &repository.MessageContentRepository{ Database: database }
}

func initStoryMessageContentRepo(database *gorm.DB) *repository.StoryMessageContentRepository{
	return &repository.StoryMessageContentRepository{ Database: database }
}

func initPostMessageContentRepo(database *gorm.DB) *repository.PostMessageContentRepository{
	return &repository.PostMessageContentRepository{ Database: database }
}

func initMessageService(repo *repository.MessageRepository) *service.MessageService{
	return &service.MessageService{ Repo: repo }
}

func initMessageContentService(repo *repository.MessageContentRepository) *service.MessageContentService{
	return &service.MessageContentService{ Repo: repo }
}

func initStoryMessageContentService(repo *repository.StoryMessageContentRepository) *service.StoryMessageContentService{
	return &service.StoryMessageContentService{ Repo: repo }
}

func initPostMessageContentService(repo *repository.PostMessageContentRepository) *service.PostMessageContentService{
	return &service.PostMessageContentService{ Repo: repo }
}

func initMessageHandler(service *service.MessageService) *handler.MessageHandler{
	return &handler.MessageHandler{ Service: service }
}

func initMessageContentHandler(service *service.MessageContentService) *handler.MessageContentHandler{
	return &handler.MessageContentHandler{ Service: service }
}

func initStoryMessageContentHandler(service *service.StoryMessageContentService) *handler.StoryMessageContentHandler{
	return &handler.StoryMessageContentHandler{ Service: service }
}

func initPostMessageContentHandler(service *service.PostMessageContentService) *handler.PostMessageContentHandler{
	return &handler.PostMessageContentHandler{ Service: service }
}

func handleFunc(handlerMessage *handler.MessageHandler, handlerMessageContent *handler.MessageContentHandler,
				handlerStoryMessageContent *handler.StoryMessageContentHandler, handlerPostMessageContent *handler.PostMessageContentHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/message/", handlerMessage.CreateMessage).Methods("POST")
	router.HandleFunc("/message_content/", handlerMessageContent.CreateMessageContent).Methods("POST")
	router.HandleFunc("/story_message_content/", handlerStoryMessageContent.CreateStoryMessageContent).Methods("POST")
	router.HandleFunc("/post_message_content/", handlerPostMessageContent.CreatePostMessageContent).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repoMessage := initMessageRepo(database)
	repoMessageContent := initMessageContentRepo(database)
	repoStoryMessageContent := initStoryMessageContentRepo(database)
	repoPostMessageContent := initPostMessageContentRepo(database)
	serviceMessage := initMessageService(repoMessage)
	serviceMessageContent := initMessageContentService(repoMessageContent)
	serviceStoryMessageContent := initStoryMessageContentService(repoStoryMessageContent)
	servicePostMessageContent := initPostMessageContentService(repoPostMessageContent)
	handlerMessage := initMessageHandler(serviceMessage)
	handlerMessageContent := initMessageContentHandler(serviceMessageContent)
	handlerStoryMessageContent := initStoryMessageContentHandler(serviceStoryMessageContent)
	handlerPostMessageContent := initPostMessageContentHandler(servicePostMessageContent)
	handleFunc(handlerMessage, handlerMessageContent, handlerStoryMessageContent, handlerPostMessageContent)
}