package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func initDB() *gorm.DB{
	dsn := initDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Post{}, &model.Activity{}, &model.Comment{}, &model.PostAlbum{}, &model.SinglePost{}, &model.PostCollection{})
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

func initActivityRepo(database *gorm.DB) *repository.ActivityRepository{
	return &repository.ActivityRepository{ Database: database }
}

func initCommentRepo(database *gorm.DB) *repository.CommentRepository{
	return &repository.CommentRepository{ Database: database }
}

func initPostRepo(database *gorm.DB) *repository.PostRepository{
	return &repository.PostRepository{ Database: database }
}

func initPostAlbumRepo(database *gorm.DB) *repository.PostAlbumRepository{
	return &repository.PostAlbumRepository{ Database: database }
}

func initPostCollectionRepo(database *gorm.DB) *repository.PostCollectionRepository{
	return &repository.PostCollectionRepository{ Database: database }
}

func initSinglePostRepo(database *gorm.DB) *repository.SinglePostRepository{
	return &repository.SinglePostRepository{ Database: database }
}

func initActivityService(repo *repository.ActivityRepository) *service.ActivityService{
	return &service.ActivityService{ Repo: repo }
}

func initCommentService(repo *repository.CommentRepository) *service.CommentService{
	return &service.CommentService{ Repo: repo }
}

func initPostService(repo *repository.PostRepository) *service.PostService{
	return &service.PostService{ Repo: repo }
}

func initPostAlbumService(repo *repository.PostAlbumRepository) *service.PostAlbumService{
	return &service.PostAlbumService{ Repo: repo }
}

func initPostCollectionService(repo *repository.PostCollectionRepository) *service.PostCollectionService{
	return &service.PostCollectionService{ Repo: repo }
}

func initSinglePostService(repo *repository.SinglePostRepository) *service.SinglePostService{
	return &service.SinglePostService{ Repo: repo }
}

func initActivityHandler(service *service.ActivityService) *handler.ActivityHandler{
	return &handler.ActivityHandler{ Service: service }
}

func initCommentHandler(service *service.CommentService) *handler.CommentHandler{
	return &handler.CommentHandler{ Service: service }
}

func initPostHandler(service *service.PostService) *handler.PostHandler{
	return &handler.PostHandler{ Service: service }
}

func initPostAlbumHandler(service *service.PostAlbumService, postService *service.PostService) *handler.PostAlbumHandler{
	return &handler.PostAlbumHandler{ Service: service, PostService: postService}
}

func initPostCollectionHandler(service *service.PostCollectionService) *handler.PostCollectionHandler{
	return &handler.PostCollectionHandler{ Service: service }
}

func initSinglePostHandler(service *service.SinglePostService, postService *service.PostService) *handler.SinglePostHandler{
	return &handler.SinglePostHandler{ Service: service, PostService: postService }
}

func initPostCollectionPostsRepo(database *gorm.DB) *repository.PostCollectionPostsRepository{
	return &repository.PostCollectionPostsRepository { Database: database }
}

func initPostCollectionPostsServices(repo *repository.PostCollectionPostsRepository) *service.PostCollectionPostsService{
	return &service.PostCollectionPostsService { Repo: repo }
}

func initPostCollectionPostsHandler(service *service.PostCollectionPostsService) *handler.PostCollectionPostsHandler{
	return &handler.PostCollectionPostsHandler { Service: service }
}

func handleFunc(handlerActivity *handler.ActivityHandler, handlerComment *handler.CommentHandler, handlerPost *handler.PostHandler,
	handlerPostAlbum *handler.PostAlbumHandler, handlerPostCollection *handler.PostCollectionHandler,
	handlerSinglePost *handler.SinglePostHandler, handlerPostCollectionPosts *handler.PostCollectionPostsHandler){

	mux := http.NewServeMux()

	mux.HandleFunc("/post/", handlerPost.CreatePost)
	mux.HandleFunc("/post_album/", handlerPostAlbum.CreatePostAlbum)
	mux.HandleFunc("/single_post/", handlerSinglePost.CreateSinglePost)
	mux.HandleFunc("/activity/", handlerActivity.CreateActivity)
	mux.HandleFunc("/comment/", handlerComment.CreateComment)
	mux.HandleFunc("/update_post/", handlerPost.UpdatePost)
	mux.HandleFunc("/post_collection/", handlerPostCollection.CreatePostCollection)
	mux.HandleFunc("/post_collection_posts/", handlerPostCollectionPosts.CreatePostCollectionPosts)

	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8084", handlerVar))
}

func main() {
	database := initDB()
	repoPostCollectionPosts := initPostCollectionPostsRepo(database)
	servicePostCollectionPosts := initPostCollectionPostsServices(repoPostCollectionPosts)
	handlerPostCollectionPosts := initPostCollectionPostsHandler(servicePostCollectionPosts)

	repoActivity := initActivityRepo(database)
	repoComment := initCommentRepo(database)
	repoPost := initPostRepo(database)
	repoPostAlbum := initPostAlbumRepo(database)
	repoPostCollection := initPostCollectionRepo(database)
	repoSinglePost := initSinglePostRepo(database)

	serviceActivity := initActivityService(repoActivity)
	serviceComment := initCommentService(repoComment)
	servicePost := initPostService(repoPost)
	servicePostAlbum := initPostAlbumService(repoPostAlbum)
	servicePostCollection := initPostCollectionService(repoPostCollection)
	serviceSinglePost := initSinglePostService(repoSinglePost)

	handlerActivity := initActivityHandler(serviceActivity)
	handlerComment := initCommentHandler(serviceComment)
	handlerPost := initPostHandler(servicePost)
	handlerPostAlbum := initPostAlbumHandler(servicePostAlbum, servicePost)
	handlerPostCollection := initPostCollectionHandler(servicePostCollection)
	handlerSinglePost := initSinglePostHandler(serviceSinglePost, servicePost)

	handleFunc(handlerActivity, handlerComment, handlerPost, handlerPostAlbum, handlerPostCollection, handlerSinglePost, handlerPostCollectionPosts)
}