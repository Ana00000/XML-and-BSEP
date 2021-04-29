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

	db.AutoMigrate(&model.InappropriateContentRequest{})
	return db
}

func initInappropriateContentRequestRepo(database *gorm.DB) *repository.InappropriateContentRequestRepository{
	return &repository.InappropriateContentRequestRepository { Database: database }
}

func initInappropriateContentRequestServices(repo *repository.InappropriateContentRequestRepository) *service.InappropriateContentRequestService{
	return &service.InappropriateContentRequestService { Repo: repo }
}

func initInappropriateContentRequestHandler(service *service.InappropriateContentRequestService) *handler.InappropriateContentRequestHandler{
	return &handler.InappropriateContentRequestHandler { Service: service }
}

func handleFunc(inappropriateContentRequestHandler *handler.InappropriateContentRequestHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/inappropriateContentRequest", inappropriateContentRequestHandler.CreateInappropriateContentRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8083"), router))
}

func main() {
	database := initDB()
	inappropriateContentRequestRepo := initInappropriateContentRequestRepo(database)
	inappropriateContentRequestService := initInappropriateContentRequestServices(inappropriateContentRequestRepo)
	inappropriateContentRequestHandler := initInappropriateContentRequestHandler(inappropriateContentRequestService)
	handleFunc(inappropriateContentRequestHandler)
}