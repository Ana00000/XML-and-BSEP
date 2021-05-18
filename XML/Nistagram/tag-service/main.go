package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
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
	_ "os"
	_ "strconv"
)

func initDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=root dbname=nistagram-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Tag{},&model.StoryTag{},&model.StoryTagStories{},&model.CommentTag{},
				   &model.CommentTagComments{},&model.PostTag{},&model.PostTagPosts{},
				   &model.PostAlbumTag{},&model.PostAlbumTagPostAlbums{})
	return db
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

func initPostAlbumTagRepo(database *gorm.DB) *repository.PostAlbumTagRepository{
	return &repository.PostAlbumTagRepository { Database: database }
}

func initPostAlbumTagPostAlbumsRepo(database *gorm.DB) *repository.PostAlbumTagPostAlbumsRepository{
	return &repository.PostAlbumTagPostAlbumsRepository { Database: database }
}

func initStoryAlbumTagRepo(database *gorm.DB) *repository.StoryAlbumTagRepository{
	return &repository.StoryAlbumTagRepository { Database: database }
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

func initPostAlbumTagServices(repo *repository.PostAlbumTagRepository) *service.PostAlbumTagService{
	return &service.PostAlbumTagService { Repo: repo }
}

func initPostAlbumTagPostAlbumsServices(repo *repository.PostAlbumTagPostAlbumsRepository) *service.PostAlbumTagPostAlbumsService{
	return &service.PostAlbumTagPostAlbumsService { Repo: repo }
}

func initStoryAlbumTagServices(repo *repository.StoryAlbumTagRepository) *service.StoryAlbumTagService{
	return &service.StoryAlbumTagService { Repo: repo }
}

func initStoryAlbumTagStoryAlbumsServices(repo *repository.StoryAlbumTagStoryAlbumsRepository) *service.StoryAlbumTagStoryAlbumsService{
	return &service.StoryAlbumTagStoryAlbumsService { Repo: repo }
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

func initPostAlbumTagHandler(service *service.PostAlbumTagService, tagService *service.TagService) *handler.PostAlbumTagHandler{
	return &handler.PostAlbumTagHandler { Service: service, TagService : tagService }
}

func initPostAlbumTagPostAlbumsHandler(service *service.PostAlbumTagPostAlbumsService) *handler.PostAlbumTagPostAlbumsHandler{
	return &handler.PostAlbumTagPostAlbumsHandler { Service: service }
}

func initStoryAlbumTagHandler(service *service.StoryAlbumTagService, tagService *service.TagService) *handler.StoryAlbumTagHandler{
	return &handler.StoryAlbumTagHandler { Service: service, TagService : tagService }
}

func initStoryAlbumTagStoryAlbumsHandler(service *service.StoryAlbumTagStoryAlbumsService) *handler.StoryAlbumTagStoryAlbumsHandler{
	return &handler.StoryAlbumTagStoryAlbumsHandler { Service: service }
}

func handleFunc(handlerTag *handler.TagHandler, handlerPostTag *handler.PostTagHandler, handlerStoryTag *handler.StoryTagHandler,
	handlerCommentTag *handler.CommentTagHandler, handlerCommentTagComments *handler.CommentTagCommentsHandler,
	handlerPostTagPosts *handler.PostTagPostsHandler, handlerStoryTagStories *handler.StoryTagStoriesHandler,
	handlerPostAlbumTag *handler.PostAlbumTagHandler, handlerPostAlbumTagPostAlbums *handler.PostAlbumTagPostAlbumsHandler,
	handlerStoryAlbumTag *handler.StoryAlbumTagHandler, handlerStoryAlbumTagStoryAlbums *handler.StoryAlbumTagStoryAlbumsHandler){

	mux := http.NewServeMux()

	mux.HandleFunc("/tag/", handlerTag.CreateTag)
	mux.HandleFunc("/post_tag/", handlerPostTag.CreatePostTag)
	mux.HandleFunc("/story_tag/", handlerStoryTag.CreateStoryTag)
	mux.HandleFunc("/comment_tag/", handlerCommentTag.CreateCommentTag)
	mux.HandleFunc("/comment_tag_comments/", handlerCommentTagComments.CreateCommentTagComments)
	mux.HandleFunc("/post_tag_posts/", handlerPostTagPosts.CreatePostTagPosts)
	mux.HandleFunc("/story_tag_stories/", handlerStoryTagStories.CreateStoryTagStories)
	mux.HandleFunc("/post_album_tag/", handlerPostAlbumTag.CreatePostAlbumTag)
	mux.HandleFunc("/post_album_tag_post_albums/", handlerPostAlbumTagPostAlbums.CreatePostAlbumTagPostAlbums)
	mux.HandleFunc("/story_album_tag/", handlerStoryAlbumTag.CreateStoryAlbumTag)
	mux.HandleFunc("/story_album_tag_story_albums/", handlerStoryAlbumTagStoryAlbums.CreateStoryAlbumTagStoryAlbums)

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

	repoPostAlbumTag := initPostAlbumTagRepo(database)
	servicePostAlbumTag := initPostAlbumTagServices(repoPostAlbumTag)
	handlerPostAlbumTag := initPostAlbumTagHandler(servicePostAlbumTag, serviceTag)

	repoPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsRepo(database)
	servicePostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsServices(repoPostAlbumTagPostAlbums)
	handlerPostAlbumTagPostAlbums := initPostAlbumTagPostAlbumsHandler(servicePostAlbumTagPostAlbums)

	repoStoryAlbumTag := initStoryAlbumTagRepo(database)
	serviceStoryAlbumTag := initStoryAlbumTagServices(repoStoryAlbumTag)
	handlerStoryAlbumTag := initStoryAlbumTagHandler(serviceStoryAlbumTag, serviceTag)

	repoStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsRepo(database)
	serviceStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsServices(repoStoryAlbumTagStoryAlbums)
	handlerStoryAlbumTagStoryAlbums := initStoryAlbumTagStoryAlbumsHandler(serviceStoryAlbumTagStoryAlbums)

	handleFunc(handlerTag, handlerPostTag, handlerStoryTag, handlerCommentTag, handlerCommentTagComments,
		       handlerPostTagPosts, handlerStoryTagStories, handlerPostAlbumTag, handlerPostAlbumTagPostAlbums,
		       handlerStoryAlbumTag, handlerStoryAlbumTagStoryAlbums)
}