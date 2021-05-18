package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
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

	db.AutoMigrate(&model.Content{}, &model.AdvertisementContent{},&model.CommentContent{},&model.PostAlbumContent{},&model.SinglePostContent{},&model.SingleStoryContent{},&model.MessageContent{},&model.StoryAlbumContent{})
	return db
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

func initCommentContentRepo(database *gorm.DB) *repository.CommentContentRepository{
	return &repository.CommentContentRepository { Database: database }
}

func initCommentContentService(repo *repository.CommentContentRepository) *service.CommentContentService{
	return &service.CommentContentService { Repo: repo }
}

func initCommentContentHandler(service *service.CommentContentService) *handler.CommentContentHandler{
	return &handler.CommentContentHandler { Service: service }
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
	handlerCommentContent *handler.CommentContentHandler, handlerMessageContent *handler.MessageContentHandler){

	mux := http.NewServeMux()

	mux.HandleFunc("/content/", handlerContent.CreateContent)
	mux.HandleFunc("/single_post_content/", handlerSinglePostContent.CreateSinglePostContent)
	mux.HandleFunc("/single_story_content/", handlerSingleStoryContent.CreateSingleStoryContent)
	mux.HandleFunc("/advertisement_content/", handlerAdvertisementContent.CreateAdvertisementContent)
	mux.HandleFunc("/post_album_content/", handlerPostAlbumContent.CreatePostAlbumContent)
	mux.HandleFunc("/story_album_content/", handlerStoryAlbumContent.CreateStoryAlbumContent)
	mux.HandleFunc("/comment_content/", handlerCommentContent.CreateCommentContent)
	mux.HandleFunc("/message_content/", handlerMessageContent.CreateMessageContent)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8085", handlerVar))
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

	repoCommentContent := initCommentContentRepo(database)
	serviceCommentContent := initCommentContentService(repoCommentContent)
	handlerCommentContent := initCommentContentHandler(serviceCommentContent)

	repoMessageContent := initMessageContentRepo(database)
	serviceMessageContent := initMessageContentService(repoMessageContent)
	handlerMessageContent := initMessageContentHandler(serviceMessageContent)
	handleFunc(handlerContent, handlerAdvertisementContent,handlerPostAlbumContent,handlerSinglePostContent,handlerStoryAlbumContent,handlerSingleStoryContent,handlerCommentContent,handlerMessageContent)
}