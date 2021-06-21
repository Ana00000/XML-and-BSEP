package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/post-service/service"
	"gopkg.in/go-playground/validator.v9"
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

	db.AutoMigrate(&model.Post{}, &model.Activity{}, &model.Comment{},
	               &model.PostAlbum{}, &model.SinglePost{}, &model.PostCollection{},
	               &model.PostCollectionPosts{})
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

func initPostCollectionPostsRepo(database *gorm.DB) *repository.PostCollectionPostsRepository{
	return &repository.PostCollectionPostsRepository { Database: database }
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

func initPostCollectionPostsServices(repo *repository.PostCollectionPostsRepository) *service.PostCollectionPostsService{
	return &service.PostCollectionPostsService { Repo: repo }
}

func initActivityHandler(SinglePostService *service.SinglePostService,logInfo *logrus.Logger, logError *logrus.Logger,service *service.ActivityService) *handler.ActivityHandler{
	return &handler.ActivityHandler{
		Service:           service,
		SinglePostService: SinglePostService,
		LogInfo:           logInfo,
		LogError:          logError,
	}
}

func initCommentHandler(SinglePostService *service.SinglePostService,logInfo *logrus.Logger, logError *logrus.Logger,service *service.CommentService, validate *validator.Validate) *handler.CommentHandler{
	return &handler.CommentHandler{
		Service:           service,
		SinglePostService: SinglePostService,
		Validator:         validate,
		LogInfo:           logInfo,
		LogError:          logError,
	}
}

func initPostHandler(logInfo *logrus.Logger, logError *logrus.Logger,postService *service.PostService) *handler.PostHandler{
	return &handler.PostHandler{ LogInfo: logInfo, LogError: logError, PostService: postService}
}

func initPostAlbumHandler(logInfo *logrus.Logger, logError *logrus.Logger, service *service.PostAlbumService, postService *service.PostService) *handler.PostAlbumHandler{
	return &handler.PostAlbumHandler{ LogInfo: logInfo, LogError: logError, Service: service, PostService: postService}
}

func initPostCollectionHandler(logInfo *logrus.Logger, logError *logrus.Logger, service *service.PostCollectionService) *handler.PostCollectionHandler{
	return &handler.PostCollectionHandler{ LogInfo: logInfo, LogError: logError, Service: service }
}

func initSinglePostHandler(logInfo *logrus.Logger, logError *logrus.Logger, singlePostService *service.SinglePostService, postService *service.PostService,) *handler.SinglePostHandler{
	return &handler.SinglePostHandler{ LogInfo: logInfo, LogError: logError, SinglePostService: singlePostService, PostService: postService}
}

func initPostCollectionPostsHandler(logInfo *logrus.Logger, logError *logrus.Logger, service *service.PostCollectionPostsService) *handler.PostCollectionPostsHandler{
	return &handler.PostCollectionPostsHandler { LogInfo: logInfo, LogError: logError, Service: service }
}


func handleFunc(handlerActivity *handler.ActivityHandler, handlerComment *handler.CommentHandler, handlerPost *handler.PostHandler,
	handlerPostAlbum *handler.PostAlbumHandler, handlerPostCollection *handler.PostCollectionHandler,
	handlerSinglePost *handler.SinglePostHandler, handlerPostCollectionPosts *handler.PostCollectionPostsHandler){

	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/post/", handlerPost.CreatePost).Methods("POST")
	router.HandleFunc("/post_album/", handlerPostAlbum.CreatePostAlbum).Methods("POST")
	router.HandleFunc("/single_post/", handlerSinglePost.CreateSinglePost).Methods("POST")
	router.HandleFunc("/activity/", handlerActivity.CreateActivity).Methods("POST")
	router.HandleFunc("/update_activity/", handlerActivity.UpdateActivity).Methods("POST")

	router.HandleFunc("/find_all_likes_for_post", handlerActivity.FindAllLikesForPost).Methods("GET")
	router.HandleFunc("/find_all_dislikes_for_post", handlerActivity.FindAllDislikesForPost).Methods("GET")
	router.HandleFunc("/find_all_favorites_for_post", handlerActivity.FindAllFavoritesForPost).Methods("GET")
	router.HandleFunc("/find_all_activities_for_post", handlerActivity.FindAllActivitiesForPost).Methods("GET")

	router.HandleFunc("/comment/", handlerComment.CreateComment).Methods("POST")
	router.HandleFunc("/find_all_comments_for_post", handlerComment.FindAllCommentsForPost).Methods("GET")
	router.HandleFunc("/find_all_user_comments", handlerComment.FindAllUserComments).Methods("GET")
	router.HandleFunc("/update_post/", handlerPost.UpdatePost).Methods("POST")
	router.HandleFunc("/post_collection/", handlerPostCollection.CreatePostCollection).Methods("POST")
	router.HandleFunc("/post_collection_posts/", handlerPostCollectionPosts.CreatePostCollectionPosts).Methods("POST")

	router.HandleFunc("/find_all_post_collections_for_reg", handlerPostCollection.FindAllPostCollectionsForUserRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_post_collection_posts_for_post", handlerPostCollectionPosts.FindAllPostCollectionPostsForPost).Methods("GET")

	router.HandleFunc("/find_all_posts_for_not_reg", handlerSinglePost.FindAllPostsForUserNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_posts_for_reg", handlerSinglePost.FindAllPostsForUserRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_following_posts", handlerSinglePost.FindAllFollowingPosts).Methods("GET")
	router.HandleFunc("/find_selected_post_not_reg", handlerSinglePost.FindSelectedPostByIdForNotRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_selected_post_reg", handlerSinglePost.FindSelectedPostByIdForRegisteredUsers).Methods("GET")
	router.HandleFunc("/find_all_public_posts_not_reg/", handlerSinglePost.FindAllPublicPostsNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_public_posts_reg", handlerSinglePost.FindAllPublicPostsRegisteredUser).Methods("GET")


	router.HandleFunc("/find_all_public_album_posts_reg", handlerPostAlbum.FindAllPublicAlbumPostsRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_public_album_posts_not_reg/", handlerPostAlbum.FindAllPublicAlbumPostsNotRegisteredUser).Methods("GET")
	router.HandleFunc("/find_all_following_post_albums", handlerPostAlbum.FindAllFollowingPostAlbums).Methods("GET")

	router.HandleFunc("/find_all_album_posts_for_logged_user", handlerPostAlbum.FindAllAlbumPostsForLoggedUser).Methods("GET")
	router.HandleFunc("/find_selected_post_album_for_logged_user", handlerPostAlbum.FindSelectedPostAlbumByIdForLoggedUser).Methods("GET")
	router.HandleFunc("/find_all_posts_for_logged_user", handlerSinglePost.FindAllPostsForLoggedUser).Methods("GET")
	router.HandleFunc("/find_selected_post_for_logged_user", handlerSinglePost.FindSelectedPostByIdForLoggedUser).Methods("GET")

	//SEARCH FOR NOT REGISTERED USER
	router.HandleFunc("/find_all_tags_for_public_posts/", handlerSinglePost.FindAllTagsForPublicPosts).Methods("GET")
	router.HandleFunc("/find_all_locations_for_public_posts/", handlerSinglePost.FindAllLocationsForPublicPosts).Methods("GET")

	//metoda koja se poziva kada neregistrovani user pretrazi tag pa klikne na njega - prikazuju se svi PUBLIC, NOT DELETED postovi sa tim tagom
	router.HandleFunc("/find_all_posts_for_tag", handlerSinglePost.FindAllPostsForTag).Methods("GET")
	router.HandleFunc("/find_all_posts_for_location", handlerSinglePost.FindAllPostsForLocation).Methods("GET")

	//SEARCH FOR REGISTERED USER
	router.HandleFunc("/find_all_tags_for_public_and_friends_posts", handlerSinglePost.FindAllTagsForPublicAndFollowingPosts).Methods("GET")
	router.HandleFunc("/find_all_locations_for_public_friends_posts", handlerSinglePost.FindAllLocationsForPublicAndFollowingPosts).Methods("GET")

	//metoda koja se poziva kada neregistrovani user pretrazi tag pa klikne na njega - prikazuju se svi PUBLIC, NOT DELETED postovi sa tim tagom
	router.HandleFunc("/find_all_posts_for_tag_reg_user", handlerSinglePost.FindAllPostsForTagRegUser).Methods("GET")
	router.HandleFunc("/find_all_posts_for_location_reg_user", handlerSinglePost.FindAllPostsForLocationRegUser).Methods("GET")

	router.HandleFunc("/find_owner_of_post/{id}", handlerSinglePost.FindOwnerOfPost).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	logInfo := logrus.New()
	logError := logrus.New()

	LogInfoFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logInfoPOST.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	LogErrorFile, err := os.OpenFile(os.Getenv("LOG_URL")+"/logErrorPOST.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	logInfo.Out = LogInfoFile
	logInfo.Formatter = &logrus.JSONFormatter{}
	logError.Out = LogErrorFile
	logError.Formatter = &logrus.JSONFormatter{}

	database := initDB()
	validator := validator.New()

	repoActivity := initActivityRepo(database)
	repoComment := initCommentRepo(database)
	repoPost := initPostRepo(database)
	repoPostAlbum := initPostAlbumRepo(database)
	repoPostCollection := initPostCollectionRepo(database)
	repoSinglePost := initSinglePostRepo(database)
	repoPostCollectionPosts := initPostCollectionPostsRepo(database)

	serviceActivity := initActivityService(repoActivity)
	serviceComment := initCommentService(repoComment)
	servicePost := initPostService(repoPost)
	servicePostAlbum := initPostAlbumService(repoPostAlbum)
	servicePostCollection := initPostCollectionService(repoPostCollection)
	serviceSinglePost := initSinglePostService(repoSinglePost)
	servicePostCollectionPosts := initPostCollectionPostsServices(repoPostCollectionPosts)

	handlerActivity := initActivityHandler(serviceSinglePost,logInfo, logError, serviceActivity)
	handlerComment := initCommentHandler(serviceSinglePost,logInfo, logError, serviceComment, validator)
	handlerPost := initPostHandler(logInfo, logError, servicePost)
	handlerPostAlbum := initPostAlbumHandler(logInfo, logError, servicePostAlbum, servicePost)
	handlerPostCollection := initPostCollectionHandler(logInfo, logError, servicePostCollection)
	handlerSinglePost := initSinglePostHandler(logInfo, logError, serviceSinglePost, servicePost)
	handlerPostCollectionPosts := initPostCollectionPostsHandler(logInfo, logError, servicePostCollectionPosts)

	handleFunc(handlerActivity, handlerComment, handlerPost, handlerPostAlbum, handlerPostCollection, handlerSinglePost, handlerPostCollectionPosts)
}