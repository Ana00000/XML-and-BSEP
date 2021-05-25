package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
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

	db.AutoMigrate(&model.Content{}, &model.AdvertisementContent{},&model.PostAlbumContent{},&model.SinglePostContent{},&model.SingleStoryContent{},&model.MessageContent{},&model.StoryAlbumContent{})
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

func initAdvertisementContentRepo(database *gorm.DB) *repository.AdvertisementContentRepository{
	return &repository.AdvertisementContentRepository { Database: database }
}

func initAdvertisementContentService(repo *repository.AdvertisementContentRepository) *service.AdvertisementContentService{
	return &service.AdvertisementContentService { Repo: repo }
}

func initAdvertisementContentHandler(service *service.AdvertisementContentService) *handler.AdvertisementContentHandler{
	return &handler.AdvertisementContentHandler { Service: service }
}

func initContentRepo(database *gorm.DB) *repository.ContentRepository{
	return &repository.ContentRepository { Database: database }
}

func initContentService(repo *repository.ContentRepository) *service.ContentService{
	return &service.ContentService { Repo: repo }
}

func initContentHandler(service *service.ContentService) *handler.ContentHandler{
	return &handler.ContentHandler { Service: service }
}

func initPostAlbumContentRepo(database *gorm.DB) *repository.PostAlbumContentRepository{
	return &repository.PostAlbumContentRepository { Database: database }
}

func initPostAlbumContentService(repo *repository.PostAlbumContentRepository) *service.PostAlbumContentService{
	return &service.PostAlbumContentService { Repo: repo }
}

func initPostAlbumContentHandler(service *service.PostAlbumContentService, contentService *service.ContentService) *handler.PostAlbumContentHandler{
	return &handler.PostAlbumContentHandler { Service: service, ContentService: contentService  }
}

func initStoryAlbumContentRepo(database *gorm.DB) *repository.StoryAlbumContentRepository{
	return &repository.StoryAlbumContentRepository { Database: database }
}

func initStoryAlbumContentService(repo *repository.StoryAlbumContentRepository) *service.StoryAlbumContentService{
	return &service.StoryAlbumContentService { Repo: repo }
}

func initStoryAlbumContentHandler(service *service.StoryAlbumContentService, contentService * service.ContentService) *handler.StoryAlbumContentHandler{
	return &handler.StoryAlbumContentHandler { Service: service, ContentService: contentService}
}

func initSingleStoryContentRepo(database *gorm.DB) *repository.SingleStoryContentRepository{
	return &repository.SingleStoryContentRepository { Database: database }
}

func initSingleStoryContentService(repo *repository.SingleStoryContentRepository) *service.SingleStoryContentService{
	return &service.SingleStoryContentService { Repo: repo }
}

func initSingleStoryContentHandler(service *service.SingleStoryContentService, contentService *service.ContentService) *handler.SingleStoryContentHandler{
	return &handler.SingleStoryContentHandler { Service: service, ContentService: contentService }
}

func initSinglePostContentRepo(database *gorm.DB) *repository.SinglePostContentRepository{
	return &repository.SinglePostContentRepository { Database: database }
}

func initSinglePostContentService(repo *repository.SinglePostContentRepository) *service.SinglePostContentService{
	return &service.SinglePostContentService { Repo: repo }
}

func initSinglePostContentHandler(service *service.SinglePostContentService, contentService *service.ContentService) *handler.SinglePostContentHandler{
	return &handler.SinglePostContentHandler { Service: service, ContentService: contentService }
}

func initMessageContentRepo(database *gorm.DB) *repository.MessageContentRepository{
	return &repository.MessageContentRepository { Database: database }
}

func initMessageContentService(repo *repository.MessageContentRepository) *service.MessageContentService{
	return &service.MessageContentService { Repo: repo }
}

func initMessageContentHandler(service *service.MessageContentService) *handler.MessageContentHandler{
	return &handler.MessageContentHandler { Service: service }
}

