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

	db.AutoMigrate(&model.User{}, &model.RegisteredUser{})
	return db
}

func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}

func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler { Service: service }
}

func initRegisteredUserHandler(service *service.RegisteredUserService) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler { Service: service }
}

func handleFunc(userHandler *handler.UserHandler, registeredUserHandler *handler.RegisteredUserHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user/", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	userRepo := initUserRepo(database)
	registeredUserRepo := initRegisteredUserRepo(database)
	userService := initUserService(userRepo)
	registeredUserService := initRegisteredUserService(registeredUserRepo)
	userHandler := initUserHandler(userService)
	registeredUserHandler := initRegisteredUserHandler(registeredUserService)
	handleFunc(userHandler, registeredUserHandler)
}