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

	db.AutoMigrate(&model.User{}, &model.RegisteredUser{}, &model.Admin{})
	return db
}

func initUserRepo(database *gorm.DB) *repository.UserRepository{
	return &repository.UserRepository { Database: database }
}

func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository{
	return &repository.RegisteredUserRepository { Database: database }
}

func initAdminRepo(database *gorm.DB) *repository.AdminRepository{
	return &repository.AdminRepository { Database: database }
}

func initUserService(repo *repository.UserRepository) *service.UserService{
	return &service.UserService { Repo: repo }
}

func initRegisteredUserService(repo *repository.RegisteredUserRepository) *service.RegisteredUserService{
	return &service.RegisteredUserService { Repo: repo }
}

func initAdminService(repo *repository.AdminRepository) *service.AdminService{
	return &service.AdminService { Repo: repo }
}

func initUserHandler(service *service.UserService) *handler.UserHandler{
	return &handler.UserHandler { Service: service }
}

func initRegisteredUserHandler(service *service.RegisteredUserService) *handler.RegisteredUserHandler{
	return &handler.RegisteredUserHandler { Service: service }
}

func initAdminHandler(service *service.AdminService) *handler.AdminHandler{
	return &handler.AdminHandler { Service: service }
}

func handleFunc(userHandler *handler.UserHandler, registeredUserHandler *handler.RegisteredUserHandler, adminHandler *handler.AdminHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user/", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/registered_user/", registeredUserHandler.CreateRegisteredUser).Methods("POST")
	router.HandleFunc("/admin/", adminHandler.CreateAdmin).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	userRepo := initUserRepo(database)
	registeredUserRepo := initRegisteredUserRepo(database)
	adminRepo := initAdminRepo(database)
	userService := initUserService(userRepo)
	registeredUserService := initRegisteredUserService(registeredUserRepo)
	adminService := initAdminService(adminRepo)
	userHandler := initUserHandler(userService)
	registeredUserHandler := initRegisteredUserHandler(registeredUserService)
	adminHandler := initAdminHandler(adminService)
	handleFunc(userHandler, registeredUserHandler, adminHandler)
}