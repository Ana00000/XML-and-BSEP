package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	contentRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	contentService "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	locationRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	locationService "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	settingsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/story-service/service"
	tagsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	userRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
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

	db.AutoMigrate(&model.Story{},&model.StoryAlbum{},&model.SingleStory{}, &model.StoryHighlight{}, &model.SingleStoryStoryHighlights{})
	return db
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

func initSingleStoryHandler(singleStoryService *service.SingleStoryService, storyService *service.StoryService,classicUserService * userService.ClassicUserService, classicUserFollowingsService * userService.ClassicUserFollowingsService, profileSettings *settingsService.ProfileSettingsService, storyContentService *contentService.SingleStoryContentService,locationService *locationService.LocationService, storyTagStoriesService *tagsService.StoryTagStoriesService,tagService *tagsService.TagService) *handler.SingleStoryHandler{
	return &handler.SingleStoryHandler { SingleStoryService: singleStoryService, StoryService: storyService, ClassicUserService: classicUserService, ClassicUserFollowingsService: classicUserFollowingsService, ProfileSettings: profileSettings, StoryContentService: storyContentService, LocationService: locationService, StoryTagStoriesService: storyTagStoriesService, TagService: tagService }
}

func initStoryHighlightHandler(service *service.StoryHighlightService) *handler.StoryHighlightHandler{
	return &handler.StoryHighlightHandler { Service: service }
}

func initSingleStoryStoryHighlightsHandler(service *service.SingleStoryStoryHighlightsService) *handler.SingleStoryStoryHighlightsHandler{
	return &handler.SingleStoryStoryHighlightsHandler { Service: service }
}


// CLASSIC USER
func initClassicUserRepo(database *gorm.DB) *userRepository.ClassicUserRepository{
	return &userRepository.ClassicUserRepository{ Database: database }
}

func initClassicUserService(repo *userRepository.ClassicUserRepository) *userService.ClassicUserService{
	return &userService.ClassicUserService{ Repo: repo }
}

// CLASSIC USER FOLLOWINGS
func initClassicUserFollowingsRepo(database *gorm.DB) *userRepository.ClassicUserFollowingsRepository{
	return &userRepository.ClassicUserFollowingsRepository{ Database: database }
}

func initClassicUserFollowingsService(repo *userRepository.ClassicUserFollowingsRepository) *userService.ClassicUserFollowingsService{
	return &userService.ClassicUserFollowingsService{ Repo: repo }
}

// PROFILE SETTINGS
func initProfileSettingsRepo(database *gorm.DB) *settingsRepository.ProfileSettingsRepository{
	return &settingsRepository.ProfileSettingsRepository{ Database: database }
}

func initProfileSettingsService(repo *settingsRepository.ProfileSettingsRepository) *settingsService.ProfileSettingsService{
	return &settingsService.ProfileSettingsService{ Repo: repo }
}

// POST CONTENT
func initStoryContentRepo(database *gorm.DB) *contentRepository.SingleStoryContentRepository{
	return &contentRepository.SingleStoryContentRepository{ Database: database }
}

func initStoryContentService(repo *contentRepository.SingleStoryContentRepository) *contentService.SingleStoryContentService{
	return &contentService.SingleStoryContentService{ Repo: repo }
}

// LOCATION
func initLocationRepo(database *gorm.DB) *locationRepository.LocationRepository{
	return &locationRepository.LocationRepository{ Database: database }
}

func initLocationService(repo *locationRepository.LocationRepository) *locationService.LocationService{
	return &locationService.LocationService{ Repo: repo }
}

// POST TAG POST
func initStoryTagStoriesRepo(database *gorm.DB) *tagsRepository.StoryTagStoriesRepository{
	return &tagsRepository.StoryTagStoriesRepository{ Database: database }
}

func initStoryTagStoriesService(repo *tagsRepository.StoryTagStoriesRepository) *tagsService.StoryTagStoriesService{
	return &tagsService.StoryTagStoriesService{ Repo: repo }
}

// TAG
func initTagRepo(database *gorm.DB) *tagsRepository.TagRepository{
	return &tagsRepository.TagRepository{ Database: database }
}

