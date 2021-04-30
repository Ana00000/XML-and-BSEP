package main

import (
	"./handler"
	"./model"
	"./repository"
	"./service"
	"fmt"
	"github.com/gorilla/mux"
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

	db.AutoMigrate(&model.Activity{})
	return db
}

func initActivityRepo(database *gorm.DB) *repository.ActivityRepository{
	return &repository.ActivityRepository{ Database: database }
}

func initActivityService(repo *repository.ActivityRepository) *service.ActivityService{
	return &service.ActivityService{ Repo: repo }
}

func initActivityHandler(service *service.ActivityService) *handler.ActivityHandler{
	return &handler.ActivityHandler{ Service: service }
}

func handleFunc(handlerActivity *handler.ActivityHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/activity/", handlerActivity.CreateActivity).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repoActivity := initActivityRepo(database)
	serviceActivity := initActivityService(repoActivity)
	handlerActivity := initActivityHandler(serviceActivity)
	handleFunc(handlerActivity)
}