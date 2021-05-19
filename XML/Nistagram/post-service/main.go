package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	settingsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	userRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/repository"
	userService "github.com/xml/XML-and-BSEP/XML/Nistagram/user-service/service"
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

	db.AutoMigrate(&model.Post{}, &model.Activity{}, &model.Comment{}, &model.PostAlbum{}, &model.SinglePost{}, &model.PostCollection{})
	return db
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

func initPostHandler(postService *service.PostService, classicUserService *userService.ClassicUserService, classicUserFollowingsService *userService.ClassicUserFollowingsService, profileSettings *settingsService.ProfileSettingsService) *handler.PostHandler{
	return &handler.PostHandler{ PostService: postService, ClassicUserService: classicUserService, ClassicUserFollowingsService: classicUserFollowingsService, ProfileSettings: profileSettings}
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






func handleFunc(handlerActivity *handler.ActivityHandler, handlerComment *handler.CommentHandler, handlerPost *handler.PostHandler,
	handlerPostAlbum *handler.PostAlbumHandler, handlerPostCollection *handler.PostCollectionHandler,
	handlerSinglePost *handler.SinglePostHandler, handlerPostCollectionPosts *handler.PostCollectionPostsHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/post/", handlerPost.CreatePost).Methods("POST")
	router.HandleFunc("/post_album/", handlerPostAlbum.CreatePostAlbum).Methods("POST")
	router.HandleFunc("/single_post/", handlerSinglePost.CreateSinglePost).Methods("POST")
	router.HandleFunc("/activity/", handlerActivity.CreateActivity).Methods("POST")
	router.HandleFunc("/comment/", handlerComment.CreateComment).Methods("POST")
	router.HandleFunc("/update_post/", handlerPost.UpdatePost).Methods("POST")
	router.HandleFunc("/post_collection/", handlerPostCollection.CreatePostCollection).Methods("POST")
	router.HandleFunc("/post_collection_posts/", handlerPostCollectionPosts.CreatePostCollectionPosts).Methods("POST")

	router.HandleFunc("/find_all_posts_for_not_reg", handlerPost.FindAllPostsForUserNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_posts_for_reg", handlerPost.FindAllPostsForUserRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_following_posts", handlerPost.FindAllFollowingPosts).Methods("GET")
	router.HandleFunc("/find_selected_post_not_reg", handlerPost.FindSelectedPostByIdForNotRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_selected_post_reg", handlerPost.FindSelectedPostByIdForRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_all_public_posts_not_reg/", handlerPost.FindAllPublicPostsNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_public_posts_reg", handlerPost.FindAllPublicPostsRegisteredUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8084", cors(router)))
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
	repoClassicUser := initClassicUserRepo(database)
	repoClassicUserFollowings := initClassicUserFollowingsRepo(database)
	repoProfileSettings := initProfileSettingsRepo(database)

	serviceActivity := initActivityService(repoActivity)
	serviceComment := initCommentService(repoComment)
	servicePost := initPostService(repoPost)
	servicePostAlbum := initPostAlbumService(repoPostAlbum)
	servicePostCollection := initPostCollectionService(repoPostCollection)
	serviceSinglePost := initSinglePostService(repoSinglePost)
	serviceClassicUser := initClassicUserService(repoClassicUser)
	serviceClassicUserFollowings := initClassicUserFollowingsService(repoClassicUserFollowings)
	serviceProfileSettings := initProfileSettingsService(repoProfileSettings)

	handlerActivity := initActivityHandler(serviceActivity)
	handlerComment := initCommentHandler(serviceComment)
	handlerPost := initPostHandler(servicePost, serviceClassicUser, serviceClassicUserFollowings, serviceProfileSettings)
	handlerPostAlbum := initPostAlbumHandler(servicePostAlbum, servicePost)
	handlerPostCollection := initPostCollectionHandler(servicePostCollection)
	handlerSinglePost := initSinglePostHandler(serviceSinglePost, servicePost)

	handleFunc(handlerActivity, handlerComment, handlerPost, handlerPostAlbum, handlerPostCollection, handlerSinglePost, handlerPostCollectionPosts)
}