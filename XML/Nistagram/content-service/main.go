package main

import (
	"./handler"
	"./model"
	"./repository"
	"./service"
	"fmt"
	_ "fmt"
	_ "github.com/antchfx/xpath"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	db.AutoMigrate(&model.Content{}, &model.AdvertisementContent{})
	return db
}

func initContentRepo(database *gorm.DB) *repository.ContentRepository{
	return &repository.ContentRepository { Database: database }
}

func initAdvertisementContentRepo(database *gorm.DB) *repository.AdvertisementContentRepository{
	return &repository.AdvertisementContentRepository { Database: database }
}

func initContentService(repo *repository.ContentRepository) *service.ContentService{
	return &service.ContentService { Repo: repo }
}

func initAdvertisementContentService(repo *repository.AdvertisementContentRepository) *service.AdvertisementContentService{
	return &service.AdvertisementContentService { Repo: repo }
}

func initContentHandler(service *service.ContentService) *handler.ContentHandler{
	return &handler.ContentHandler { Service: service }
}

func initAdvertisementContentHandler(service *service.AdvertisementContentService) *handler.AdvertisementContentHandler{
	return &handler.AdvertisementContentHandler { Service: service }
}

func handleFunc(handlerContent *handler.ContentHandler, handlerAdvertisementContent *handler.AdvertisementContentHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/content/", handlerContent.CreateContent).Methods("POST")
	router.HandleFunc("/advertisement_content/", handlerAdvertisementContent.CreateAdvertisementContent).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repoContent := initContentRepo(database)
	repoAdvertisementContent := initAdvertisementContentRepo(database)
	serviceContent := initContentService(repoContent)
	serviceAdvertisementContent := initAdvertisementContentService(repoAdvertisementContent)
	handlerContent := initContentHandler(serviceContent)
	handlerAdvertisementContent := initAdvertisementContentHandler(serviceAdvertisementContent)
	handleFunc(handlerContent, handlerAdvertisementContent)
}