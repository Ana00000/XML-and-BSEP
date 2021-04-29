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

	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.MessageContent{})
	return db
}

func initMessageRepo(database *gorm.DB) *repository.MessageRepository{
	return &repository.MessageRepository{ Database: database }
}

func initMessageContentRepo(database *gorm.DB) *repository.MessageContentRepository{
	return &repository.MessageContentRepository{ Database: database }
}

func initMessageService(repo *repository.MessageRepository) *service.MessageService{
	return &service.MessageService{ Repo: repo }
}

func initMessageContentService(repo *repository.MessageContentRepository) *service.MessageContentService{
	return &service.MessageContentService{ Repo: repo }
}

func initMessageHandler(service *service.MessageService) *handler.MessageHandler{
	return &handler.MessageHandler{ Service: service }
}

func initMessageContentHandler(service *service.MessageContentService) *handler.MessageContentHandler{
	return &handler.MessageContentHandler{ Service: service }
}

func handleFunc(handlerMessage *handler.MessageHandler, handlerMessageContent *handler.MessageContentHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/message/", handlerMessage.CreateMessage).Methods("POST")
	router.HandleFunc("/message_content/", handlerMessageContent.CreateMessageContent).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repoMessage := initMessageRepo(database)
	repoMessageContent := initMessageContentRepo(database)
	serviceMessage := initMessageService(repoMessage)
	serviceMessageContent := initMessageContentService(repoMessageContent)
	handlerMessage := initMessageHandler(serviceMessage)
	handlerMessageContent := initMessageContentHandler(serviceMessageContent)
	handleFunc(handlerMessage, handlerMessageContent)
}