func initTagService(repo *tagsRepository.TagRepository) *tagsService.TagService{
	return &tagsService.TagService{ Repo: repo }
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


	router.HandleFunc("/story/",handlerStory.CreateStory).Methods("POST")
	router.HandleFunc("/story_album/", handlerStoryAlbum.CreateStoryAlbum).Methods("POST")
	router.HandleFunc("/single_story/", handlerSingleStory.CreateSingleStory).Methods("POST")
	router.HandleFunc("/story_highlight/", handlerStoryHighlight.CreateStoryHighlight).Methods("POST")
	router.HandleFunc("/single_story_story_highlights/",handlerSingleStoryStoryHighlights.CreateSingleStoryStoryHighlights).Methods("POST")

	// NOT REGISTERED USER
	router.HandleFunc("/find_all_stories_for_not_reg", handlerSingleStory.FindAllStoriesForUserNotRegisteredUser).Methods("GET") // kada neregistrovani user otvori PUBLIC usera sa spiska i onda na njegovom profilu vidi PUBLIC i NOT EXPIRED STORIJE
	router.HandleFunc("/find_all_public_stories_not_reg/", handlerSingleStory.FindAllPublicStoriesNotRegisteredUser).Methods("GET") // tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA


	// REGISTERED USER
	router.HandleFunc("/find_all_public_stories_reg", handlerSingleStory.FindAllPublicStoriesRegisteredUser).Methods("GET") // tab PUBLIC STORIES za reg usera - prikazuju se svi PUBLIC, NOT EXPIRED I OD PUBLIC USERA KOJI NISU ON!
	router.HandleFunc("/find_all_stories_for_reg", handlerSingleStory.FindAllStoriesForUserRegisteredUser).Methods("GET")


	router.HandleFunc("/find_all_following_stories", handlerSingleStory.FindAllFollowingStories).Methods("GET")
	//router.HandleFunc("/find_selected_story_not_reg", handlerSingleStory.FindSelectedStoryByIdForNotRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_selected_story_reg", handlerSingleStory.FindSelectedStoryByIdForRegisteredUsers).Methods("GET")


	log.Fatal(http.ListenAndServe(":8086", cors(router)))

}

func main() {
	database := initDB()

	repoClassicUser := initClassicUserRepo(database)
	repoClassicUserFollowings := initClassicUserFollowingsRepo(database)
	repoProfileSettings := initProfileSettingsRepo(database)
	repoStoryContent := initStoryContentRepo(database)
	repoLocation := initLocationRepo(database)
	repoStoryTagStories := initStoryTagStoriesRepo(database)
	repoTag := initTagRepo(database)

	serviceClassicUser := initClassicUserService(repoClassicUser)
	serviceClassicUserFollowings := initClassicUserFollowingsService(repoClassicUserFollowings)
	serviceProfileSettings := initProfileSettingsService(repoProfileSettings)
	serviceStoryContent := initStoryContentService(repoStoryContent)
	serviceLocation := initLocationService(repoLocation)
	serviceStoryTagStories := initStoryTagStoriesService(repoStoryTagStories)
	serviceTag := initTagService(repoTag)

	repoStory := initStoryRepo(database)
	serviceStory := initStoryServices(repoStory)
	handlerStory := initStoryHandler(serviceStory)

	repoStoryAlbum := initStoryAlbumRepo(database)
	serviceStoryAlbum := initStoryAlbumServices(repoStoryAlbum)
	handlerStoryAlbum := initStoryAlbumHandler(serviceStoryAlbum, serviceStory)

	repoSingleStory := initSingleStoryRepo(database)
	serviceSingleStory := initSingleStoryServices(repoSingleStory)
	handlerSingleStory := initSingleStoryHandler(serviceSingleStory, serviceStory, serviceClassicUser, serviceClassicUserFollowings, serviceProfileSettings, serviceStoryContent, serviceLocation, serviceStoryTagStories, serviceTag)

	repoStoryHighlight := initStoryHighlightRepo(database)
	serviceStoryHighlight := initStoryHighlightServices(repoStoryHighlight)
	handlerStoryHighlight := initStoryHighlightHandler(serviceStoryHighlight)

	repoSingleStoryStoryHighlights := initSingleStoryStoryHighlightsRepo(database)
	serviceSingleStoryStoryHighlights := initSingleStoryStoryHighlightsServices(repoSingleStoryStoryHighlights)
	handlerSingleStoryStoryHighlights := initSingleStoryStoryHighlightsHandler(serviceSingleStoryStoryHighlights)

	handleFunc(handlerStory,handlerStoryAlbum,handlerStoryHighlight,handlerSingleStoryStoryHighlights,handlerSingleStory)
}