func handleFunc(handlerContent *handler.ContentHandler, handlerAdvertisementContent *handler.AdvertisementContentHandler,
	handlerPostAlbumContent *handler.PostAlbumContentHandler, handlerSinglePostContent *handler.SinglePostContentHandler,
	handlerStoryAlbumContent *handler.StoryAlbumContentHandler, handlerSingleStoryContent *handler.SingleStoryContentHandler,
	handlerMessageContent *handler.MessageContentHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/content/", handlerContent.CreateContent).Methods("POST")
	router.HandleFunc("/single_post_content/", handlerSinglePostContent.CreateSinglePostContent).Methods("POST")
	router.HandleFunc("/single_story_content/", handlerSingleStoryContent.CreateSingleStoryContent).Methods("POST")
	router.HandleFunc("/advertisement_content/", handlerAdvertisementContent.CreateAdvertisementContent).Methods("POST")
	router.HandleFunc("/post_album_content/", handlerPostAlbumContent.CreatePostAlbumContent).Methods("POST")
	router.HandleFunc("/story_album_content/", handlerStoryAlbumContent.CreateStoryAlbumContent).Methods("POST")
	//router.HandleFunc("/comment_content/", handlerCommentContent.CreateCommentContent).Methods("POST")
	router.HandleFunc("/message_content/", handlerMessageContent.CreateMessageContent).Methods("POST")
	router.HandleFunc("/uploadPostMedia/", handlerSinglePostContent.Upload).Methods("POST")
	router.HandleFunc("/uploadPostAlbumMedia/", handlerPostAlbumContent.Upload).Methods("POST")
	router.HandleFunc("/uploadStoryMedia/", handlerSingleStoryContent.Upload).Methods("POST")
	router.HandleFunc("/uploadStoryAlbumMedia/", handlerStoryAlbumContent.Upload).Methods("POST")
	router.HandleFunc("/find_all_contents_for_stories/", handlerSingleStoryContent.FindAllContentsForStories).Methods("POST")
	router.HandleFunc("/find_all_contents_for_story/", handlerSingleStoryContent.FindAllContentsForStory).Methods("POST")
	router.HandleFunc("/find_all_contents_for_story_album/", handlerStoryAlbumContent.FindAllContentsForStoryAlbum).Methods("POST")
	router.HandleFunc("/find_all_contents_for_story_albums/", handlerStoryAlbumContent.FindAllContentsForStoryAlbums).Methods("POST")
	router.HandleFunc("/find_all_contents_for_posts/", handlerSinglePostContent.FindAllContentsForPosts).Methods("POST")
	router.HandleFunc("/find_all_contents_for_post/", handlerSinglePostContent.FindAllContentsForPost).Methods("POST")
	router.HandleFunc("/find_all_contents_for_post_album/", handlerPostAlbumContent.FindAllContentsForPostAlbum).Methods("POST")
	router.HandleFunc("/find_all_contents_for_post_albums/", handlerPostAlbumContent.FindAllContentsForPostAlbums).Methods("POST")

	router.HandleFunc("/find_single_story_content_for_story_id", handlerSingleStoryContent.FindSingleStoryContentForStoryId).Methods("GET")


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {

	database := initDB()
	repoContent := initContentRepo(database)
	serviceContent := initContentService(repoContent)
	handlerContent := initContentHandler(serviceContent)

	repoAdvertisementContent := initAdvertisementContentRepo(database)
	serviceAdvertisementContent := initAdvertisementContentService(repoAdvertisementContent)
	handlerAdvertisementContent := initAdvertisementContentHandler(serviceAdvertisementContent)

	repoPostAlbumContent := initPostAlbumContentRepo(database)
	servicePostAlbumContent := initPostAlbumContentService(repoPostAlbumContent)
	handlerPostAlbumContent := initPostAlbumContentHandler(servicePostAlbumContent, serviceContent)

	repoSinglePostContent := initSinglePostContentRepo(database)
	serviceSinglePostContent := initSinglePostContentService(repoSinglePostContent)
	handlerSinglePostContent := initSinglePostContentHandler(serviceSinglePostContent, serviceContent)

	repoStoryAlbumContent := initStoryAlbumContentRepo(database)
	serviceStoryAlbumContent := initStoryAlbumContentService(repoStoryAlbumContent)
	handlerStoryAlbumContent := initStoryAlbumContentHandler(serviceStoryAlbumContent, serviceContent)

	repoSingleStoryContent := initSingleStoryContentRepo(database)
	serviceSingleStoryContent := initSingleStoryContentService(repoSingleStoryContent)
	handlerSingleStoryContent := initSingleStoryContentHandler(serviceSingleStoryContent, serviceContent)

	repoMessageContent := initMessageContentRepo(database)
	serviceMessageContent := initMessageContentService(repoMessageContent)
	handlerMessageContent := initMessageContentHandler(serviceMessageContent)

	handleFunc(handlerContent, handlerAdvertisementContent,handlerPostAlbumContent,handlerSinglePostContent,handlerStoryAlbumContent,handlerSingleStoryContent,handlerMessageContent)
}