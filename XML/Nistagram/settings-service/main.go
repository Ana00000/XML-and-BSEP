package main

import (
	registerUserPath "../user-service/model"
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

	db.AutoMigrate(&model.ProfileSettings{}, &registerUserPath.RegisteredUser{})
	return db
}

func initRepo(database *gorm.DB) *repository.ProfileSettingsRepository{
	return &repository.ProfileSettingsRepository { Database: database }
}

func initServices(repo *repository.ProfileSettingsRepository) *service.ProfileSettingsService{
	return &service.ProfileSettingsService { Repo: repo }
}

func initHandler(service *service.ProfileSettingsService) *handler.ProfileSettingsHandler{
	return &handler.ProfileSettingsHandler { Service: service }
}

func handleFunc(handler *handler.ProfileSettingsHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/settings/", handler.CreateProfileSettings).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}