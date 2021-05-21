package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
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

	db.AutoMigrate(&model.Tag{}, &model.StoryTagStories{},
				   &model.CommentTagComments{}, &model.PostTagPosts{},
				   &model.PostAlbumTagPostAlbums{}, &model.StoryAlbumTagStoryAlbums{})
	return db
}



func initPostTagPostsRepo(database *gorm.DB) *repository.PostTagPostsRepository{
	return &repository.PostTagPostsRepository { Database: database }
}

func initStoryTagStoriesRepo(database *gorm.DB) *repository.StoryTagStoriesRepository{
	return &repository.StoryTagStoriesRepository { Database: database }
}

func initCommentTagCommentsRepo(database *gorm.DB) *repository.CommentTagCommentsRepository{
	return &repository.CommentTagCommentsRepository { Database: database }
}


func initPostAlbumTagPostAlbumsRepo(database *gorm.DB) *repository.PostAlbumTagPostAlbumsRepository{
	return &repository.PostAlbumTagPostAlbumsRepository { Database: database }
}


func initStoryAlbumTagStoryAlbumsRepo(database *gorm.DB) *repository.StoryAlbumTagStoryAlbumsRepository{
	return &repository.StoryAlbumTagStoryAlbumsRepository { Database: database }
}

func initTagRepo(database *gorm.DB) *repository.TagRepository{
	return &repository.TagRepository { Database: database }
}

func initTagServices(repo *repository.TagRepository) *service.TagService{
	return &service.TagService { Repo: repo }
}

func initPostTagPostsServices(repo *repository.PostTagPostsRepository) *service.PostTagPostsService{
	return &service.PostTagPostsService { Repo: repo }
}

func initCommentTagCommentsServices(repo *repository.CommentTagCommentsRepository) *service.CommentTagCommentsService{
	return &service.CommentTagCommentsService { Repo: repo }
}

func initStoryTagStoriesServices(repo *repository.StoryTagStoriesRepository) *service.StoryTagStoriesService{
	return &service.StoryTagStoriesService { Repo: repo }
}


func initPostAlbumTagPostAlbumsServices(repo *repository.PostAlbumTagPostAlbumsRepository) *service.PostAlbumTagPostAlbumsService{
	return &service.PostAlbumTagPostAlbumsService { Repo: repo }
}



func initStoryAlbumTagStoryAlbumsServices(repo *repository.StoryAlbumTagStoryAlbumsRepository) *service.StoryAlbumTagStoryAlbumsService{
	return &service.StoryAlbumTagStoryAlbumsService { Repo: repo }
}

func initTagHandler(service *service.TagService) *handler.TagHandler{
	return &handler.TagHandler { Service: service }
}


func initStoryTagStoriesHandler(service *service.StoryTagStoriesService) *handler.StoryTagStoriesHandler{
	return &handler.StoryTagStoriesHandler { Service: service }
}

func initCommentTagCommentsHandler(service *service.CommentTagCommentsService) *handler.CommentTagCommentsHandler{
	return &handler.CommentTagCommentsHandler { Service: service }
}

func initPostTagPostsHandler(service *service.PostTagPostsService) *handler.PostTagPostsHandler{
	return &handler.PostTagPostsHandler { Service: service }
}


func initPostAlbumTagPostAlbumsHandler(service *service.PostAlbumTagPostAlbumsService) *handler.PostAlbumTagPostAlbumsHandler{
	return &handler.PostAlbumTagPostAlbumsHandler { Service: service }
}



func initStoryAlbumTagStoryAlbumsHandler(service *service.StoryAlbumTagStoryAlbumsService) *handler.StoryAlbumTagStoryAlbumsHandler{
	return &handler.StoryAlbumTagStoryAlbumsHandler { Service: service }
}

func handleFunc(handlerTag *handler.TagHandler, handlerCommentTagComments *handler.CommentTagCommentsHandler,
	handlerPostTagPosts *handler.PostTagPostsHandler, handlerStoryTagStories *handler.StoryTagStoriesHandler,  handlerPostAlbumTagPostAlbums *handler.PostAlbumTagPostAlbumsHandler, handlerStoryAlbumTagStoryAlbums *handler.StoryAlbumTagStoryAlbumsHandler){


	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/tag/", handlerTag.CreateTag).Methods("POST")
	router.HandleFunc("/comment_tag_comments/", handlerCommentTagComments.CreateCommentTagComments).Methods("POST")
	router.HandleFunc("/post_tag_posts/", handlerPostTagPosts.CreatePostTagPosts).Methods("POST")
	router.HandleFunc("/story_tag_stories/", handlerStoryTagStories.CreateStoryTagStories).Methods("POST")
	router.HandleFunc("/post_album_tag_post_albums/", handlerPostAlbumTagPostAlbums.CreatePostAlbumTagPostAlbums).Methods("POST")
	router.HandleFunc("/story_album_tag_story_albums/", handlerStoryAlbumTagStoryAlbums.CreateStoryAlbumTagStoryAlbums).Methods("POST")


	log.Fatal(http.ListenAndServe(":8082", cors(router)))

}

func main() {
	database := initDB()

	repoTag := initTagRepo(database)
	serviceTag := initTagServices(repoTag)
	handlerTag := initTagHandler(serviceTag)

	repoPostTagPosts := initPostTagPostsRepo(database)
	servicePostTagPosts := initPostTagPostsServices(repoPostTagPosts)
	handlerPostTagPosts := initPostTagPostsHandler(servicePostTagPosts)

	repoStoryTagStories := initStoryTagStoriesRepo(database)
	serviceStoryTagStories := initStoryTagStoriesServices(repoStoryTagStories)
	handlerStoryTagStories := initStoryTagStoriesHandler(serviceStoryTagStories)

	repoCommentTagComments := initCommentTagCommentsRepo(database)
	serviceCommentTagComments := initCommentTagCommentsServices(repoCommentTagComments)
	handlerCommentTagComments := initCommentTagCommentsHandler(serviceCommentTagComments)


	repoPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsRepo(database)
	servicePostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsServices(repoPostAlbumTagPostAlbums)
	handlerPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsHandler(servicePostAlbumTagPostAlbums)


	repoStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsRepo(database)
	serviceStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsServices(repoStoryAlbumTagStoryAlbums)
	handlerStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsHandler(serviceStoryAlbumTagStoryAlbums)

	handleFunc(handlerTag, handlerCommentTagComments,handlerPostTagPosts, handlerStoryTagStories, handlerPostAlbumTagPostAlbums, handlerStoryAlbumTagStoryAlbums)
}