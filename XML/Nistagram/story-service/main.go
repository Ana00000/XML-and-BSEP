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

func initStoryAlbumHandler(service *service.StoryAlbumService, storyService *service.StoryService,classicUserService * userService.ClassicUserService, classicUserFollowingsService * userService.ClassicUserFollowingsService, profileSettings *settingsService.ProfileSettingsService, storyAlbumContentService *contentService.StoryAlbumContentService,locationService *locationService.LocationService, storyAlbumTagStoryAlbumsService *tagsService.StoryAlbumTagStoryAlbumsService,tagService *tagsService.TagService, classicUserCloseFriendsService *userService.ClassicUserCloseFriendsService) *handler.StoryAlbumHandler{
	return &handler.StoryAlbumHandler{ Service: service, StoryService: storyService, ClassicUserService: classicUserService, ClassicUserFollowingsService: classicUserFollowingsService, ProfileSettings: profileSettings, StoryAlbumContentService: storyAlbumContentService, LocationService: locationService, StoryAlbumTagStoryAlbumsService: storyAlbumTagStoryAlbumsService, TagService: tagService, ClassicUserCloseFriendsService:  classicUserCloseFriendsService}
}

func initSingleStoryHandler(singleStoryService *service.SingleStoryService, storyService *service.StoryService) *handler.SingleStoryHandler{
	return &handler.SingleStoryHandler { SingleStoryService: singleStoryService, StoryService: storyService}
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


	router.HandleFunc("/story/",handlerStory.CreateStory).Methods("POST")
	router.HandleFunc("/story_album/", handlerStoryAlbum.CreateStoryAlbum).Methods("POST")
	router.HandleFunc("/single_story/", handlerSingleStory.CreateSingleStory).Methods("POST")
	router.HandleFunc("/story_highlight/", handlerStoryHighlight.CreateStoryHighlight).Methods("POST")
	router.HandleFunc("/single_story_story_highlights/",handlerSingleStoryStoryHighlights.CreateSingleStoryStoryHighlights).Methods("POST")

	router.HandleFunc("/find_single_story_for_id", handlerSingleStory.FindSingleStoryForId).Methods("GET")
	router.HandleFunc("/find_all_story_highlights_for_user", handlerStoryHighlight.FindAllStoryHighlightsForUser).Methods("GET")
	router.HandleFunc("/find_all_single_story_story_highlights_for_story", handlerSingleStoryStoryHighlights.FindAllSingleStoryStoryHighlightsForStory).Methods("GET")
	router.HandleFunc("/find_all_single_story_story_highlights_for_story_highlight", handlerSingleStoryStoryHighlights.FindAllSingleStoryStoryHighlightsForStoryHighlight).Methods("GET")

	// NOT REGISTERED USER
	router.HandleFunc("/find_all_stories_for_not_reg", handlerSingleStory.FindAllStoriesForUserNotRegisteredUser).Methods("GET") // kada neregistrovani user otvori PUBLIC usera sa spiska i onda na njegovom profilu vidi PUBLIC i NOT EXPIRED STORIJE
	router.HandleFunc("/find_all_public_stories_not_reg/", handlerSingleStory.FindAllPublicStoriesNotRegisteredUser).Methods("GET") // tab PUBLIC STORIES kada neregistroviani korisnik otvori sve PUBLIC, NOT EXPIRED I OD PUBLIC USERA

	router.HandleFunc("/find_all_album_stories_for_logged_user", handlerStoryAlbum.FindAllAlbumStoriesForLoggedUser).Methods("GET")
	router.HandleFunc("/find_selected_story_album_for_logged_user", handlerStoryAlbum.FindSelectedStoryAlbumByIdForLoggedUser).Methods("GET")
	router.HandleFunc("/find_all_public_album_stories_reg", handlerStoryAlbum.FindAllPublicAlbumStoriesRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_public_album_stories_not_reg/", handlerStoryAlbum.FindAllPublicAlbumStoriesNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_following_story_albums", handlerStoryAlbum.FindAllFollowingStoryAlbums).Methods("GET")

	// REGISTERED USER
	router.HandleFunc("/find_all_public_stories_reg", handlerSingleStory.FindAllPublicStoriesRegisteredUser).Methods("GET") // tab PUBLIC STORIES za reg usera - prikazuju se svi PUBLIC, NOT EXPIRED I OD PUBLIC USERA KOJI NISU ON!
	router.HandleFunc("/find_all_stories_for_reg", handlerSingleStory.FindAllStoriesForUserRegisteredUser).Methods("GET") // metoda koja se poziva kada registrovani user udje na profil nekog usera
	router.HandleFunc("/find_all_following_stories", handlerSingleStory.FindAllFollowingStories).Methods("GET") // tab FOLLOWING stories znaci svi storiji koji su PUBLIC I ALL FRIENDS , CLOSE FRIENDS storiji za one usere kojima je ulogovani user close friend

	router.HandleFunc("/find_selected_story_reg", handlerSingleStory.FindSelectedStoryByIdForRegisteredUsers).Methods("GET")//metoda koju ulogovani user poziva kada hoce da otvori svoj stori (kako bi ga dodao u HIGHLIGHTS)
	router.HandleFunc("/find_all_stories_for_logged_user", handlerSingleStory.FindAllStoriesForLoggedUser).Methods("GET") // metoda koju ulogovani user poziva kada hoce da vidi sve svoje storije (znaci i expired samo da nisu deleted)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))

}

func main() {
	database := initDB()
	repoStory := initStoryRepo(database)
	serviceStory := initStoryServices(repoStory)
	handlerStory := initStoryHandler(serviceStory)

	repoStoryAlbum := initStoryAlbumRepo(database)
	repoStoryAlbumContent := initStoryAlbumContentRepo(database)
	repoStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsRepo(database)
	repoClassicUserCloseFriends := initClassicUserCloseFriendsRepo(database)
	serviceStoryAlbum := initStoryAlbumServices(repoStoryAlbum)
	serviceStoryAlbumContent := initStoryAlbumContentService(repoStoryAlbumContent)
	serviceStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsService(repoStoryAlbumTagStoryAlbums)
	serviceClassicUserCloseFriends := initClassicUserCloseFriendsService(repoClassicUserCloseFriends)
	handlerStoryAlbum := initStoryAlbumHandler(serviceStoryAlbum, serviceStory, serviceClassicUser, serviceClassicUserFollowings, serviceProfileSettings, serviceStoryAlbumContent, serviceLocation, serviceStoryAlbumTagStoryAlbums, serviceTag, serviceClassicUserCloseFriends)

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