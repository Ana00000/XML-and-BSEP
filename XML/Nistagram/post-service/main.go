package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	contentRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/repository"
	contentService "github.com/xml/XML-and-BSEP/XML/Nistagram/content-service/service"
	locationRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	locationService "github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	settingsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/repository"
	settingsService "github.com/xml/XML-and-BSEP/XML/Nistagram/settings-service/service"
	tagsRepository "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	tagsService "github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
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

func initPostHandler(postService *service.PostService) *handler.PostHandler{
	return &handler.PostHandler{ PostService: postService}
}

func initPostAlbumHandler(service *service.PostAlbumService, postService *service.PostService) *handler.PostAlbumHandler{
	return &handler.PostAlbumHandler{ Service: service, PostService: postService}
}

func initPostCollectionHandler(service *service.PostCollectionService) *handler.PostCollectionHandler{
	return &handler.PostCollectionHandler{ Service: service }
}

func initSinglePostHandler(singlePostService *service.SinglePostService, postService *service.PostService,classicUserService * userService.ClassicUserService, classicUserFollowingsService * userService.ClassicUserFollowingsService, profileSettings *settingsService.ProfileSettingsService, postContentService *contentService.SinglePostContentService,locationService *locationService.LocationService, postTagPostsService *tagsService.PostTagPostsService,tagService *tagsService.TagService) *handler.SinglePostHandler{
	return &handler.SinglePostHandler{ SinglePostService: singlePostService, PostService: postService, ClassicUserService: classicUserService, ClassicUserFollowingsService: classicUserFollowingsService, ProfileSettings: profileSettings, PostContentService: postContentService, LocationService: locationService, PostTagPostsService: postTagPostsService, TagService: tagService }
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

// POST CONTENT
func initPostContentRepo(database *gorm.DB) *contentRepository.SinglePostContentRepository{
	return &contentRepository.SinglePostContentRepository{ Database: database }
}

func initPostContentService(repo *contentRepository.SinglePostContentRepository) *contentService.SinglePostContentService{
	return &contentService.SinglePostContentService{ Repo: repo }
}

// LOCATION
func initLocationRepo(database *gorm.DB) *locationRepository.LocationRepository{
	return &locationRepository.LocationRepository{ Database: database }
}

func initLocationService(repo *locationRepository.LocationRepository) *locationService.LocationService{
	return &locationService.LocationService{ Repo: repo }
}

// POST TAG POST
func initPostTagPostRepo(database *gorm.DB) *tagsRepository.PostTagPostsRepository{
	return &tagsRepository.PostTagPostsRepository{ Database: database }
}

func initPostTagPostService(repo *tagsRepository.PostTagPostsRepository) *tagsService.PostTagPostsService{
	return &tagsService.PostTagPostsService{ Repo: repo }
}

// TAG
func initTagRepo(database *gorm.DB) *tagsRepository.TagRepository{
	return &tagsRepository.TagRepository{ Database: database }
}

func initTagService(repo *tagsRepository.TagRepository) *tagsService.TagService{
	return &tagsService.TagService{ Repo: repo }
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

	router.HandleFunc("/find_all_likes_for_post", handlerActivity.FindAllLikesForPost).Methods("GET")
	router.HandleFunc("/find_all_dislikes_for_post", handlerActivity.FindAllDislikesForPost).Methods("GET")
	router.HandleFunc("/find_all_favorites_for_post", handlerActivity.FindAllFavoritesForPost).Methods("GET")
	router.HandleFunc("/find_all_activities_for_post", handlerActivity.FindAllActivitiesForPost).Methods("GET")

	router.HandleFunc("/comment/", handlerComment.CreateComment).Methods("POST")
	router.HandleFunc("/update_post/", handlerPost.UpdatePost).Methods("POST")
	router.HandleFunc("/post_collection/", handlerPostCollection.CreatePostCollection).Methods("POST")
	router.HandleFunc("/post_collection_posts/", handlerPostCollectionPosts.CreatePostCollectionPosts).Methods("POST")

	router.HandleFunc("/find_all_posts_for_not_reg", handlerSinglePost.FindAllPostsForUserNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_posts_for_reg", handlerSinglePost.FindAllPostsForUserRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_following_posts", handlerSinglePost.FindAllFollowingPosts).Methods("GET")
	router.HandleFunc("/find_selected_post_not_reg", handlerSinglePost.FindSelectedPostByIdForNotRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_selected_post_reg", handlerSinglePost.FindSelectedPostByIdForRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_all_public_posts_not_reg/", handlerSinglePost.FindAllPublicPostsNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_public_posts_reg", handlerSinglePost.FindAllPublicPostsRegisteredUser).Methods("GET")

	router.HandleFunc("/find_all_posts_for_logged_user", handlerSinglePost.FindAllPostsForLoggedUser).Methods("GET")
	router.HandleFunc("/find_selected_post_for_logged_user", handlerSinglePost.FindSelectedPostByIdForLoggedUser).Methods("GET")


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
	repoPostContent := initPostContentRepo(database)
	repoLocation := initLocationRepo(database)
	repoPostTagPost := initPostTagPostRepo(database)
	repoTag := initTagRepo(database)

	serviceActivity := initActivityService(repoActivity)
	serviceComment := initCommentService(repoComment)
	servicePost := initPostService(repoPost)
	servicePostAlbum := initPostAlbumService(repoPostAlbum)
	servicePostCollection := initPostCollectionService(repoPostCollection)
	serviceSinglePost := initSinglePostService(repoSinglePost)
	serviceClassicUser := initClassicUserService(repoClassicUser)
	serviceClassicUserFollowings := initClassicUserFollowingsService(repoClassicUserFollowings)
	serviceProfileSettings := initProfileSettingsService(repoProfileSettings)
	servicePostContent := initPostContentService(repoPostContent)
	serviceLocation := initLocationService(repoLocation)
	servicePostTagPost := initPostTagPostService(repoPostTagPost)
	serviceTag := initTagService(repoTag)

	handlerActivity := initActivityHandler(serviceActivity)
	handlerComment := initCommentHandler(serviceComment)
	handlerPost := initPostHandler(servicePost)
	handlerPostAlbum := initPostAlbumHandler(servicePostAlbum, servicePost)
	handlerPostCollection := initPostCollectionHandler(servicePostCollection)
	handlerSinglePost := initSinglePostHandler(serviceSinglePost, servicePost, serviceClassicUser, serviceClassicUserFollowings, serviceProfileSettings, servicePostContent, serviceLocation, servicePostTagPost, serviceTag)

	handleFunc(handlerActivity, handlerComment, handlerPost, handlerPostAlbum, handlerPostCollection, handlerSinglePost, handlerPostCollectionPosts)
}