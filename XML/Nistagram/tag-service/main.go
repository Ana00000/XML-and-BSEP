package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
	"gopkg.in/go-playground/validator.v9"
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

	db.AutoMigrate(&model.Tag{}, &model.UserTag{}, &model.StoryTagStories{},
				   &model.CommentTagComments{}, &model.PostTagPosts{},
				   &model.PostAlbumTagPostAlbums{}, &model.StoryAlbumTagStoryAlbums{})
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

func initUserTagRepo(database *gorm.DB) *repository.UserTagRepository{
	return &repository.UserTagRepository { Database: database }
}

func initTagServices(repo *repository.TagRepository) *service.TagService{
	return &service.TagService { Repo: repo }
}

func initUserTagServices(repo *repository.UserTagRepository) *service.UserTagService{
	return &service.UserTagService { Repo: repo }
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

func initTagHandler(service *service.TagService, validator *validator.Validate) *handler.TagHandler{
	return &handler.TagHandler {
		Service: service,
		Validator: validator,
	}
}

func initUserTagHandler(service *service.UserTagService, tagService * service.TagService, validator *validator.Validate) *handler.UserTagHandler{
	return &handler.UserTagHandler {
		Service: service,
		TagService: tagService,
		Validator: validator,
	}
}

func initStoryTagStoriesHandler(service *service.StoryTagStoriesService) *handler.StoryTagStoriesHandler{
	return &handler.StoryTagStoriesHandler { Service: service }
}

func initCommentTagCommentsHandler(service *service.CommentTagCommentsService,tagService *service.TagService) *handler.CommentTagCommentsHandler{
	return &handler.CommentTagCommentsHandler { Service: service, TagService: tagService}
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

func handleFunc(handlerTag *handler.TagHandler, handlerUserTag *handler.UserTagHandler, handlerCommentTagComments *handler.CommentTagCommentsHandler,
	handlerPostTagPosts *handler.PostTagPostsHandler, handlerStoryTagStories *handler.StoryTagStoriesHandler,  handlerPostAlbumTagPostAlbums *handler.PostAlbumTagPostAlbumsHandler, handlerStoryAlbumTagStoryAlbums *handler.StoryAlbumTagStoryAlbumsHandler){


	router := mux.NewRouter().StrictSlash(true)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "Access-Control-Allow-Headers"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"https://localhost:8081"}),
		handlers.AllowCredentials(),
	)

	router.HandleFunc("/tag/", handlerTag.CreateTag).Methods("POST")
	router.HandleFunc("/get_tag_name_by_id/{id}", handlerTag.FindTagNameById).Methods("GET")
	router.HandleFunc("/find_all_hashtags/", handlerTag.FindAllHashTags).Methods("GET")
	router.HandleFunc("/user_tag/", handlerUserTag.CreateUserTag).Methods("POST")
	router.HandleFunc("/create_user_tag_for_registered_user/", handlerUserTag.CreateUserTagForRegisteredUser).Methods("POST")
	router.HandleFunc("/find_all_taggable_users_post/", handlerUserTag.FindAllTaggableUsersPost).Methods("GET")
	router.HandleFunc("/find_all_taggable_users_story/", handlerUserTag.FindAllTaggableUsersStory).Methods("GET")
	router.HandleFunc("/find_all_taggable_users_comment/", handlerUserTag.FindAllTaggableUsersComment).Methods("GET")
	router.HandleFunc("/comment_tag_comments/", handlerCommentTagComments.CreateCommentTagComments).Methods("POST")
	router.HandleFunc("/post_tag_posts/", handlerPostTagPosts.CreatePostTagPosts).Methods("POST")
	router.HandleFunc("/story_tag_stories/", handlerStoryTagStories.CreateStoryTagStories).Methods("POST")
	router.HandleFunc("/post_album_tag_post_albums/", handlerPostAlbumTagPostAlbums.CreatePostAlbumTagPostAlbums).Methods("POST")
	router.HandleFunc("/story_album_tag_story_albums/", handlerStoryAlbumTagStoryAlbums.CreateStoryAlbumTagStoryAlbums).Methods("POST")

	router.HandleFunc("/find_comment_tag_comments_for_comment/{id}", handlerCommentTagComments.FindAllCommentTagCommentsForComment).Methods("GET")
	router.HandleFunc("/get_tag_by_name/{name}", handlerTag.FindTagByName).Methods("GET")

	router.HandleFunc("/find_tag_id", handlerTag.FindTagForId).Methods("GET")
	router.HandleFunc("/find_story_tag_stories_for_story_id", handlerStoryTagStories.FindStoryTagStoriesForStoryId).Methods("GET")

	router.HandleFunc("/find_all_tags_for_stories/", handlerStoryTagStories.FindAllTagsForStories).Methods("POST")
	router.HandleFunc("/find_all_tags_for_story/", handlerStoryTagStories.FindAllTagsForStory).Methods("POST")

	router.HandleFunc("/find_all_tags_for_posts/", handlerPostTagPosts.FindAllTagsForPosts).Methods("POST")
	router.HandleFunc("/find_all_tags_for_post/", handlerPostTagPosts.FindAllTagsForPost).Methods("POST")

	router.HandleFunc("/find_all_tags_for_post_tag_posts/", handlerPostTagPosts.FindAllTagsForPostsTagPosts).Methods("POST")

	router.HandleFunc("/find_all_tags_for_post_album_tag_post_albums/", handlerPostAlbumTagPostAlbums.FindAllTagsForPostAlbumTagPostAlbums).Methods("POST")
	router.HandleFunc("/find_all_tags_for_post_album/", handlerPostAlbumTagPostAlbums.FindAllTagsForPostAlbum).Methods("POST")
	//FindPostIdsByTagId
	router.HandleFunc("/find_post_ids_by_tag_id/{tagID}", handlerPostTagPosts.FindPostIdsByTagId).Methods("GET")

	router.HandleFunc("/find_all_tags_for_story_album_tag_story_albums/",handlerStoryAlbumTagStoryAlbums.FindAllTagsForStoryAlbumTagStoryAlbums).Methods("POST")
	router.HandleFunc("/find_all_tags_for_story_album/", handlerStoryAlbumTagStoryAlbums.FindAllTagsForStoryAlbum).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), cors(router)))
}

