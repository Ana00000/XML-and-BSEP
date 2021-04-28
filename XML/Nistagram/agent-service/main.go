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

	db.AutoMigrate(&model.Product{})
	return db
}

func initRepo(database *gorm.DB) *repository.ProductRepository{
	return &repository.ProductRepository { Database: database }
}

func initServices(repo *repository.ProductRepository) *service.ProductService{
	return &service.ProductService { Repo: repo }
}

func initHandler(service *service.ProductService) *handler.ProductHandler{
	return &handler.ProductHandler { Service: service }
}

func handleFunc(handler *handler.ProductHandler){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.CreateProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8082"), router))
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}