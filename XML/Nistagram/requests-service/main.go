package main

import (
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/requests-service/service"
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	db.AutoMigrate(&model.InappropriateContentRequest{}, &model.PostICR{}, &model.StoryICR{},
	               &model.CommentICR{}, &model.VerificationRequest{}, &model.AgentRegistrationRequest{})
	return db
}

func initInappropriateContentRequestRepo(database *gorm.DB) *repository.InappropriateContentRequestRepository{
	return &repository.InappropriateContentRequestRepository { Database: database }
}

func initPostICRRepo(database *gorm.DB) *repository.PostICRRepository{
	return &repository.PostICRRepository { Database: database }
}

func initStoryICRRepo(database *gorm.DB) *repository.StoryICRRepository{
	return &repository.StoryICRRepository { Database: database }
}

func initCommentICRRepo(database *gorm.DB) *repository.CommentICRRepository{
	return &repository.CommentICRRepository { Database: database }
}

func initVerificationRequestRepo(database *gorm.DB) *repository.VerificationRequestRepository{
	return &repository.VerificationRequestRepository { Database: database }
}

func initAgentRegistrationRequestRepo(database *gorm.DB) *repository.AgentRegistrationRequestRepository{
	return &repository.AgentRegistrationRequestRepository { Database: database }
}

func initInappropriateContentRequestServices(repo *repository.InappropriateContentRequestRepository) *service.InappropriateContentRequestService{
	return &service.InappropriateContentRequestService { Repo: repo }
}

func initPostICRServices(repo *repository.PostICRRepository) *service.PostICRService{
	return &service.PostICRService { Repo: repo }
}

func initStoryICRServices(repo *repository.StoryICRRepository) *service.StoryICRService{
	return &service.StoryICRService { Repo: repo }
}

func initCommentICRServices(repo *repository.CommentICRRepository) *service.CommentICRService{
	return &service.CommentICRService { Repo: repo }
}

func initVerificationRequestServices(repo *repository.VerificationRequestRepository) *service.VerificationRequestService{
	return &service.VerificationRequestService { Repo: repo }
}

func initAgentRegistrationRequestServices(repo *repository.AgentRegistrationRequestRepository) *service.AgentRegistrationRequestService{
	return &service.AgentRegistrationRequestService { Repo: repo }
}

func initInappropriateContentRequestHandler(service *service.InappropriateContentRequestService) *handler.InappropriateContentRequestHandler{
	return &handler.InappropriateContentRequestHandler { Service: service }
}

func initPostICRHandler(service *service.PostICRService) *handler.PostICRHandler{
	return &handler.PostICRHandler { Service: service }
}

func initStoryICRHandler(service *service.StoryICRService) *handler.StoryICRHandler{
	return &handler.StoryICRHandler { Service: service }
}

func initCommentICRHandler(service *service.CommentICRService) *handler.CommentICRHandler{
	return &handler.CommentICRHandler { Service: service }
}

func initVerificationRequestHandler(service *service.VerificationRequestService) *handler.VerificationRequestHandler{
	return &handler.VerificationRequestHandler { Service: service }
}

func initAgentRegistrationRequestHandler(service *service.AgentRegistrationRequestService) *handler.AgentRegistrationRequestHandler{
	return &handler.AgentRegistrationRequestHandler { Service: service }
}

func handleFunc(inappropriateContentRequestHandler *handler.InappropriateContentRequestHandler, postICRHandler *handler.PostICRHandler,
	storyICRHandler *handler.StoryICRHandler, commentICRHandler *handler.CommentICRHandler, verificationRequestHandler *handler.VerificationRequestHandler,
	agentRegistrationRequestHandler *handler.AgentRegistrationRequestHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/inappropriateContentRequest", inappropriateContentRequestHandler.CreateInappropriateContentRequest).Methods("POST")
	router.HandleFunc("/postICR", postICRHandler.CreatePostICR).Methods("POST")
	router.HandleFunc("/storyICR", storyICRHandler.CreateStoryICR).Methods("POST")
	router.HandleFunc("/commentICR", commentICRHandler.CreateCommentICR).Methods("POST")
	router.HandleFunc("/verificationRequest", verificationRequestHandler.CreateVerificationRequest).Methods("POST")
	router.HandleFunc("/agentRegistrationRequestHandler", agentRegistrationRequestHandler.CreateAgentRegistrationRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()

	inappropriateContentRequestRepo := initInappropriateContentRequestRepo(database)
	postICRRepo := initPostICRRepo(database)
	storyICRRepo := initStoryICRRepo(database)
	commentICRRepo := initCommentICRRepo(database)
	verificationRequestRepo := initVerificationRequestRepo(database)
	agentRegistrationRequestRepo := initAgentRegistrationRequestRepo(database)

	inappropriateContentRequestService := initInappropriateContentRequestServices(inappropriateContentRequestRepo)
	postICRService := initPostICRServices(postICRRepo)
	storyICRService := initStoryICRServices(storyICRRepo)
	commentICRService := initCommentICRServices(commentICRRepo)
	verificationRequestService := initVerificationRequestServices(verificationRequestRepo)
	agentRegistrationRequestService := initAgentRegistrationRequestServices(agentRegistrationRequestRepo)

	inappropriateContentRequestHandler := initInappropriateContentRequestHandler(inappropriateContentRequestService)
	postICRHandler := initPostICRHandler(postICRService)
	storyICRHandler := initStoryICRHandler(storyICRService)
	commentICRHandler := initCommentICRHandler(commentICRService)
	verificationRequestRHandler := initVerificationRequestHandler(verificationRequestService)
	agentRegistrationRequestHandler := initAgentRegistrationRequestHandler(agentRegistrationRequestService)

	handleFunc(inappropriateContentRequestHandler, postICRHandler,storyICRHandler, commentICRHandler,
		verificationRequestRHandler,agentRegistrationRequestHandler)
}