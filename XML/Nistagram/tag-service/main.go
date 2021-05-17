package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/tag-service/service"
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

	db.AutoMigrate(&model.Tag{},&model.StoryTag{},&model.StoryTagStories{},&model.CommentTag{},&model.CommentTagComments{},&model.PostTag{},&model.PostTagPosts{})
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

func initPostTagRepo(database *gorm.DB) *repository.PostTagRepository{
	return &repository.PostTagRepository { Database: database }
}

func initStoryTagRepo(database *gorm.DB) *repository.StoryTagRepository{
	return &repository.StoryTagRepository { Database: database }
}

func initCommentTagRepo(database *gorm.DB) *repository.CommentTagRepository{
	return &repository.CommentTagRepository { Database: database }
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

func initTagRepo(database *gorm.DB) *repository.TagRepository{
	return &repository.TagRepository { Database: database }
}

func initTagServices(repo *repository.TagRepository) *service.TagService{
	return &service.TagService { Repo: repo }
}

func initPostTagServices(repo *repository.PostTagRepository) *service.PostTagService{
	return &service.PostTagService { Repo: repo }
}

func initCommentTagServices(repo *repository.CommentTagRepository) *service.CommentTagService{
	return &service.CommentTagService { Repo: repo }
}

func initStoryTagServices(repo *repository.StoryTagRepository) *service.StoryTagService{
	return &service.StoryTagService { Repo: repo }
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

func initTagHandler(service *service.TagService) *handler.TagHandler{
	return &handler.TagHandler { Service: service }
}

func initPostTagHandler(service *service.PostTagService, tagService * service.TagService) *handler.PostTagHandler{
	return &handler.PostTagHandler { Service: service, TagService: tagService }
}

func initStoryTagHandler(service *service.StoryTagService, tagService *service.TagService) *handler.StoryTagHandler{
	return &handler.StoryTagHandler { Service: service, TagService: tagService }
}

func initCommentTagHandler(service *service.CommentTagService) *handler.CommentTagHandler{
	return &handler.CommentTagHandler { Service: service }
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

func handleFunc(handlerTag *handler.TagHandler,handlerPostTag *handler.PostTagHandler,handlerStoryTag *handler.StoryTagHandler,
	handlerCommentTag *handler.CommentTagHandler, handlerCommentTagComments *handler.CommentTagCommentsHandler,handlerPostTagPosts *handler.PostTagPostsHandler,
	handlerStoryTagStories *handler.StoryTagStoriesHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/comment_tag/", handlerCommentTag.CreateCommentTag).Methods("POST")
	router.HandleFunc("/comment_tag_comments/", handlerCommentTagComments.CreateCommentTagComments).Methods("POST")
	router.HandleFunc("/post_tag_posts/", handlerPostTagPosts.CreatePostTagPosts).Methods("POST")
	router.HandleFunc("/story_tag_stories/", handlerStoryTagStories.CreateStoryTagStories).Methods("POST")

	mux := http.NewServeMux()
	mux.HandleFunc("/tag/", handlerTag.CreateTag)
	mux.HandleFunc("/post_tag/", handlerPostTag.CreatePostTag)
	mux.HandleFunc("/story_tag/", handlerStoryTag.CreateStoryTag)
	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8082", handlerVar))
}

func main() {
	database := initDB()
	repoTag := initTagRepo(database)
	serviceTag := initTagServices(repoTag)
	handlerTag := initTagHandler(serviceTag)

	repoPostTag := initPostTagRepo(database)
	servicePostTag := initPostTagServices(repoPostTag)
	handlerPostTag := initPostTagHandler(servicePostTag, serviceTag)

	repoStoryTag := initStoryTagRepo(database)
	serviceStoryTag := initStoryTagServices(repoStoryTag)
	handlerStoryTag := initStoryTagHandler(serviceStoryTag, serviceTag)

	repoCommentTag := initCommentTagRepo(database)
	serviceCommentTag := initCommentTagServices(repoCommentTag)
	handlerCommentTag := initCommentTagHandler(serviceCommentTag)

	repoPostTagPosts := initPostTagPostsRepo(database)
	servicePostTagPosts := initPostTagPostsServices(repoPostTagPosts)
	handlerPostTagPosts := initPostTagPostsHandler(servicePostTagPosts)

	repoStoryTagStories := initStoryTagStoriesRepo(database)
	serviceStoryTagStories := initStoryTagStoriesServices(repoStoryTagStories)
	handlerStoryTagStories := initStoryTagStoriesHandler(serviceStoryTagStories)

	repoCommentTagComments := initCommentTagCommentsRepo(database)
	serviceCommentTagComments := initCommentTagCommentsServices(repoCommentTagComments)
	handlerCommentTagComments := initCommentTagCommentsHandler(serviceCommentTagComments)
	handleFunc(handlerTag,handlerPostTag,handlerStoryTag,handlerCommentTag,handlerCommentTagComments,handlerPostTagPosts,handlerStoryTagStories)
}