func main() {
	database := initDB()
	validator := validator.New()

	repoTag := initTagRepo(database)
	serviceTag := initTagServices(repoTag)
	handlerTag := initTagHandler(serviceTag, validator)

	repoUserTag := initUserTagRepo(database)
	serviceUserTag := initUserTagServices(repoUserTag)
	handlerUserTag := initUserTagHandler(serviceUserTag, serviceTag, validator)

	repoPostTagPosts := initPostTagPostsRepo(database)
	servicePostTagPosts := initPostTagPostsServices(repoPostTagPosts)
	handlerPostTagPosts := initPostTagPostsHandler(servicePostTagPosts)

	repoStoryTagStories := initStoryTagStoriesRepo(database)
	serviceStoryTagStories := initStoryTagStoriesServices(repoStoryTagStories)
	handlerStoryTagStories := initStoryTagStoriesHandler(serviceStoryTagStories)

	repoCommentTagComments := initCommentTagCommentsRepo(database)
	serviceCommentTagComments := initCommentTagCommentsServices(repoCommentTagComments)
	handlerCommentTagComments := initCommentTagCommentsHandler(serviceCommentTagComments, serviceTag)


	repoPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsRepo(database)
	servicePostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsServices(repoPostAlbumTagPostAlbums)
	handlerPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsHandler(servicePostAlbumTagPostAlbums)


	repoStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsRepo(database)
	serviceStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsServices(repoStoryAlbumTagStoryAlbums)
	handlerStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsHandler(serviceStoryAlbumTagStoryAlbums)

	handleFunc(handlerTag, handlerUserTag, handlerCommentTagComments,handlerPostTagPosts, handlerStoryTagStories, handlerPostAlbumTagPostAlbums, handlerStoryAlbumTagStoryAlbums)
}