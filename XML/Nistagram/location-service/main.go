package main

import (
	_ "fmt"
	_ "github.com/antchfx/xpath"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/handler"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/model"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/repository"
	"github.com/xml/XML-and-BSEP/XML/Nistagram/location-service/service"
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

	db.AutoMigrate(&model.Location{})
	return db
}

func initRepo(database *gorm.DB) *repository.LocationRepository{
	return &repository.LocationRepository { Database: database }
}

func initServices(repo *repository.LocationRepository) *service.LocationService{
	return &service.LocationService { Repo: repo }
}

func initHandler(service *service.LocationService) *handler.LocationHandler{
	return &handler.LocationHandler { Service: service }
}

func handleFunc(handler *handler.LocationHandler){
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.CreateLocation)
	handlerVar := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8083", handlerVar))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}