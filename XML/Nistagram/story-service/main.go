package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn :=initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Story{},&model.StoryAlbum{},&model.SingleStory{}, &model.StoryHighlight{}, &model.SingleStoryStoryHighlights{})
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

func initStoryRepo(database *gorm.DB) *repository.StoryRepository{
	return &repository.StoryRepository { Database: database }
}

func initStoryAlbumRepo(database *gorm.DB) *repository.StoryAlbumRepository{
	return &repository.StoryAlbumRepository { Database: database }
}

func initSingleStoryRepo(database *gorm.DB) *repository.SingleStoryRepository{
	return &repository.SingleStoryRepository { Database: database }
}

func initStoryHighlightRepo(database *gorm.DB) *repository.StoryHighlightRepository{
	return &repository.StoryHighlightRepository { Database: database }
}

func initSingleStoryStoryHighlightsRepo(database *gorm.DB) *repository.SingleStoryStoryHighlightsRepository{
	return &repository.SingleStoryStoryHighlightsRepository { Database: database }
}

func initSingleStoryStoryHighlightsServices(repo *repository.SingleStoryStoryHighlightsRepository) *service.SingleStoryStoryHighlightsService{
	return &service.SingleStoryStoryHighlightsService { Repo: repo }
}

func initStoryHighlightServices(repo *repository.StoryHighlightRepository) *service.StoryHighlightService{
	return &service.StoryHighlightService { Repo: repo }
}

func initSingleStoryServices(repo *repository.SingleStoryRepository) *service.SingleStoryService{
	return &service.SingleStoryService { Repo: repo }
}

func initStoryAlbumServices(repo *repository.StoryAlbumRepository) *service.StoryAlbumService{
	return &service.StoryAlbumService { Repo: repo }
}

func initStoryServices(repo *repository.StoryRepository) *service.StoryService{
	return &service.StoryService { Repo: repo }
}

func initStoryHandler(service *service.StoryService) *handler.StoryHandler{
	return &handler.StoryHandler { Service: service }
}

func initStoryAlbumHandler(service *service.StoryAlbumService, storyService * service.StoryService) *handler.StoryAlbumHandler{
	return &handler.StoryAlbumHandler { Service: service, StoryService: storyService }
}

func initSingleStoryHandler(service *service.SingleStoryService, storyService *service.StoryService) *handler.SingleStoryHandler{
	return &handler.SingleStoryHandler { Service: service, StoryService: storyService}
}

func initStoryHighlightHandler(service *service.StoryHighlightService) *handler.StoryHighlightHandler{
	return &handler.StoryHighlightHandler { Service: service }
}

func initSingleStoryStoryHighlightsHandler(service *service.SingleStoryStoryHighlightsService) *handler.SingleStoryStoryHighlightsHandler{
	return &handler.SingleStoryStoryHighlightsHandler { Service: service }
}

func handleFunc(handlerStory *handler.StoryHandler, handlerStoryAlbum *handler.StoryAlbumHandler, handlerStoryHighlight *handler.StoryHighlightHandler,
	handlerSingleStoryStoryHighlights *handler.SingleStoryStoryHighlightsHandler,handlerSingleStory *handler.SingleStoryHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/story/", handlerStory.CreateStory).Methods("POST")
	router.HandleFunc("/story_album/", handlerStoryAlbum.CreateStoryAlbum).Methods("POST")
	router.HandleFunc("/story_highlight/", handlerStoryHighlight.CreateStoryHighlight).Methods("POST")
	router.HandleFunc("/single_story_story_highlights/", handlerSingleStoryStoryHighlights.CreateSingleStoryStoryHighlights).Methods("POST")
	router.HandleFunc("/single_story/", handlerSingleStory.CreateSingleStory).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	database := initDB()

	repoStory := initStoryRepo(database)
	serviceStory := initStoryServices(repoStory)
	handlerStory := initStoryHandler(serviceStory)

	repoStoryAlbum := initStoryAlbumRepo(database)
	serviceStoryAlbum := initStoryAlbumServices(repoStoryAlbum)
	handlerStoryAlbum := initStoryAlbumHandler(serviceStoryAlbum, serviceStory)

	repoSingleStory := initSingleStoryRepo(database)
	serviceSingleStory := initSingleStoryServices(repoSingleStory)
	handlerSingleStory := initSingleStoryHandler(serviceSingleStory, serviceStory)

	repoStoryHighlight := initStoryHighlightRepo(database)
	serviceStoryHighlight := initStoryHighlightServices(repoStoryHighlight)
	handlerStoryHighlight := initStoryHighlightHandler(serviceStoryHighlight)

	repoSingleStoryStoryHighlights := initSingleStoryStoryHighlightsRepo(database)
	serviceSingleStoryStoryHighlights := initSingleStoryStoryHighlightsServices(repoSingleStoryStoryHighlights)
	handlerSingleStoryStoryHighlights := initSingleStoryStoryHighlightsHandler(serviceSingleStoryStoryHighlights)

	handleFunc(handlerStory,handlerStoryAlbum,handlerStoryHighlight,handlerSingleStoryStoryHighlights,handlerSingleStory)
}